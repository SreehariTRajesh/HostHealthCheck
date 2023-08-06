package host_utils

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/disk"
)

func HostDiskProber(path string) *disk.UsageStat {
	diskUsageStats, _ := disk.Usage(path)
	currTime := time.Now()
	log.Println(currTime.Format(""), "Total Disk Space:", diskUsageStats.Total)
	log.Println(currTime.Format(""), "Free Disk Space:", diskUsageStats.Free)
	log.Println(currTime.Format(""), "Used Disk Space:", diskUsageStats.Used)
	log.Println(currTime.Format(""), "Used Disk Space(%):", diskUsageStats.UsedPercent)
	return diskUsageStats

}
