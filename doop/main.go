package main

import (
	"os"
)

func main() {
	if len(os.Args) == 4 {
		v1 := os.Args[1]
		op := os.Args[2]
		v2 := os.Args[3]
		nbr1 := ToDigit(v1)
		nbr2 := ToDigit(v2)
		if checker(v1, v2, op) && nbr1 < 9223372036854775807 && nbr1 > -9223372036854775807 && nbr2 < 9223372036854775807 && nbr2 > -9223372036854775807 {
			if op == "/" {
				if nbr2 == 0 {
					message := "No division by 0\n"
					os.Stdout.WriteString(message)
					return
				}
			}
			if op == "%" {
				if nbr2 == 0 {
					message := "No modulo by 0\n"
					os.Stdout.WriteString(message)
					return
				}
			}

			result := calculation(nbr1, op, nbr2)
			res := itoa(result) + "\n"
			message := res
			os.Stdout.WriteString(message)
		}
	}
}

func ToDigit(s string) int {
	sign := 1
	var result int
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
		for ; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				digit := int(s[i] - '0')
				result = result*10 + digit
			} else {
				return 0
			}
		}
	}
	return result * sign
}

func checker(num1, num2, op string) bool {
	res := true
	for i := 0; i < len(num1); i++ {
		if (num1[i] < 48 || num1[i] > 57) && num1[i] != '-' {
			res = false
		}
	}
	for i := 0; i < len(num2); i++ {
		if (num2[i] < 48 || num2[i] > 57) && num2[i] != '-' {
			res = false
		}
	}
	if op != "*" && op != "+" && op != "-" && op != "/" && op != "%" {
		res = false
	}
	return res
}

func calculation(a int, b string, c int) int {
	var res int
	if b == "*" {
		res = a * c
	}
	if b == "-" {
		res = a - c
	}
	if b == "+" {
		res = a + c
	}
	if b == "/" {
		res = a / c
	}
	if b == "%" {
		res = a % c
	}
	return res
}

func itoa(n int) string {
	var res string
	negative := false
	if n == 0 {
		return "0"
	}

	if n < 0 {
		negative = true
		n = n * -1
	}

	for n != 0 {
		digit := n % 10
		res = string(rune(digit+'0')) + res
		n /= 10
	}

	if negative {
		res = "-" + res
	}
	return res
}
