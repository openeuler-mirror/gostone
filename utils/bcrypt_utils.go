package utils

func GetPwd(password string) string {
	return HashPw(password, GenSalt("$2b", 12))
}

func CheckPwd(pwd string, hash string) bool {
	return CheckPw(pwd, hash)
}
