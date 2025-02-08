package collectors

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/showwin/speedtest-go/speedtest"
)

var (
	downloadSpeed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pi_monitor_speedtest_download",
		Help: "Current download speed in Mbps",
	}, []string{"server"})

	uploadSpeed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pi_monitor_speedtest_upload",
		Help: "Current upload speed in Mbps",
	}, []string{"server"})

	ping = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pi_monitor_speedtest_ping",
		Help: "Current ping in milliseconds",
	}, []string{"server"})
)

func init() {
	prometheus.MustRegister(downloadSpeed)
	prometheus.MustRegister(uploadSpeed)
	prometheus.MustRegister(ping)
}

func collectorSpeedtest() {
	for {
		speedtestClient := speedtest.New()
		serverList, err := speedtestClient.FetchServers()

		if err != nil {
			log.Println("Error getting server list: ", err)
			continue
		}

		targets, err := serverList.FindServer([]int{})

		if err != nil {
			log.Println("Error finding server: ", err)
			continue
		}

		for _, server := range targets {

			server.PingTest(nil)
			server.DownloadTest()
			server.UploadTest()

			downloadSpeed.WithLabelValues(server.Name).Set(float64(server.DLSpeed))
			uploadSpeed.WithLabelValues(server.Name).Set(float64(server.ULSpeed))
			ping.WithLabelValues(server.Name).Set(float64(server.Latency))

			server.Context.Reset()
		}

		time.Sleep(5 * time.Minute)
	}
}
