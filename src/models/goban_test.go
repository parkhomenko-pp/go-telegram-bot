package models

import (
	"testing"
)

func TestNewGoban(t *testing.T) {
	goban := NewGoban7()
	if goban.size != 7 {
		t.Errorf("expected size 7, got %d", goban.size)
	}
	if len(goban.dots) != 7 {
		t.Errorf("expected 7 rows, got %d", len(goban.dots))
	}
	for _, row := range goban.dots {
		if len(row) != 7 {
			t.Errorf("expected 7 columns, got %d", len(row))
		}
	}
}

func TestPlaceBlack(t *testing.T) {
	goban := NewGoban7()

	err := goban.PlaceBlack(3, 3)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	err = goban.PlaceBlack(3, 3)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceBlack(-1, 3)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestPlaceWhite(t *testing.T) {
	goban := NewGoban7()

	err := goban.PlaceWhite(3, 3)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	err = goban.PlaceWhite(3, 3)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceWhite(7, 7)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestPrint(t *testing.T) {
	goban := NewGoban7()
	goban.PlaceBlack(3, 3)
	goban.PlaceWhite(4, 4)

	expectedOutput := "·······\n·······\n·······\n···⚫···\n····⚪️··\n·······\n·······\n"
	if goban.String() != expectedOutput {
		t.Errorf("expected %v, got %v", expectedOutput, goban.String())
	}
}

func TestChangeTheme(t *testing.T) {
	goban := NewGoban7()
	newTheme := NewDarkGobanTheme()
	goban.ChangeTheme(newTheme)

	if goban.theme != *newTheme {
		t.Errorf("expected theme %v, got %v", newTheme, goban.theme)
	}
}
