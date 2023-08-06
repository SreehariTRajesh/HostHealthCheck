package webutils

import (
	"encoding/json"
	"goprober/host_utils"
	"log"
	"net/http"
)

type Disk_Usage struct {
	TotalDiskSpace uint64  `json: "TotalDiskSpace`
	FreeDiskSpace  uint64  `json: "FreeDiskSpace"`
	DiskSpaceUsed  uint64  `json: "UsedDiskSpace"`
	DiskUsedPCT    float64 `json: "DiskSpacePCT"`
}

func Disk_Usage_API() func(http.ResponseWriter, *http.Request) {
	DiskUsageHandler := func(res http.ResponseWriter, req *http.Request) {
		diskUsageStats := host_utils.HostDiskProber("/")
		disk_usage_data := Disk_Usage{
			TotalDiskSpace: diskUsageStats.Total,
			FreeDiskSpace:  diskUsageStats.Free,
			DiskSpaceUsed:  diskUsageStats.Used,
			DiskUsedPCT:    diskUsageStats.UsedPercent,
		}
		jsonDiskData, err := json.Marshal(disk_usage_data)
		if err != nil {
			log.Println("Error:", err)
		} else {
			res.Header().Set("Content-Type", "application/json")
			res.Write(jsonDiskData)
		}
	}
	return DiskUsageHandler
}
