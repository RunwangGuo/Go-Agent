package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"Go-Agent/internal/common"
)

func Listen() {
	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var hb common.Heartbeat
		if err := json.Unmarshal(body, &hb); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// 统一格式化输出
		fmt.Printf("[Server] ✅ Received heartbeat | Host: %-18s | CPU: %5.2f%% | MEM: %5.2f%% | Time: %s\n",
			hb.Hostname, hb.CPU, hb.MemMB,
			time.Now().Format("15:04:05"))

		// 若想同时打印原始 JSON，可以保留这一行
		// hb.Pretty()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("[Server] Listening on :8080 ...")
	http.ListenAndServe(":8080", nil)
}
