package replay

type ShareCodeStruct struct {
	MatchId   string
	OutcomeId string
	TokenId   string
}

type Demo struct {
	Name     string `json:"name"`     // demo名称
	Filename string `json:"filename"` // demo文件名
	Md5      string `json:"md5"`      // demo的md5值 用于区分
}

type HlaeLaunchOption struct {
	Override        bool     `json:"override"`        //
	MmcfgEnabled    bool     `json:"mmcfgEnabled"`    //
	MmcfgDir        string   `json:"mmcfgDir"`        //
	LaunchOption    string   `json:"launchOption"`    //
	GfxEnabled      bool     `json:"gfxEnabled"`      //
	GfxWidth        int      `json:"gfxWidth"`        //
	GfxHeight       int      `json:"gfxHeight"`       //
	GfxFull         bool     `json:"gfxFull"`         //
	UseCustomLoader bool     `json:"useCustomLoader"` //
	HookDllPaths    []string `json:"hookDllPaths"`    //
	Envs            []string `json:"envs"`            //
	Inited          bool     `json:"inited"`          //
}
