package piscine

import "github.com/01-edu/z01"

func Isvalid(bs string) bool {
	if len(bs) < 2 {
		return false
	}
	for i := 0; i < len(bs); i++ {
		for j := i + 1; j < len(bs); j++ {
			if bs[i] == bs[j] || bs[i] == '-' || bs[i] == '+' {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	res := ""
	isNeg := false
	if !Isvalid(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbrBase(9, base)
		PrintNbrBase(223372036854775808, base)
		return
	}

	if nbr < 0 {
		nbr = -nbr
		isNeg = true

	}

	for nbr != 0 {
		res = string(base[nbr%len(base)]) + res
		nbr /= len(base)
	}

	if isNeg {
		res = "-" + res
	}

	for i := range res {
		z01.PrintRune(rune(res[i]))
	}
}
