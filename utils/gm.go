package utils

import (
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"strings"
)

func GenSM3Pwd(pwd string) string {
	return byteToString(sm3.Sm3Sum([]byte(pwd)))
}

func CheckSM3Pwd(pwd string, hash string) bool {
	digest := sm3.Sm3Sum([]byte(pwd))
	return strings.EqualFold(byteToString(digest), hash)
}

func byteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	return ret
}
