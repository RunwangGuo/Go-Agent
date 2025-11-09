package agent

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"Go-Agent/internal/common"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// Start 启动心跳上报
func Start() {
	// 1️⃣ 加载配置文件
	if err := common.LoadConfig("config.yaml"); err != nil {
		panic(err)
	}

	// 2️⃣ 初始化日志系统
	common.InitLogger(common.Cfg.LogLevel)

	client := &http.Client{Timeout: 5 * time.Second}
	hostname, _ := os.Hostname()

	common.Log.Infof("[Agent] Started. Reporting to %s", common.Cfg.ServerURL)

	for {
		// ==== 1. CPU 使用率（取 1 秒平均）====
		cpuPercent, err := cpu.Percent(time.Second, false)
		if err != nil || len(cpuPercent) == 0 {
			common.Log.Warnf("CPU collect error: %v", err)
			continue
		}

		// ==== 2. 系统内存使用率 ====
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			common.Log.Warnf("Memory collect error: %v", err)
			continue
		}

		// ==== 3. 构建心跳数据 ====
		hb := common.Heartbeat{
			Hostname: hostname,
			CPU:      cpuPercent[0],      // CPU 占用率 (%)
			MemMB:    vmStat.UsedPercent, // 内存使用率 (%)
			Time:     time.Now().Format(time.RFC3339),
		}

		// ==== 4. 序列化并上报 ====
		data, _ := json.Marshal(hb)
		resp, err := client.Post(common.Cfg.ServerURL, "application/json", bytes.NewReader(data))
		if err != nil {
			common.Log.Errorf("HTTP Post error: %v", err)
		} else {
			resp.Body.Close()
			common.Log.Infof("✅ Sent heartbeat | CPU: %.2f%% | MEM: %.2f%%", hb.CPU, hb.MemMB)
		}

		time.Sleep(time.Duration(common.Cfg.IntervalSec) * time.Second)
	}
}
