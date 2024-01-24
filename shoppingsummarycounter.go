package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	ls := make(map[string]int)
	var item string
	for _, e := range str {
		if e == 32 {
			ls[item] += 1
			item = ""
		} else if e != 32 {
			item += string(byte(e))
		}
	}
	ls[item] += 1

	return ls
}
