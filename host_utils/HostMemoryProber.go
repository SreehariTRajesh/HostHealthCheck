package host_utils

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/mem"
)

func HostMemoryStats() *mem.VirtualMemoryStat {
	memoryStats, _ := mem.VirtualMemory()
	currTime := time.Now()
	log.Println(currTime.Format(""), "Total Memory:", memoryStats.Total)
	log.Println(currTime.Format(""), "Available Memory:", memoryStats.Available)
	log.Println(currTime.Format(""), "Used Memory:", memoryStats.Used)
	log.Println(currTime.Format(""), "Used Memory(%):", memoryStats.UsedPercent)
	return memoryStats
}
