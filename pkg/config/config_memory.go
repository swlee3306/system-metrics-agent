package config

type MEMInfo struct {
	State           string `json:"state"`
	TotalOnline     string `json:"total_online"`
	TotalOffline    string `json:"total_offline"`
	MemoryBlockSize string `json:"memory_block_size"`
}

type MemoryUsage struct {
	Total     float64
	Used      float64
	Available float64
	Usage     float64
}
