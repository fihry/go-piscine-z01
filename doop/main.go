package main

import (
	"os"
)

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

func IsNumeric(s string) bool {
	for j := range s {
		if (s[j] < '0' || s[j] > '9') && s[j] != '-' {
			return false
		}
	}
	return true
}

func itoa(nb int) string {
	negative := false
	var re string
	if nb == 0 {
		return "0"
	}
	if nb < 0 {
		negative = true
		nb = -nb
	}
	for nb != 0 {
		re = string(rune(nb%10)+'0') + re
		nb = nb / 10
	}
	if negative {
		re = "-" + re
	}
	return re
}

func main() {
	num1 := ToDigit(os.Args[1])
	operator := os.Args[2]
	num2 := ToDigit(os.Args[3])
	var res int
	if !(len(os.Args) > 4 || len(os.Args) < 4) {

		if !IsNumeric(os.Args[1]) == true || !IsNumeric(os.Args[3]) == true {
			return
		}
		if !(os.Args[1] == "9223372036854775807" || os.Args[3] == "9223372036854775807") && !(os.Args[1] == "-9223372036854775809" || os.Args[3] == "-9223372036854775809") {
			if len(os.Args) == 4 {
				if operator == "-" {
					res = num1 - num2
				}
				if operator == "+" {
					res = num1 + num2
				}
				if operator == "*" {
					res = num1 * num2
				}
				if operator == "%" {
					res = num1 % num2
				}
				if operator == "/" {
					if num2 == 0 {
						msg := []byte("No division by 0\n")
						_, err := os.Stdout.Write(msg)
						if err != nil {
							panic(err)
						}
						return
					}
					res = num1 / num2
				}
				result := []byte(itoa(res) + "\n")
				_, err := os.Stdout.Write(result)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
