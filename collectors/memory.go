package collectors

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/mem"
)

var memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "pi_monitor_memory_usage",
	Help: "Current memory usage",
})

func init() {
	prometheus.MustRegister(memoryUsage)
}

func collectorMemoryUsage() {
	for {
		v, err := mem.VirtualMemory()
		if err != nil {
			log.Println("Error getting memory usage: ", err)
			continue
		}
		memoryUsage.Set(v.UsedPercent)
		time.Sleep(5 * time.Second)
	}
}
