package collectors

func StartCollectors() {
	go collectorCpuUsage()
	go collectorMemoryUsage()
}
