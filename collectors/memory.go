package collectors

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_usage",
		Help: "Current memory usage",
	})

	memoryTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_total",
		Help: "Total memory",
	})

	memoryFree = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_free",
		Help: "Free memory",
	})

	memoryUsed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_used",
		Help: "Used memory",
	})
)

func init() {
	prometheus.MustRegister(memoryUsage)
	prometheus.MustRegister(memoryTotal)
	prometheus.MustRegister(memoryFree)
	prometheus.MustRegister(memoryUsed)
}

func collectorMemoryUsage() {
	for {
		v, err := mem.VirtualMemory()
		if err != nil {
			log.Println("Error getting memory usage: ", err)
			continue
		}
		memoryUsage.Set(v.UsedPercent)
		memoryTotal.Set(float64(v.Total))
		memoryFree.Set(float64(v.Free))
		memoryUsed.Set(float64(v.Used))

		time.Sleep(5 * time.Second)
	}
}
