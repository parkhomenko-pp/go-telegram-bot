package models

import (
	"image/color"
	"testing"
)

func TestNewLightGobanTheme(t *testing.T) {
	theme := NewLightGobanTheme()

	if theme.id != light {
		t.Errorf("expected id %v, got %v", light, theme.id)
	}

	expectedBlackFill := color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF}
	if theme.blackStoneFill != expectedBlackFill {
		t.Errorf("expected blackStoneFill %v, got %v", expectedBlackFill, theme.blackStoneFill)
	}

	expectedWhiteFill := color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF}
	if theme.whiteStoneFill != expectedWhiteFill {
		t.Errorf("expected whiteStoneFill %v, got %v", expectedWhiteFill, theme.whiteStoneFill)
	}
}

func TestNewDarkGobanTheme(t *testing.T) {
	theme := NewDarkGobanTheme()

	if theme.id != dark {
		t.Errorf("expected id %v, got %v", dark, theme.id)
	}

	expectedBlackFill := color.RGBA{R: 0x20, G: 0x20, B: 0x24, A: 0xFF}
	if theme.blackStoneFill != expectedBlackFill {
		t.Errorf("expected blackStoneFill %v, got %v", expectedBlackFill, theme.blackStoneFill)
	}

	expectedWhiteFill := color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF}
	if theme.whiteStoneFill != expectedWhiteFill {
		t.Errorf("expected whiteStoneFill %v, got %v", expectedWhiteFill, theme.whiteStoneFill)
	}
}

func TestGetFilePathName(t *testing.T) {
	lightTheme := NewLightGobanTheme()
	if lightTheme.GetFilePathName() != "light" {
		t.Errorf("expected 'light', got %v", lightTheme.GetFilePathName())
	}

	darkTheme := NewDarkGobanTheme()
	if darkTheme.GetFilePathName() != "dark" {
		t.Errorf("expected 'dark', got %v", darkTheme.GetFilePathName())
	}
}
