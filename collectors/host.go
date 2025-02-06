package collectors

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/host"
)

var updatime = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "pi_monitor_uptime_seconds",
	Help: "Uptime in seconds",
})

func init() {
	prometheus.MustRegister(updatime)
}

func collectorUptime() {
	for {
		uptimeValue, err := host.Uptime()

		if err != nil {
			log.Println("Error getting uptime: ", err)
			continue
		}

		updatime.Set(float64(uptimeValue))
		time.Sleep(5 * time.Second)
	}
}
