package utils

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestCheckPwd(t *testing.T) {
	flag := CheckPw("D60SbpzY0817", "$2b$12$Pew1nj86Jw/pywqi.lTU5eUgD9C6DJXIyVqa2vGP3lP5ORAVsPhiW")
	assert.Equal(t, true, flag)
}

func TestCheckSM3Pwd2(t *testing.T) {
	flag := CheckSM3Pwd("D60SbpzY0817", "a964a104a4282aaade1eea6bb455820922de2f1e5f7ebc524550c0f70ac6ec5d")
	assert.Equal(t, true, flag)

}

func TestGenPwd(t *testing.T) {
	pwd := HashPw("123456", GenSalt("$2b", 12))
	fmt.Println(pwd)
	pwd2 := HashPw("654321", GenSalt("$2b", 12))
	fmt.Println(pwd2)
	assert.Equal(t, true, CheckPw("123456", pwd))
	assert.Equal(t, true, CheckPw("654321", pwd2))
}
