package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = -42
	ptr.y = 21
}

func ItoaConv(i int) string {
	var a string
	isNegative := false
	if i < 0 {
		isNegative = true
		i = i * -1
	}
	if i == 0 {
		return "0"
	}
	for i != 0 {
		a = string(i%10+'0') + a
		i = i / 10
	}
	if isNegative {
		a = "-" + a
	}
	return a
}

func main() {
	points := &point{
		x: 0,
		y: 0,
	}
	setPoint(points)
	xStr := ItoaConv(points.x)
	yStr := ItoaConv(points.y)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, r := range xStr {
		z01.PrintRune(r)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, r := range yStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
