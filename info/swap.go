package info

import (
	"github.com/shirou/gopsutil/mem"
	"time"
)

type SwapMemoryInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

var swapInfo SwapMemoryInfo

func init() {
	swapInfo = SwapMemoryInfo{}
	go updateSwapInfo()
}

func updateSwapInfo() {
	for {
		info, _ := mem.SwapMemory()

		swapInfo.Total = info.Total
		swapInfo.Used = info.Used
		swapInfo.Free = info.Free
		swapInfo.UsedPercent = info.UsedPercent

		time.Sleep(1000 * time.Millisecond)
	}
}

// getSwapMemoryInfo 获取交换分区信息
func getSwapMemoryInfo() *SwapMemoryInfo {
	return &swapInfo
}
