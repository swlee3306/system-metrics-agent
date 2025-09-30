package alerting

import (
	"context"
	"fmt"
	"time"
)

type AlertLevel string

const (
	Info     AlertLevel = "info"
	Warning  AlertLevel = "warning"
	Critical AlertLevel = "critical"
)

type Alert struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Level       AlertLevel            `json:"level"`
	Metric      string                `json:"metric"`
	Threshold   float64               `json:"threshold"`
	Operator    string                 `json:"operator"` // "gt", "lt", "eq", "gte", "lte"
	Duration    time.Duration          `json:"duration"`
	Enabled     bool                  `json:"enabled"`
	Labels      map[string]string     `json:"labels"`
	Annotations map[string]string     `json:"annotations"`
	LastFired   *time.Time            `json:"last_fired,omitempty"`
	Firing      bool                  `json:"firing"`
}

type AlertRule struct {
	Alert
	CheckFunc func(value float64) bool
}

type AlertManager struct {
	rules    map[string]*AlertRule
	handlers []AlertHandler
	ctx      context.Context
	cancel   context.CancelFunc
}

type AlertHandler interface {
	HandleAlert(alert *Alert, value float64) error
}

type AlertManagerConfig struct {
	CheckInterval time.Duration
	MaxRetries    int
}

func NewAlertManager(config AlertManagerConfig) *AlertManager {
	ctx, cancel := context.WithCancel(context.Background())
	
	manager := &AlertManager{
		rules:    make(map[string]*AlertRule),
		handlers: make([]AlertHandler, 0),
		ctx:      ctx,
		cancel:   cancel,
	}
	
	go manager.startMonitoring(config.CheckInterval)
	
	return manager
}

func (am *AlertManager) AddRule(rule *AlertRule) {
	am.rules[rule.ID] = rule
}

func (am *AlertManager) RemoveRule(id string) {
	delete(am.rules, id)
}

func (am *AlertManager) AddHandler(handler AlertHandler) {
	am.handlers = append(am.handlers, handler)
}

func (am *AlertManager) startMonitoring(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for {
		select {
		case <-am.ctx.Done():
			return
		case <-ticker.C:
			am.checkAlerts()
		}
	}
}

func (am *AlertManager) checkAlerts() {
	for _, rule := range am.rules {
		if !rule.Enabled {
			continue
		}
		
		// Get current metric value (this would be implemented with actual metric collection)
		value := am.getMetricValue(rule.Metric)
		
		shouldFire := rule.CheckFunc(value)
		
		if shouldFire && !rule.Firing {
			// Alert is firing
			rule.Firing = true
			now := time.Now()
			rule.LastFired = &now
			
			// Notify handlers
			for _, handler := range am.handlers {
				handler.HandleAlert(&rule.Alert, value)
			}
		} else if !shouldFire && rule.Firing {
			// Alert is resolved
			rule.Firing = false
		}
	}
}

func (am *AlertManager) getMetricValue(metric string) float64 {
	// This would be implemented with actual metric collection
	// For now, return a mock value
	return 75.0
}

func (am *AlertManager) Close() {
	am.cancel()
}

// Predefined alert rules
func NewCPUAlert(threshold float64) *AlertRule {
	return &AlertRule{
		Alert: Alert{
			ID:          "cpu_high_usage",
			Name:        "High CPU Usage",
			Description: "CPU usage is above threshold",
			Level:       Critical,
			Metric:      "system_cpu_usage_percent",
			Threshold:   threshold,
			Operator:    "gt",
			Duration:    5 * time.Minute,
			Enabled:     true,
			Labels: map[string]string{
				"severity": "critical",
				"service":  "system",
			},
			Annotations: map[string]string{
				"summary":     "High CPU usage detected",
				"description": fmt.Sprintf("CPU usage is above %.1f%%", threshold),
			},
		},
		CheckFunc: func(value float64) bool {
			return value > threshold
		},
	}
}

func NewMemoryAlert(threshold float64) *AlertRule {
	return &AlertRule{
		Alert: Alert{
			ID:          "memory_high_usage",
			Name:        "High Memory Usage",
			Description: "Memory usage is above threshold",
			Level:       Warning,
			Metric:      "system_memory_usage_percent",
			Threshold:   threshold,
			Operator:    "gt",
			Duration:    5 * time.Minute,
			Enabled:     true,
			Labels: map[string]string{
				"severity": "warning",
				"service":  "system",
			},
			Annotations: map[string]string{
				"summary":     "High memory usage detected",
				"description": fmt.Sprintf("Memory usage is above %.1f%%", threshold),
			},
		},
		CheckFunc: func(value float64) bool {
			return value > threshold
		},
	}
}

func NewDiskAlert(threshold float64) *AlertRule {
	return &AlertRule{
		Alert: Alert{
			ID:          "disk_high_usage",
			Name:        "High Disk Usage",
			Description: "Disk usage is above threshold",
			Level:       Critical,
			Metric:      "system_disk_usage_percent",
			Threshold:   threshold,
			Operator:    "gt",
			Duration:    5 * time.Minute,
			Enabled:     true,
			Labels: map[string]string{
				"severity": "critical",
				"service":  "system",
			},
			Annotations: map[string]string{
				"summary":     "High disk usage detected",
				"description": fmt.Sprintf("Disk usage is above %.1f%%", threshold),
			},
		},
		CheckFunc: func(value float64) bool {
			return value > threshold
		},
	}
}
