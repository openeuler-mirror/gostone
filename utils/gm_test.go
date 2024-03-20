package utils

import (
	"fmt"
	"testing"
)

func TestGenSM3Pwd(t *testing.T) {
	fmt.Println(GenSM3Pwd("123456"))
}

func TestCheckSM3Pwd(t *testing.T) {
	hash := GenSM3Pwd("123")
	result := CheckSM3Pwd("123", hash)
	fmt.Println(result)
	if !result {
		t.Error("")
		return
	}
}
