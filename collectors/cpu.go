package collectors

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/matiasmartin00/pi-monitor/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	cpuInterval, _ = time.ParseDuration("5s")

	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_cpu_usage",
		Help: "Current CPU usage",
	})

	cpuCores = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_cpu_cores",
		Help: "Number of CPU cores",
	})

	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pi_monitor_cpu_temp",
		Help: "Current CPU temperature in Celsius",
	})
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(cpuCores)
	prometheus.MustRegister(cpuTemp)
}

func setupCpuInterval() {
	if config.Config.Metrics.Cpu.Interval != nil {
		cpuInterval = *config.Config.Metrics.Cpu.Interval
		return
	}
	log.Println("CPU interval not set, using default: ", cpuInterval)
}

func collectorCpuUsage() {
	if !config.Config.Metrics.Cpu.Enabled {
		return
	}
	for {
		log.Println("Getting CPU usage")
		percentage, err := cpu.Percent(1*time.Second, false)
		if err != nil {
			log.Println("Error getting CPU usage: ", err)
			continue
		}

		if len(percentage) == 0 {
			log.Println("No CPU usage data")
			continue
		}
		cpuUsage.Set(percentage[0])
		time.Sleep(cpuInterval)
	}
}

func collectorCpuTemp() {
	if !config.Config.Metrics.Cpu.Enabled {
		return
	}
	for {
		log.Println("Getting CPU temperature")
		sysPath := os.Getenv("HOST_SYS")
		if sysPath == "" {
			log.Println("HOST_SYS environment variable not set")
			return
		}

		path := fmt.Sprintf("%s/class/thermal/thermal_zone0/temp", sysPath)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Println("CPU temperature not available")
			return
		}

		data, err := os.ReadFile(path)

		if err != nil {
			log.Println("Error getting CPU temperature: ", err)
			continue
		}

		tempStr := strings.TrimSpace(string(data))

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			log.Println("Error parsing CPU temperature: ", err)
			continue
		}

		// The temperature is in millidegrees Celsius so we divide by 1000 to get Celsius
		cpuTemp.Set(temp / 1000.0)
		time.Sleep(cpuInterval)
	}
}

func collectorCpuCores() {
	if !config.Config.Metrics.Cpu.Enabled {
		return
	}
	log.Println("Getting CPU cores")
	cores, err := cpu.Counts(true)
	if err != nil {
		log.Println("Error getting CPU cores: ", err)
		return
	}
	cpuCores.Set(float64(cores))
}
