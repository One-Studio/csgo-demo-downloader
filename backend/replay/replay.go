package replay

import (
	"downloader/backend/tool"
	"errors"
	"os"
	"path"

	ptools "github.com/One-Studio/ptools/pkg"
)

func ListDemo(demoDir string) ([]Demo, error) {
	if demoDir == "" {
		return nil, errors.New("录像下载路径为空")
	}

	p := tool.FormatPath(demoDir)
	suffix := []string{".dem", ".zip", ".bz2"}

	if !ptools.IsFileExisted(p) {
		if err := os.MkdirAll(p, os.ModePerm); err != nil {
			return nil, errors.New("录像下载路径不存在: " + p)
		}
	}

	tFiles, err := ptools.ListDir(p, suffix)
	if err != nil {
		return nil, errors.New("list dir failed. " + err.Error())
	}

	var files []Demo
	for _, file := range tFiles {
		// 处理每个文件
		name := path.Base(tool.FormatPath(file))
		filename := path.Base(tool.FormatPath(file))
		md5 := ""
		files = append(files, Demo{name, filename, md5})
	}

	return files, nil
}
