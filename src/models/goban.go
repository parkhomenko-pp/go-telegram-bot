package models

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"
)

type Goban struct {
	size  uint8
	dots  [][]uint8
	theme GobanTheme
}

const (
	empty = 0
	black = 1
	white = 2

	startSizePx     = 60
	rectangleSizePx = 151
	stoneRadPx      = 55
)

func newGoban(size uint8) *Goban {
	dots := make([][]uint8, size)
	for i := range dots {
		dots[i] = make([]uint8, size)
	}
	return &Goban{size: size, dots: dots, theme: *NewLightGobanTheme()}
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

func (g *Goban) ChangeTheme(theme *GobanTheme) {
	g.theme = *theme
}

func (g *Goban) Print() {
	println("Size:", g.size)
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

func (g *Goban) PlaceBlack(s, i int) error {
	if s < 0 || s >= len(g.dots) || i < 0 || i >= len(g.dots) {
		return errors.New("out of range")
	}
	if g.dots[i][s] != empty {
		return errors.New("already placed")
	}
	g.dots[i][s] = black
	return nil
}

func (g *Goban) PlaceWhite(s, i int) error {
	if s < 0 || s >= len(g.dots) || i < 0 || i >= len(g.dots) {
		return errors.New("out of range")
	}
	if g.dots[i][s] != empty {
		return errors.New("already placed")
	}

	g.dots[i][s] = white
	return nil
}

func (g *Goban) String() string {
	var result string
	for _, row := range g.dots {
		for _, dot := range row {
			switch dot {
			case empty:
				result += "·"
			case black:
				result += "⚫"
			case white:
				result += "⚪️"
			}
		}
		result += "\n"
	}
	return result
}

func DrawCircle(img draw.Image, cx, cy, r int, col color.Color) {
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if x*x+y*y <= r*r {
				img.Set(cx+x, cy+y, col)
			}
		}
	}
}

func (g *Goban) loadBackground() (image.Image, error) {
	sourceImageFile, err := os.Open(
		"media/gobans/" + strconv.Itoa(int(g.size)) + "-" + g.theme.GetFilePathName() + ".png",
	)
	if err != nil {
		println(err.Error())
	}

	defer sourceImageFile.Close()

	return png.Decode(sourceImageFile)
}

func (g *Goban) GetImage() **image.RGBA {
	gobanImage, _ := g.loadBackground()

	drawableImage := image.NewRGBA(gobanImage.Bounds())
	draw.Draw(drawableImage, gobanImage.Bounds(), gobanImage, image.Point{}, draw.Src)

	for i, row := range g.dots {
		for j, dot := range row {
			if dot == empty {
				continue
			}

			jPosition := startSizePx + (j-1)*rectangleSizePx
			iPosition := startSizePx + (i-1)*rectangleSizePx

			if dot == black {
				DrawCircle(
					drawableImage,
					jPosition, iPosition,
					stoneRadPx,
					g.theme.blackStoneStroke,
				)
				DrawCircle(
					drawableImage,
					jPosition, iPosition,
					stoneRadPx-2,
					g.theme.blackStoneFill,
				)
				continue
			}

			if dot == white {
				DrawCircle(
					drawableImage,
					jPosition, iPosition,
					stoneRadPx,
					g.theme.whiteStoneStroke,
				)
				DrawCircle(
					drawableImage,
					jPosition, iPosition,
					stoneRadPx-2,
					g.theme.whiteStoneFill,
				)
				continue
			}
		}
	}

	return &drawableImage
}
