package cpu

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"system-Info-collector/pkg/config"
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
