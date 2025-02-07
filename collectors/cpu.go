package collectors

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_cpu_usage",
		Help: "Current CPU usage",
	})

	cpuCores = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_cpu_cores",
		Help: "Number of CPU cores",
	})
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(cpuCores)
}

func collectorCpuUsage() {
	for {
		percentage, err := cpu.Percent(1*time.Second, false)
		if err != nil {
			log.Println("Error getting CPU usage: ", err)
			continue
		}
		cpuUsage.Set(percentage[0])
		time.Sleep(5 * time.Second)
	}
}

func collectorCpuCores() {
	cores, err := cpu.Counts(true)
	if err != nil {
		log.Println("Error getting CPU cores: ", err)
		return
	}
	cpuCores.Set(float64(cores))
}
