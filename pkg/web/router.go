package web

import (
	"fmt"
	"net/http"
	"system-Info-collector/internal/collector/cpu"
	"system-Info-collector/internal/collector/disk"
	"system-Info-collector/internal/collector/exporter"
	"system-Info-collector/internal/collector/memory"
	"system-Info-collector/internal/collector/network"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	fnc := "StartWebServer"
	fmt.Println("Starting web server:", fnc)

	prometheus.MustRegister(exporter.CPUUsageGauge)
	prometheus.MustRegister(exporter.MemoryUsageGauge)
	prometheus.MustRegister(exporter.DiskUsageGauge)

	// Add /metrics and /health handlers to the default HTTP server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"ok"}`))
		})
		if err := http.ListenAndServe(":9091", nil); err != nil {
			fmt.Printf("Metric server error: %v\n", err)
		}
	}()

	r := gin.Default()
	r.Use(gin.Recovery())

	apiGroup := r.Group("/api/v1.0")
	{
		apiGroup.GET("/cpuinfo", func(c *gin.Context) {
			cpuInfos, err := cpu.GetCpuInfo()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to get CPU info"})
				return
			}
			c.JSON(200, cpuInfos)
		})

		apiGroup.GET("/memoryinfo", func(c *gin.Context) {
			memoryInfos, err := memory.GetMemoryInfo()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to get memory info"})
				return
			}
			c.JSON(200, memoryInfos)
		})

		apiGroup.GET("/diskinfo", func(c *gin.Context) {
			diskInfos, err := disk.GetDiskInfo()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to get disk info"})
				return
			}
			c.JSON(200, diskInfos)
		})

		apiGroup.GET("/networkinfo", func(c *gin.Context) {
			networkInfos, err := network.GetNetworkInfo()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to get network info"})
				return
			}
			c.JSON(200, networkInfos)
		})
	}
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Web server error: %v\n", err)
	}
}
