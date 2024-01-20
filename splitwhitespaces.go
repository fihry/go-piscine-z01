package piscine

func SplitWhiteSpaces(s string) []string {
	sta := []string{}
	st := ""
	for i := 0; i < len(s); i++ {
		if !(s[i] == ' ' || s[i] == '\t' || s[i] == '\n') {
			st += string(s[i])
		} else if st != "" {
			sta = append(sta, st)
			st = ""
		}
	}
	if st != "" {
		sta = append(sta, st)
	}
	return sta
}
