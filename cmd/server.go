package cmd

import (
	"system-Info-collector/internal/collector/exporter"
	"system-Info-collector/pkg/web"

	"time"
)

func StartSystemAgent() {
	// ✅ Start metric exporter synchronously first
	go exporter.MetricExporter(time.Duration(10) * time.Second)

	// ✅ Start web server last
	web.StartWebServer()
}
