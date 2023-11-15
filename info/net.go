package info

import (
	"github.com/shirou/gopsutil/net"
	"time"
)

type NetInfo struct {
	UploadSpeed   float64 `json:"upload_speed"`
	DownloadSpeed float64 `json:"download_speed"`
}

var netInfo NetInfo

func init() {
	netInfo = NetInfo{}
	go updateNetInfo()
}

func updateNetInfo() {
	var (
		uploadBytes       uint64
		downloadBytes     uint64
		lastUploadBytes   uint64
		lastDownloadBytes uint64
	)
	for {
		uploadBytes = 0
		downloadBytes = 0

		info, _ := net.IOCounters(true)
		for _, v := range info {
			uploadBytes += v.BytesSent
			downloadBytes += v.BytesRecv
		}

		if lastUploadBytes == 0 || lastDownloadBytes == 0 {
			lastUploadBytes = uploadBytes
			lastDownloadBytes = downloadBytes
		}

		netInfo.UploadSpeed = (float64(uploadBytes) - float64(lastUploadBytes)) / 1
		netInfo.DownloadSpeed = (float64(downloadBytes) - float64(lastDownloadBytes)) / 1

		lastUploadBytes = uploadBytes
		lastDownloadBytes = downloadBytes

		time.Sleep(1000 * time.Millisecond)
	}
}

// getNetInfo 获取网络信息
func getNetInfo() *NetInfo {
	return &netInfo
}
