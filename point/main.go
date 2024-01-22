package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func toStr(i int) string {
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
		a = string(rune(i%10+'0')) + a
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
	result := "x = " + toStr(points.x) + ", y = " + toStr(points.y) + "\n"
	for _, r := range result {
		z01.PrintRune(r)
	}
}
