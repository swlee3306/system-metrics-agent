package config

type DiskInfo struct {
	Filesystems []struct {
		Filesystem string `json:"filesystem"`
		Size       string `json:"size"`
		Used       string `json:"used"`
		Avail      string `json:"avail"`
		Use        string `json:"use%"`
		MountedOn  string `json:"mounted_on"`
	} `json:"Filesystems"`
}

type DiskUsage struct {
	Total float64
	Used  float64
	Free  float64
	Usage float64
}
