package info

type SystemInfo struct {
	Cpu  CpuInfo           `json:"cpu"`
	Mem  VirtualMemoryInfo `json:"mem"`
	Net  NetInfo           `json:"net"`
	Swap SwapMemoryInfo    `json:"swap"`
	Disk DiskInfo          `json:"disk"`
	Host HostInfo          `json:"host"`
}

var systemInfo SystemInfo

func init() {
	systemInfo = SystemInfo{}
}

func GetSystemInfo() *SystemInfo {
	systemInfo.Cpu = *getCpuInfo()
	systemInfo.Mem = *getVirtualMemoryInfo()
	systemInfo.Net = *getNetInfo()
	systemInfo.Swap = *getSwapMemoryInfo()
	systemInfo.Disk = *getDiskInfo()
	systemInfo.Host = *getHostInfo()
	return &systemInfo
}
