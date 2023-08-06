package webutils

import (
	"encoding/json"
	"goprober/host_utils"
	"log"
	"net/http"
)

type CPU_Usage struct {
	CoreCount       int         `json:"CoreCount"`
	PerCoreUsage    interface{} `json:"PerCoreUsage"`
	Total_CPU_Usage interface{} `json:"Total_CPU_Usage"`
}

func CPU_Usage_API() func(http.ResponseWriter, *http.Request) {
	CPUUsageHandler := func(res http.ResponseWriter, req *http.Request) {
		Total_CPU_Usage := host_utils.HostCPUStatsTotal()
		Per_Processor_CPU_Usage := host_utils.HostCPUStatsPerCPU()
		cpu_data := CPU_Usage{
			CoreCount:       len(Per_Processor_CPU_Usage),
			PerCoreUsage:    Per_Processor_CPU_Usage,
			Total_CPU_Usage: Total_CPU_Usage,
		}
		jsonCPUData, err := json.Marshal(cpu_data)
		if err != nil {
			log.Println("Error:", err)
		} else {
			res.Header().Set("Content-Type", "application/json")
			res.Write(jsonCPUData)
		}
	}
	return CPUUsageHandler
}
