package models

import "errors"

type Goban struct {
	dots [][]uint8
}

const (
	empty = 0
	black = 1
	white = 2
)

func newGoban(size int) *Goban {
	dots := make([][]uint8, size)

	for i := range dots {
		dots[i] = make([]uint8, size)
	}

	return &Goban{dots: dots}
}

func NewGoban7() *Goban {
	return newGoban(7)
}

func NewGoban9() *Goban {
	return newGoban(9)
}

func NewGoban13() *Goban {
	return newGoban(13)
}

func NewGoban19() *Goban {
	return newGoban(19)
}

func (g *Goban) Print() {
	for _, row := range g.dots {
		for _, dot := range row {
			switch dot {
			case empty:
				print("·")
			case black:
				print("⚫")
			case white:
				print("⚪️")
			}
		}
		println()
	}
}

func (g *Goban) PlaceBlack(x, y int) error {
	if x < 0 || x >= len(g.dots) || y < 0 || y >= len(g.dots) {
		return errors.New("out of range")
	}

	if g.dots[y][x] != empty {
		return errors.New("already placed")
	}

	g.dots[y][x] = black
	return nil
}

func (g *Goban) PlaceWhite(x, y int) error {
	if x < 0 || x >= len(g.dots) || y < 0 || y >= len(g.dots) {
		return errors.New("out of range")
	}

	if g.dots[y][x] != empty {
		return errors.New("already placed")
	}

	g.dots[y][x] = white
	return nil
}
