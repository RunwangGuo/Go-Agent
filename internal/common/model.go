package common

import (
	"encoding/json"
	"fmt"
)

// Heartbeat 心跳数据结构
type Heartbeat struct {
	Hostname string  `json:"hostname"`
	CPU      float64 `json:"cpu"`
	MemMB    float64 `json:"mem_mb"`
	Time     string  `json:"time"`
}

// Pretty 打印格式化心跳
func (h Heartbeat) Pretty() {
	b, _ := json.MarshalIndent(h, "", "  ")
	fmt.Println(string(b))
}
