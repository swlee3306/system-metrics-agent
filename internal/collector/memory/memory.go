package memory

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"system-Info-collector/pkg/config"
)

func GetMemoryInfo() ([]*config.MEMInfo, error) {
	cmd := exec.Command("lsmem")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	mems := make([]*config.MEMInfo, 0)
	mem := &config.MEMInfo{}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	foundState := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// 첫 번째 테이블에서 STATE 값을 추출
		fields := strings.Fields(line)
		if len(fields) >= 4 && (fields[2] == "online" || fields[2] == "offline") {
			mem.State = fields[2]
			foundState = true
			continue
		}

		// `:` 구분자로 된 키-값 형식에서 나머지 정보 추출
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Memory block size":
			mem.MemoryBlockSize = value
		case "Total online memory":
			mem.TotalOnline = value
		case "Total offline memory":
			mem.TotalOffline = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// STATE 값이 발견되지 않은 경우, 기본값을 할당
	if !foundState {
		mem.State = "unknown"
	}

	// mem 객체를 리스트에 추가
	mems = append(mems, mem)
	return mems, nil
}
