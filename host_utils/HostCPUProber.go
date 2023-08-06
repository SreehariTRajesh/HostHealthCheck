package host_utils

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func HostCPUStatsTotal() float64 {
	totalCPUStats, _ := cpu.Percent(time.Second, false)
	currentTime := time.Now()
	log.Println(currentTime.Format(""), "Total CPU Usage(%):", totalCPUStats)
	return totalCPUStats[0]
}

func HostCPUStatsPerCPU() []float64 {
	perCPUStats, _ := cpu.Percent(time.Second, true)
	currentTime := time.Now()
	for i := 0; i < len(perCPUStats); i++ {
		log.Println(currentTime.Format(""), "CPU Usage(%): of Core", i, ":", perCPUStats[i])
	}
	return perCPUStats
}
