package piscine

func LoafOfBread(str string) string {
	res := ""
	crWord := ""
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	for _, c := range str {
		if c != ' ' && len(crWord) != 5 {
			crWord += string(c)
		} else if len(crWord) == 5 {
			res += crWord + " "
			crWord = ""
		}
	}
	if crWord != "" {
		res += crWord
	}
	return res + "\n"
}
