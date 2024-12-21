package models

import (
	"image/color"
)

const (
	light = 0
	dark  = 1
)

type GobanTheme struct {
	id uint8

	blackStoneFill   color.Color
	blackStoneStroke color.Color
	whiteStoneFill   color.Color
	whiteStoneStroke color.Color
}

func (t *GobanTheme) GetFilePathName() string {
	switch t.id {
	case light:
		return "light"
	case dark:
		return "dark"
	default:
		return ""
	}
}

func NewLightGobanTheme() *GobanTheme {
	return &GobanTheme{
		id:               light,
		blackStoneFill:   color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
		blackStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
		whiteStoneFill:   color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		whiteStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
	}
}

func NewDarkGobanTheme() *GobanTheme {
	return &GobanTheme{
		id:               dark,
		blackStoneFill:   color.RGBA{R: 0x20, G: 0x20, B: 0x24, A: 0xFF},
		blackStoneStroke: color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		whiteStoneFill:   color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},
		whiteStoneStroke: color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},
	}
}
