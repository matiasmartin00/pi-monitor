package collectors

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func collectorCpuTemp() {
	for {
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

		// La temperatura viene en miligrados, hay que dividir por 1000
		cpuTemp.Set(temp / 1000.0)
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
