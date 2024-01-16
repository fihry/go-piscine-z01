package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == toFind[0] {
			return i
		}
	}
	return -1
}
