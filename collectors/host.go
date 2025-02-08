package collectors

import (
	"log"
	"time"

	"github.com/matiasmartin00/pi-monitor/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/host"
)

var (
	hostInterval, _ = time.ParseDuration("5s")
	updatime        = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_uptime_seconds",
		Help: "Uptime in seconds",
	})
)

func init() {
	prometheus.MustRegister(updatime)
}

func setupHostInterval() {
	if config.Config.Metrics.Host.Interval != nil {
		hostInterval = *config.Config.Metrics.Host.Interval
		return
	}
	log.Println("Host interval not set, using default: ", hostInterval)
}

func collectorUptime() {
	if !config.Config.Metrics.Host.Enabled {
		return
	}
	for {
		log.Println("Getting uptime")
		uptimeValue, err := host.Uptime()

		if err != nil {
			log.Println("Error getting uptime: ", err)
			continue
		}

		updatime.Set(float64(uptimeValue))
		time.Sleep(hostInterval)
	}
}
