package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}
	switch {
	case n%2 == 0 && n%3 == 0:
		return "rock and roll\n"
	case n%2 == 0:
		return "rock\n"
	case n%3 == 0:
		return "roll\n"
	default:
		return "error: non divisible\n"
	}
}
