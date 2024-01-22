package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x string
	y string
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := &point{
		x: "0",
		y: "0",
	}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, r := range points.x {
		z01.PrintRune(r)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, r := range points.y {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
