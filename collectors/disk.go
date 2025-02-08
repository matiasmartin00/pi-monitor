package collectors

import (
	"log"
	"time"

	"github.com/matiasmartin00/pi-monitor/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/disk"
)

var (
	diskInterval, _ = time.ParseDuration("5s")

	diskTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_disk_total",
		Help: "Total disk space in bytes",
	})

	diskUsed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_disk_used",
		Help: "Used disk space in bytes",
	})

	diskUsagePercent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_disk_usage",
		Help: "Current disk usage in percent",
	})
)

func init() {
	prometheus.MustRegister(diskTotal)
	prometheus.MustRegister(diskUsed)
	prometheus.MustRegister(diskUsagePercent)
}

func setupDiskInterval() {
	if config.Config.Metrics.Disk.Interval != nil {
		diskInterval = *config.Config.Metrics.Disk.Interval
		return
	}
	log.Println("Disk interval not set, using default: ", diskInterval)
}

func collectorDiskUsage() {
	if !config.Config.Metrics.Disk.Enabled {
		return
	}
	for {
		log.Println("Getting disk usage")
		usage, err := disk.Usage("/")

		if err != nil {
			log.Println("Error getting disk usage: ", err)
			continue
		}

		diskTotal.Set(float64(usage.Total))
		diskUsed.Set(float64(usage.Used))
		diskUsagePercent.Set(usage.UsedPercent)

		time.Sleep(diskInterval)
	}
}
