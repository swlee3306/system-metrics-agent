package config

type CPUInfo struct {
	Processor string `json:"processor"`
	ModelName string `json:"model_name"`
	Cpus string `json:"cpus"`
	Vendor_ID   string `json:"Vendor ID"`
	Threads_per_core string `json:"threads_per_core"`
	Cores_per_socket string `json:"cores_per_socket"`
	Socket string `json:"socket"`
}

// CpuUsage represents CPU usage statistics
type CpuUsage struct {
	User   float64
	System float64
	Idle   float64
	Total  float64
}
