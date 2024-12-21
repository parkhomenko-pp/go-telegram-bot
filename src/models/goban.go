package models

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"strconv"
)

type Goban struct {
	size      uint8
	dots      [][]uint8
	lastColor uint8
	lastI     uint8
	lastJ     uint8
	theme     GobanTheme
}

const (
	empty = 0
	black = 1
	white = 2

	startSizePx     = 60
	rectangleSizePx = 151
	stoneRadPx      = 55
	lastStoneRadPx  = 14
)

func newGoban(size uint8) *Goban {
	dots := make([][]uint8, size)
	for i := range dots {
		dots[i] = make([]uint8, size)
	}
	return &Goban{size: size, dots: dots, theme: *NewLightGobanTheme(), lastColor: 0}
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

func (g *Goban) place(s, i int, color uint8) error {
	g.dots[i][s] = color
	g.lastI = uint8(i)
	g.lastJ = uint8(s)
	g.lastColor = color
	return nil
}

func (g *Goban) PlaceBlack(s, i int) error {
	if s < 0 || s >= len(g.dots) || i < 0 || i >= len(g.dots) {
		return errors.New("out of range")
	}
	if g.dots[i][s] != empty {
		return errors.New("already placed")
	}
	if g.lastColor == black {
		return errors.New("cannot place two black")
	}

	return g.place(s, i, black)
}

func (g *Goban) PlaceWhite(s, i int) error {
	if s < 0 || s >= len(g.dots) || i < 0 || i >= len(g.dots) {
		return errors.New("out of range")
	}
	if g.dots[i][s] != empty {
		return errors.New("already placed")
	}
	if g.lastColor == white {
		return errors.New("cannot place two white")
	}

	return g.place(s, i, white)
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
			dist := math.Sqrt(float64(x*x + y*y))
			if dist <= float64(r) {
				alpha := 1.0
				if dist > float64(r)-1 {
					alpha = float64(r) - dist
				}
				originalColor := img.At(cx+x, cy+y)
				r1, g1, b1, a1 := originalColor.RGBA()
				r2, g2, b2, a2 := col.RGBA()
				newR := uint8((float64(r1)*(1-alpha) + float64(r2)*alpha) / 256)
				newG := uint8((float64(g1)*(1-alpha) + float64(g2)*alpha) / 256)
				newB := uint8((float64(b1)*(1-alpha) + float64(b2)*alpha) / 256)
				newA := uint8((float64(a1)*(1-alpha) + float64(a2)*alpha) / 256)
				img.Set(cx+x, cy+y, color.RGBA{newR, newG, newB, newA})
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
				if i == int(g.lastI) && j == int(g.lastJ) {
					DrawCircle(
						drawableImage,
						jPosition, iPosition,
						lastStoneRadPx,
						g.theme.lastBlackStoneFill,
					)

				}
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
				if i == int(g.lastI) && j == int(g.lastJ) {
					DrawCircle(
						drawableImage,
						jPosition, iPosition,
						lastStoneRadPx,
						g.theme.lastWhiteStoneFill,
					)
				}
				continue
			}
		}
	}

	return &drawableImage
}
