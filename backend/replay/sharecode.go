package replay

import (
	"downloader/backend/tool"
	"math/big"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	ptools "github.com/One-Studio/ptools/pkg"
	"github.com/pkg/browser"
)

func (p *ShareCodeStruct) init(MatchId string, OutcomeId string, TokenId string) {
	p.MatchId = MatchId
	p.OutcomeId = OutcomeId
	p.TokenId = TokenId
}

func PureShareCode(shareCode string) (shareCodes []string) {
	reg := regexp.MustCompile(`CSGO-(.*)`)
	codes := reg.FindAllStringSubmatch(shareCode, -1)
	if codes == nil {
		return []string{}
	}

	for _, code := range codes {
		shareCodes = append(shareCodes, strings.ReplaceAll(code[1], "-", ""))
	}

	return
}

/**
 * @Author 黄鱼
 * @解析shareCode得到数组
 * @Date 2021/10/24 | 2022/4/26 Purp1e修改
 * @Param string shareCode
 * @return ShareCodeStruct
 **/
func ShareCodeDecode(shareCode string) (string, string, string) {
	bigNumber := big.NewInt(0)
	DICTIONARY := "ABCDEFGHJKLMNOPQRSTUVWXYZabcdefhijkmnopqrstuvwxyz23456789"
	DICTIONARYLength := big.NewInt(int64(len(DICTIONARY)))
	for i := len(shareCode) - 1; i >= 0; i-- {
		bigNumber = bigNumber.Add(bigNumber.Mul(bigNumber, DICTIONARYLength), big.NewInt(int64(strings.Index(DICTIONARY, string(shareCode[i])))))
	}
	bin := ""
	number := bigNumber
	for {
		if number.String() == "0" {
			break
		}
		mod := big.NewInt(2)
		mod.Mod(number, mod)
		if mod.String() == "1" {
			number.Sub(number, big.NewInt(1))
		}
		bin += mod.String()
		number.Div(number, big.NewInt(2))
	}
	var byteArray []string
	for i := 0; i < len(bin); i = i + 8 {
		startIndex := i + 8
		if startIndex >= len(bin) {
			startIndex = len(bin)
		}
		s := bin[i:startIndex]
		sre := ""
		for i := len(s) - 1; i >= 0; i-- {
			sre += string(s[i])
		}
		for {
			if len(sre) == 8 {
				break
			}
			sre = "0" + sre
		}
		byteArray = append(byteArray, sre)
	}
	// if len(byteArray) != 19{
	//	byteArray = append(byteArray, "00000000")
	// }
	MatchId := ""
	OutcomeId := ""
	TokenId := ""
	for i, j := len(byteArray)-8, 0; j < 8; i, j = i+1, j+1 {
		MatchId += byteArray[i]
	}
	MatchIdInt, _ := new(big.Int).SetString(MatchId, 2)
	for i, j := len(byteArray)-16, 0; j < 8; i, j = i+1, j+1 {
		OutcomeId += byteArray[i]
	}
	OutcomeIdInt, _ := new(big.Int).SetString(OutcomeId, 2)
	for i := 0; i < len(byteArray)-16; i++ {
		TokenId += byteArray[i]
	}
	TokenIdInt, _ := new(big.Int).SetString(TokenId, 2)
	shareCodeStruct := new(ShareCodeStruct)
	shareCodeStruct.init(MatchIdInt.String(), OutcomeIdInt.String(), TokenIdInt.String())
	return MatchIdInt.String(), OutcomeIdInt.String(), TokenIdInt.String()
}

func GetDemoLink(shareLink string, toolPath string) (urls []string) {
	// 确保工具存在
	toolPath = tool.FormatPath(toolPath)
	converter := path.Join(toolPath, "boiler-writter.exe")
	if !ptools.IsFileExisted(converter) {
		log.Println("converter not found")
		return
	}

	// 正则表达提取若干分享代码纯享版 -> TTVy4mt9mLn8VFabeamsoCi5P, ...,  | steam://rungame/730/76561202255233023/+csgo_download_match%20CSGO-TTVy4-mt9mL-n8VFa-beams-oCi5P
	codes := PureShareCode(shareLink)
	if len(codes) < 1 {
		log.Info("无有效分享代码")
		return []string{}
	}

	for _, code := range codes {
		// 确保Steam开启，CSGO关闭
		for i := 0; i < 3; i++ {
			command := "wmic process where name='steam.exe' get executablepath /value"
			_, err := ptools.CMD(command)
			if err != nil {
				log.Println(err)
				// 未检测到时启动csgo
				if err = browser.OpenURL("steam://rungame"); err != nil {
					log.Println(err)
				}
				time.Sleep(time.Second * 2)
			} else {
				break
			}
		}

		// 检测csgo进程并强制关闭
		command := "wmic process where name='csgo.exe' get executablepath /value"
		if _, err := ptools.CMD(command); err != nil {
			log.Println(err)
		} else if err := tool.KillProcess("csgo"); err != nil {
			log.Println(err)
		}

		mId, oId, tvId := ShareCodeDecode(code)
		tPath := path.Join(toolPath, "link_"+code+".txt")

		// 调用Converter
		out, err := ptools.ExecArgs([]string{converter, tPath, mId, oId, tvId})
		if err != nil {
			log.Println(err, out)
			switch out {
			case "-1":
				out = "error code -1: invalid params passing to converter"
			case "1":
				out = "error code 1: FatalError"
			case "2":
				out = "error code 2: InvalidArgs"
			case "3":
				out = "error code 3: CommunicationFailure"
			case "4":
				out = "error code 4: AlreadyConnectedToGC"
			case "5":
				out = "error code 5: SteamRestartRequired"
			case "6":
				out = "error code 6: SteamNotRunningOrLoggedIn"
			case "7":
				out = "error code 7: SteamUserNotLoggedIn"
			case "8":
				out = "error code 8: NoMatchesFound"
			case "9":
				out = "error code 9: WriteFileFailure"
			}
			log.Println(out)
		} else {
			// 读取URL
			url, err := ptools.ReadAll(tPath)
			if err != nil {
				log.Println(err, url)
			} else {
				urls = append(urls, url)
			}

			// 清理临时txt文件
			if err := os.Remove(tPath); err != nil {
				log.Println(err)
			}
		}
	}

	// 备用代码，先暂时用Converter
	// // 解析demo分享链接
	// // 获取Code Ids
	// code := replay.ShareCodeDecode(shareCode)
	// println(code.MatchId, code.OutcomeId, code.TokenId)
	// // 运行boiler
	// boiler := strings.ReplaceAll(a.getConfigPath("boiler-writter.exe"), "\\", "/")
	// out, err := ExecWithSpace(strconv.Quote(boiler) + " text.txt " + code.MatchId + " " + code.OutcomeId + " " + code.TokenId)
	// if err != nil {
	// 	log.Println(err)
	// 	log.Println(string(out))
	// 	return "", err
	// }
	// // 读取结果，得到链接 需要protobuf

	return
}
