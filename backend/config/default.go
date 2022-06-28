package config

var defaultCFG = CFG{
	Version:      "",
	DemoDir:      "",
	UseExternel:  false,
	AutoDownload: false,
}

func DefaultCFG() CFG {
	return defaultCFG
}
