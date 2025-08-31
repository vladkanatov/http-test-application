package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		info := map[string]string{
			"hostname":    getHostname(),
			"time":        time.Now().Format(time.RFC3339),
			"remote_addr": r.RemoteAddr,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	})

	port := getPort()
	_ = http.ListenAndServe(":"+port, nil)
}

func getHostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return host
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}
