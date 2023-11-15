package info

import (
	"github.com/shirou/gopsutil/mem"
	"time"
)

type VirtualMemoryInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Free        uint64  `json:"free"`
	Cached      uint64  `json:"cached"`
	UsedPercent float64 `json:"used_percent"`
}

var memoryInfo VirtualMemoryInfo

func init() {
	memoryInfo = VirtualMemoryInfo{}
	go updateMemInfo()
}

func updateMemInfo() {
	for {
		info, _ := mem.VirtualMemory()
		memoryInfo.Total = info.Total
		memoryInfo.Available = info.Available
		memoryInfo.Free = info.Free
		memoryInfo.Cached = info.Cached
		memoryInfo.UsedPercent = info.UsedPercent

		time.Sleep(1000 * time.Millisecond)
	}
}

// getVirtualMemoryInfo 获取主存信息
func getVirtualMemoryInfo() *VirtualMemoryInfo {
	return &memoryInfo
}
