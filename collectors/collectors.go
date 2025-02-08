package collectors

func StartCollectors() {
	setupCpuInterval()
	setupMemoryInterval()
	setupHostInterval()
	setupDiskInterval()
	setupSpeedtestInterval()
	go collectorCpuUsage()
	go collectorMemoryUsage()
	go collectorUptime()
	go collectorCpuCores()
	go collectorDiskUsage()
	go collectorCpuTemp()
	go collectorSpeedtest()
}
