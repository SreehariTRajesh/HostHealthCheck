package webutils

import (
	"encoding/json"
	"goprober/host_utils"
	"log"
	"net/http"
)

type Memory_Usage struct {
	TotalMemory   uint64  `json: "TotalMemory"`
	UsedMemory    uint64  `json: "UsedMemory"`
	FreeMemory    uint64  `json: "FreeMemory`
	MemoryUsedPCT float64 `json: "MemoryUsedPCT"`
}

func Memory_Usage_API() func(http.ResponseWriter, *http.Request) {
	MemoryUsageHandler := func(res http.ResponseWriter, req *http.Request) {
		Host_Memory_Usage := host_utils.HostMemoryStats()
		memory_data := Memory_Usage{
			TotalMemory:   Host_Memory_Usage.Total,
			UsedMemory:    Host_Memory_Usage.Used,
			FreeMemory:    Host_Memory_Usage.Available,
			MemoryUsedPCT: Host_Memory_Usage.UsedPercent,
		}
		jsonMemoryData, err := json.Marshal(memory_data)
		if err != nil {
			log.Println("Error:", err)
		} else {
			res.Header().Set("Content-Type", "application/json")
			res.Write(jsonMemoryData)
		}
	}
	return MemoryUsageHandler
}
