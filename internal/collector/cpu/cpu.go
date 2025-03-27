package cpu

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"system-Info-collector/pkg/config"
	"time"
)

// GetCpuInfo retrieves CPU information using the `lscpu` command
func GetCpuInfo() ([]*config.CPUInfo, error) {
	cmd := exec.Command("lscpu")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	cpus := make([]*config.CPUInfo, 0)
	var cpu config.CPUInfo

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Architecture":
			cpu.Processor = value
			if value == "aarch64" {
				cpu.ModelName = "ARM Processor"
			}
		case "Model name":
			cpu.ModelName = value
		case "CPU(s)":
			cpu.Cpus = value
		case "Vendor ID":
			cpu.Vendor_ID = value
		case "Thread(s) per core":
			cpu.Threads_per_core = value
		case "Core(s) per socket":
			cpu.Cores_per_socket = value
		case "Socket(s)":
			cpu.Socket = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// 결과를 리스트에 추가
	cpus = append(cpus, &cpu)

	return cpus, nil
}

// getCpuTimes reads the CPU times from /proc/stat
func getCpuTimes() (idle, total float64, err error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 8 {
			return 0, 0, fmt.Errorf("unexpected format in /proc/stat")
		}

		var values []float64
		for _, field := range fields[1:8] { // 첫 번째 필드는 "cpu" 라벨이므로 제외
			v, err := strconv.ParseFloat(field, 64)
			if err != nil {
				return 0, 0, err
			}
			values = append(values, v)
		}

		idle = values[3] // idle time
		total = 0
		for _, v := range values {
			total += v
		}
	}

	return idle, total, scanner.Err()
}

// GetCpuUsage calculates CPU usage by sampling over a short time period
func GetCpuUsage(interval time.Duration) (*config.CpuUsage, error) {
	idle1, total1, err := getCpuTimes()
	if err != nil {
		return nil, err
	}

	time.Sleep(interval)

	idle2, total2, err := getCpuTimes()
	if err != nil {
		return nil, err
	}

	idleDelta := idle2 - idle1
	totalDelta := total2 - total1
	usage := &config.CpuUsage{}

	if totalDelta > 0 {
		usage.Idle = (idleDelta / totalDelta) * 100
		usage.Total = 100 - usage.Idle
	}

	return usage, nil
}
