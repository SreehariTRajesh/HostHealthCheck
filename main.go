package main

import (
	webutils "goprober/web_utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cpu", webutils.CPU_Usage_API())
	http.HandleFunc("/mem", webutils.Memory_Usage_API())
	http.HandleFunc("/disk", webutils.Disk_Usage_API())
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Server up and running on http://localhost:8080")
}
