package piscine

func JumpOver(str string) string {
	res := ""
	if str == "" {
		return "\n"
	}
	if len(str) > 3 {
		for i := 2; i < len(str); i += 3 {
			res += string(str[i])
		}
	}
	return res + "\n"
}
