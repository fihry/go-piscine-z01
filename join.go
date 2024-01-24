package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i, v := range strs {
		if i == 0 {
			res += v
		} else {
			res += sep + v
		}
	}
	return res
}
