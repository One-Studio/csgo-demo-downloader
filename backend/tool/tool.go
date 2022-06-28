package tool

import (
	"errors"
	ptools "github.com/One-Studio/ptools/pkg"
	"strings"
)

func FormatPath(s string) string {
	return strings.Replace(s, "\\", "/", -1)
}

func KillProcess(name string) error {
	_, err := ptools.Exec("taskkill /im " + name + ".exe /f")
	if err != nil {
		if err.Error() == "exit status 128" {
			return errors.New(name + "已停止运行")
		}
		return err
	}

	return nil
}
