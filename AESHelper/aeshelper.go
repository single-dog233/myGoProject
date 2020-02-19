package AESHelper

import (
	"encoding/base64"
	"fmt"
	"github.com/wumansgy/goEncrypt"
)

func AESDecrypt(crypt string) string {
	decodeBytes, _ := base64.StdEncoding.DecodeString(crypt)
	decodePwd, err := goEncrypt.AesCbcDecrypt(decodeBytes, []byte("123456"))
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(decodePwd)
}
