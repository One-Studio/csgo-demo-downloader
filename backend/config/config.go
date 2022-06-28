package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	// "log"

	"os"
	"path"
	"strings"

	ptools "github.com/One-Studio/ptools/pkg"
	tool "github.com/One-Studio/ptools/pkg"
	log "github.com/sirupsen/logrus"
)

// 读设置
func ReadConfig(path string) (CFG, error) {
	// 检查文件是否存在
	if exist := tool.IsFileExisted(path); exist == true {
		// 存在则读取文件
		log.Println(path)
		cont, err := ptools.ReadAll(path)

		// 初始化实例并解析JSON
		var CFGInst CFG
		err = json.Unmarshal([]byte(cont), &CFGInst) // 第二个参数要地址传递
		if err != nil {
			log.Errorf("[%T]%s, %d", err, err, 53)
			return CFG{}, err
		}

		return CFGInst, nil
	} else {
		// 设置文件不存在则初始化
		return defaultCFG, nil
	}
}

// 写设置
func SaveConfig(cfg CFG, pakPath string) error {
	// 规格化路径
	pakPath = strings.Replace(pakPath, "\\", "/", -1)
	// 检查文件是否存在
	dir := path.Dir(pakPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Error(err)
	}

	JsonData, err := json.Marshal(cfg)
	if err != nil {
		log.Errorf("[%T]%s, %d", err, err, 77)
		return err
	}

	// json写入文件
	var str bytes.Buffer
	_ = json.Indent(&str, JsonData, "", "") // "", "    "

	if err := ioutil.WriteFile(pakPath, str.Bytes(), 0666); err != nil {
		log.Errorf("[%T]%s, %d", err, err, 96)
		return err
	}

	return nil
}

// 设置转Json字符串
func Config2Json(cfg CFG) (string, error) {
	JsonData, err := json.Marshal(cfg) // 第二个参数要地址传递
	if err != nil {
		return "", err
	}

	// json写入文件
	var str bytes.Buffer
	_ = json.Indent(&str, JsonData, "", "    ")
	return str.String(), nil
}

// 检查设置不一致得到null|nil的情况
func CheckConfig(cfg CFG) CFG {

	return cfg
}
