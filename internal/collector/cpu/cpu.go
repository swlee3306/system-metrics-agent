package cpu

import (
	"bufio"
	"os"
	"strings"
	"system-Info-collector/pkg/config"
)

func GetCpuInfo() ([]*config.CPUInfo, error) {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cpus := make([]*config.CPUInfo, 0)

	var cpu config.CPUInfo

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" { //빈 줄이 나오면 CPU 정보가 끝난 것으로 간주
			cpus = append(cpus, &cpu)
			cpu = config.CPUInfo{}
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "processor":
			cpu.Processor = value
		case "model name":
			cpu.ModelName = value
		case "cpu core":
			cpu.Cores = value
		case "cpu MHz":
			cpu.CPU_MHz = value
		case "cache size":
			cpu.CacheSize = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cpus, nil
}
