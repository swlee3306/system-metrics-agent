package config

type MEMInfo struct {
	State string `json:"state"`
	TotalOnline string `json:"total_online"`
	TotalOffline string `json:"total_offline"`
	MemoryBlockSize string `json:"memory_block_size"`
}