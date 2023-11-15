package info

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
	"time"
)

type CpuInfo struct {
	ModelName string  `json:"model_name"`
	CacheSize int32   `json:"cache_size"`
	MaxHz     float64 `json:"max_hz"`
	Percent   float64 `json:"used_percent"`
	Counts    int     `json:"counts"`
	Process   int32   `json:"precess"`
	Threads   int32   `json:"threads"`
}

var cpuInfo CpuInfo

func init() {
	info, _ := cpu.Info()
	cpuInfo = CpuInfo{}

	cpuInfo.ModelName = info[0].ModelName
	cpuInfo.CacheSize = info[0].CacheSize * 1024 // 统一单位为字节
	cpuInfo.MaxHz = info[0].Mhz

	go updateCpuInfo()
}

func updateCpuInfo() {
	for {
		counts, _ := cpu.Counts(true)
		percent, _ := cpu.Percent(time.Second, false)

		cpuInfo.Counts = counts
		cpuInfo.Percent = percent[0]

		cpuInfo.Process = countProcessAndThreads()["process"]
		cpuInfo.Threads = countProcessAndThreads()["threads"]

		time.Sleep(1000 * time.Millisecond)
	}
}

// getCpuInfo 获取CPU信息
func getCpuInfo() *CpuInfo {
	return &cpuInfo
}

func countProcessAndThreads() map[string]int32 {
	// 获取所有进程的信息
	processes, _ := process.Processes()

	// 统计进程数和线程数
	var processCount int32
	var threadsCount int32

	processCount = int32(len(processes))

	// 遍历每个进程，累加线程数
	for _, proc := range processes {
		threadCountPerProcess, err := proc.NumThreads()
		if err == nil {
			threadsCount += threadCountPerProcess
		}
	}

	result := make(map[string]int32)
	result["process"] = processCount
	result["threads"] = threadsCount

	return result
}
