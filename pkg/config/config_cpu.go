package config

type CPUInfo struct {
	Processor string `json:"processor"`
	ModelName string `json:"model_name"`
	Cores     string `json:"cores"`
	CPU_MHz   string `json:"cpu_mhz"`
	CacheSize string `json:"cache_size"`
}
