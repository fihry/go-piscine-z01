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

func main() {
	points := &point{
		x: 0,
		y: 0,
	}
	setPoint(points)
	for _, r := range "x = 42, y = 21\n" {
		z01.PrintRune(r)
	}
}
