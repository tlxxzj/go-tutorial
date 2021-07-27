package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var p = make([][]uint8, dx)
	for i := range p {
		p[i] = make([]uint8, dy)
		for j := range p[i] {
			p[i][j] = uint8((i + j) / 2)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}

