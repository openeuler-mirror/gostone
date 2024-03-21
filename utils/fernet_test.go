package utils

import (
	"fmt"
	"testing"
)

/*
func TestFernetToken_Validate(t *testing.T) {
	token := "gAAAAABl0rl9ZUCdvankstlu87ffJtQpuYLCycb7zTwNP7g5o2axK6AQLlHgGKt9wbU9GMgswzDrehWzGoTmT9eSE7gfjLKOWUMUBFBBfHRsDZK1-0TaXa6-EWdKnahbkd3y7S1-5MXWmb0p8KSccNuB8Nsjg39aPoYZyoW5VVe_a2-qiYmEGGU"
	fernet := NewFernetToken()
	fernetPath = "/home/jenkins/agent/workspace/gostone-单测/etc/fernet-keys"
	loadKey()
	fmt.Println(fernet.Validate(token))
}
*/

func TestBuildAuditInfo(t *testing.T) {
	fmt.Println(buildAuditInfo())
}
