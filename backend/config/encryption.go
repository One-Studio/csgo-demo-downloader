package config

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	//"errors"

	log "github.com/sirupsen/logrus"
)

//  ----------------------- 3DES-CBC -----------------------

// TDesEncoding 3des加密
func TDesEncoding(srcByte []byte, desKey []byte, iv []byte) ([]byte, error) {
	// todo key 24位数
	block, err := des.NewTripleDESCipher(desKey) // 和des的区别

	if err != nil {
		log.Errorf("TDesEncoding: [%T]%s, %d", err, err)
		return nil, err
	}
	//对明文进行填充
	srcByte = Padding(srcByte, block.BlockSize())
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//对连续的数据块进行加密
	dst := make([]byte, len(srcByte))
	blockMode.CryptBlocks(dst, srcByte)
	return dst, nil
}

// TDesDecoding 3des解密
func TDesDecoding(pwdByte []byte, desKey []byte, iv []byte) ([]byte, error) {
	block, errBlock := des.NewTripleDESCipher(desKey) // 和des的区别
	if errBlock != nil {
		return nil, errBlock
	}
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCDecrypter(block, iv)

	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)

	//解密
	plainText := make([]byte, len(pwdByte))
	blockMode.CryptBlocks(plainText, pwdByte)
	//删除填充
	plainText = UnPadding(plainText)
	//返回明文
	return plainText, nil
}

//对明文进行填充
func Padding(plainText []byte, blockSize int) []byte {
	//计算要填充的长度
	n := blockSize - len(plainText)%blockSize
	//对原来的明文填充n个n
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

//对密文删除填充
func UnPadding(cipherText []byte) []byte {
	//取出密文最后一个字节end
	end := cipherText[len(cipherText)-1]
	//删除填充
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}
