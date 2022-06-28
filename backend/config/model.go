package config

type CFG struct {
	Version      string `json:"version"`      // 版本
	DemoDir      string `json:"demoDir"`      // 下载的文件夹
	UseExternel  bool   `json:"useExternel"`  // 外置跳转下载
	AutoDownload bool   `json:"autoDownload"` // 自动下载
}
