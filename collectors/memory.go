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
		Help: "Current memory usage in percentage",
	})

	memoryTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_total",
		Help: "Total memory in bytes",
	})

	memoryFree = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_free",
		Help: "Free memory in bytes",
	})

	memoryUsed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_used",
		Help: "Used memory in bytes",
	})

	memoryCached = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_cached",
		Help: "Cached memory in bytes",
	})

	memoryBuffers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_buffers",
		Help: "Buffers memory in bytes",
	})

	memorySReclaimable = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_memory_sreclaimable",
		Help: "SReclaimable memory in bytes",
	})
)

func init() {
	prometheus.MustRegister(memoryUsage)
	prometheus.MustRegister(memoryTotal)
	prometheus.MustRegister(memoryFree)
	prometheus.MustRegister(memoryUsed)
	prometheus.MustRegister(memoryCached)
	prometheus.MustRegister(memoryBuffers)
	prometheus.MustRegister(memorySReclaimable)
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
		memoryCached.Set(float64(v.Cached))
		memoryBuffers.Set(float64(v.Buffers))
		memorySReclaimable.Set(float64(v.Sreclaimable))

		time.Sleep(5 * time.Second)
	}
}
