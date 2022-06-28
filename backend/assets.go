package backend

import (
	"embed"
	"io/ioutil"
	"os"
	"path"

	ptools "github.com/One-Studio/ptools/pkg"
	log "github.com/sirupsen/logrus"
)

// 处理打包的附件
//go:embed public
var assets embed.FS

var files = []string{
	"steam_api.dll",
	"steam_appid.txt",
	"boiler-writter.exe",
}

func ReleaseAssets(targetDir string, embedDir string, assets embed.FS, files []string) error {
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		log.Println(err)
		return err
	}

	for _, v := range files {
		p := path.Join(targetDir, v)

		// 文件已存在则跳过 TODO: 检测文件大小，避免更新附件失败
		if ptools.IsFileExisted(p) {
			// 读取附件大小
			f, err := assets.Open(path.Join(embedDir, v))
			if err != nil {
				log.Info(err)
			}

			stat, err := f.Stat()
			if err != nil {
				log.Info(err)
			}

			stat.Size()

			// 读取已有文件大小
			file, err := os.Open(p)
			if err != nil {
				log.Info(err)
			}
			fi, err := file.Stat()
			if err != nil {
				log.Info(err)
			}

			if stat.Size() == fi.Size() {
				continue
				log.Debug("跳过", p)
			}

		}

		// 读附件
		data, err := assets.ReadFile(embedDir + v)
		if err != nil {
			log.Println(err)
			return err
		}

		// 写文件
		if err := ioutil.WriteFile(p, data, 0666); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (a *App) releaseAssets() error {
	return ReleaseAssets(a.GetConfigPath(""), "public/", assets, files)
}
