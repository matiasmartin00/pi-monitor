package collectors

func StartCollectors() {
	go collectorCpuUsage()
	go collectorMemoryUsage()
	go collectorUptime()
	go collectorCpuCores()
	go collectorDiskUsage()
	go collectorCpuTemp()
	go collectorSpeedtest()
}
