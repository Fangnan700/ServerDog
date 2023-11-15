package info

import (
	"ServerDog/utils"
	"github.com/shirou/gopsutil/host"
	"strconv"
	"time"
)

type HostInfo struct {
	Hostname        string `json:"host_name"`
	IP              string `json:"ip"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platform_family"`
	PlatformVersion string `json:"platform_version"`
	KernelArch      string `json:"kernel_arch"`
	KernelVersion   string `json:"kernel_version"`
	Location        string `json:"location"`
	Provider        string `json:"provider"`
	LastLogin       string `json:"last_login"`
	RunningTime     string `json:"running_time"`
}

var hostInfo HostInfo

func init() {
	info, _ := host.Info()

	hostInfo = HostInfo{}
	hostInfo.Hostname = info.Hostname
	hostInfo.OS = info.OS
	hostInfo.Platform = info.Platform
	hostInfo.PlatformFamily = info.PlatformFamily
	hostInfo.PlatformVersion = info.PlatformVersion
	hostInfo.KernelVersion = info.KernelVersion
	hostInfo.KernelArch = info.KernelArch
	hostInfo.LastLogin = utils.GetLastLogin()
}

// getHostInfo 获取主机信息
func getHostInfo() *HostInfo {
	info, _ := host.Info()
	bootTime := time.Unix(int64(info.BootTime), 0)
	uptime := time.Since(bootTime)

	hostInfo.RunningTime = strconv.FormatInt(int64(uptime.Seconds()), 10)
	return &hostInfo
}
