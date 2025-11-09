package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"Go-Agent/internal/common"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// Start 启动心跳上报
func Start() {
	serverURL := "http://127.0.0.1:8080/heartbeat"

	hostname, _ := os.Hostname()
	client := &http.Client{Timeout: 5 * time.Second}

	fmt.Println("[Agent] Started. Reporting to", serverURL)

	for {
		// ==== 1. CPU 使用率（取 1 秒平均）====
		cpuPercent, err := cpu.Percent(time.Second, false)
		if err != nil || len(cpuPercent) == 0 {
			fmt.Println("[Agent] CPU collect error:", err)
			continue
		}

		// ==== 2. 系统内存使用率 ====
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("[Agent] Memory collect error:", err)
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
		resp, err := client.Post(serverURL, "application/json", bytes.NewReader(data))
		if err != nil {
			fmt.Println("[Agent] Error:", err)
		} else {
			resp.Body.Close()
			fmt.Printf("[Agent] Sent heartbeat ✅ CPU: %.2f%% | MEM: %.2f%%\n", hb.CPU, hb.MemMB)
		}

		time.Sleep(3 * time.Second)
	}
}
