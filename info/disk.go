package info

import (
	"github.com/shirou/gopsutil/disk"
	"time"
)

type DiskInfo struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	ReadBytes   uint64  `json:"read_bytes"`
	WriteBytes  uint64  `json:"write_bytes"`
}

var diskInfo DiskInfo

func init() {
	diskInfo = DiskInfo{}
	go updateDiskInfo()
}

func updateDiskInfo() {
	for {
		parts, _ := disk.Partitions(true)

		var (
			total uint64 = 0
			free  uint64 = 0
			used  uint64 = 0
		)
		for _, part := range parts {
			usage, _ := disk.Usage(part.Mountpoint)
			total += usage.Total
			free += usage.Free
			used += usage.Used
		}

		diskInfo.Total = total
		diskInfo.Free = free
		diskInfo.Used = used
		diskInfo.UsedPercent = float64(used) / float64(total)

		ioCounters, _ := disk.IOCounters()
		for _, v := range ioCounters {
			diskInfo.ReadBytes += v.ReadBytes
			diskInfo.WriteBytes += v.WriteBytes
		}

		time.Sleep(5 * time.Minute)
	}
}

// getDiskInfo 获取磁盘信息
func getDiskInfo() *DiskInfo {
	return &diskInfo
}
