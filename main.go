package main

import (
	"log"
	"net/http"

	"github.com/matiasmartin00/pi-monitor/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("Starting pi-monitor")
	collectors.StartCollectors()
	http.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
	log.Println("Server started on: :8080")
	http.ListenAndServe(":8080", nil)
}
