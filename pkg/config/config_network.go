package config

// 네트워크 인터페이스 정보를 저장할 구조체
type NetworkInterfaceInfo struct {
	Interfaces []struct {
		Name         string `json:"name"`          // 인터페이스 이름
		IpAddress   string `json:"ip_address"`    // IP 주소
		MacAddress   string `json:"mac_address"`   // MAC 주소
		BondingGroup string `json:"bonding_group"` // 본딩 그룹명 (없으면 빈 문자열)
		BondingMac   string `json:"bonding_mac"`   // 본딩 MAC 주소 (없으면 빈 문자열)
	} `json:"NetworkInterfaces"`
}
