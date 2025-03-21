package network

import (
	"bytes"
	"os/exec"
	"strings"
	"system-Info-collector/pkg/config"
)

// 네트워크 정보를 가져오는 함수
func GetNetworkInfo() ([]*config.NetworkInterfaceInfo, error) {
	script := `
		#!/bin/bash

		# 물리 인터페이스 목록 가져오기 (Loopback, 가상 인터페이스 제외)
		interfaces=()
		for iface in $(ls /sys/class/net/); do
		    if [[ -e "/sys/class/net/$iface/device" ]] && [[ "$iface" != "lo" ]]; then
		        interfaces+=("$iface")
		    fi
		done

		# 본딩 인터페이스 확인 (bonding 정보 저장할 배열)
		declare -A bonded_macs
		declare -A bond_interfaces
		declare -A bond_mac_addrs  # 본딩 인터페이스의 MAC 주소 저장

		# 본딩 인터페이스 목록 탐색 (존재할 때만)
		if [[ -d /proc/net/bonding ]]; then
			for bond in /proc/net/bonding/*; do
			    bond_name=$(basename "$bond")
			    bond_mac=$(cat /sys/class/net/$bond_name/address 2>/dev/null)
			    bond_mac_addrs[$bond_name]=$bond_mac
			    
			    while read -r line; do
			        if [[ $line == "Slave Interface:"* ]]; then
			            slave_iface=$(echo "$line" | awk '{print $3}')
			            bond_interfaces[$slave_iface]=$bond_name
			        elif [[ $line == "Permanent HW addr:"* ]]; then
			            bonded_mac=$(echo "$line" | awk '{print $4}')
			            bonded_macs[$slave_iface]=$bonded_mac
			        fi
			    done < "$bond"
			done
		fi

		# 인터페이스별 MAC 주소 출력
		for iface in "${interfaces[@]}"; do
		    if [[ -n "${bond_interfaces[$iface]}" ]]; then
		        bond_name=${bond_interfaces[$iface]}
		        echo "$iface|${bonded_macs[$iface]}|$bond_name|${bond_mac_addrs[$bond_name]}"
		    elif [[ -n "${bond_mac_addrs[$iface]}" ]]; then
		        echo "$iface|${bond_mac_addrs[$iface]}|$iface|${bond_mac_addrs[$iface]}"
		    else
		        mac=$(cat /sys/class/net/$iface/address 2>/dev/null)
		        echo "$iface|${mac}| | "
		    fi
		done
	`

	cmd := exec.Command("bash", "-c", script)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	result := strings.TrimSpace(out.String())
	lines := strings.Split(result, "\n")

	var networkInterfaces []*config.NetworkInterfaceInfo

	// 결과를 구조체에 매핑
	for _, line := range lines {
		fields := strings.Split(line, "|")
		if len(fields) < 4 {
			continue
		}

		networkInterface := &config.NetworkInterfaceInfo{
			Interfaces: []struct {
				Name         string `json:"name"`          // 인터페이스 이름
				MacAddress   string `json:"mac_address"`   // MAC 주소
				BondingGroup string `json:"bonding_group"` // 본딩 그룹명 (없으면 빈 문자열)
				BondingMac   string `json:"bonding_mac"`   // 본딩 MAC 주소 (없으면 빈 문자열)
			}{
				{
					Name:         strings.TrimSpace(fields[0]),
					MacAddress:   strings.TrimSpace(fields[1]),
					BondingGroup: strings.TrimSpace(fields[2]),
					BondingMac:   strings.TrimSpace(fields[3]),
				},
			},
		}
		networkInterfaces = append(networkInterfaces, networkInterface)
	}

	return networkInterfaces, nil
}
