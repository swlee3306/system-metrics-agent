package exporter

import (
	"fmt"
	"log"
	"system-Info-collector/internal/collector/cpu"
	"system-Info-collector/internal/collector/disk"
	"system-Info-collector/internal/collector/memory"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	CPUUsageGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage_percentage",
			Help: "Current CPU usage percentage",
		},
		[]string{"processor", "model_name", "cpus", "vendor_id", "threads_per_core", "cores_per_socket", "socket"},
	)

	MemoryUsageGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "memory_usage_percentage",
			Help: "Current Memory usage percentage",
		},
		[]string{"state", "total_online", "total_offline", "memory_block_size"},
	)

	DiskUsageGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "disk_usage_percentage",
			Help: "Current Disk usage percentage",
		},
		[]string{"filesystem", "size", "used", "avail", "use_percent", "mounted_on"},
	)
)

// MetricCollector collects CPU, Memory, and Disk usage periodically
func MetricExporter(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		// Collect CPU info and usage
		cpuInfo, err := cpu.GetCpuInfo()
		if err != nil {
			log.Println("Error collecting CPU info:", err)
		}
		cpuUsage, err := cpu.GetCpuUsage(interval)
		if err == nil && cpuInfo != nil && cpuUsage != nil {
			for _, cpu := range cpuInfo {
				CPUUsageGauge.WithLabelValues(
					cpu.Processor,
					cpu.ModelName,
					cpu.Cpus,
					cpu.Vendor_ID,
					cpu.Threads_per_core,
					cpu.Cores_per_socket,
					cpu.Socket,
				).Set(cpuUsage.Total)
			}
			fmt.Printf("CPU Usage Collected: %.2f%%\n", cpuUsage.Total)
		} else {
			log.Println("Error collecting CPU usage:", err)
		}

		// Collect Memory info and usage
		memInfo, err := memory.GetMemoryInfo()
		if err != nil {
			log.Println("Error collecting Memory info:", err)
		}
		memUsage, err := memory.GetMemoryUsage()
		if err == nil && memInfo != nil && memUsage != nil {
			for _, mem := range memInfo {
				MemoryUsageGauge.WithLabelValues(
					mem.State,
					mem.TotalOnline,
					mem.TotalOffline,
					mem.MemoryBlockSize,
				).Set(memUsage.Usage)
			}
			fmt.Printf("Memory Usage Collected: %.2f%%\n", memUsage.Usage)
		} else {
			log.Println("Error collecting Memory usage:", err)
		}

		// Collect Disk info and usage
		diskInfo, err := disk.GetDiskInfo()
		if err != nil {
			log.Println("Error collecting Disk info:", err)
		}
		diskUsage, err := disk.GetDiskUsage("/")
		if err == nil && diskInfo != nil && diskUsage != nil {
			for _, disk := range diskInfo {
				for _, fs := range disk.Filesystems {
					DiskUsageGauge.WithLabelValues(
						fs.Filesystem,
						fs.Size,
						fs.Used,
						fs.Avail,
						fs.Use,
						fs.MountedOn,
					).Set(diskUsage.Usage)
				}
			}
			fmt.Printf("Disk Usage Collected: %.2f%%\n", diskUsage.Usage)
		} else {
			log.Println("Error collecting Disk usage:", err)
		}
	}
}
