package main

import (
	"testing"
)

func TestRed(t *testing.T) {
	red("123 %s","123")
}

func TestGreen(t *testing.T) {
	green("123%s","123")
}

func TestBlue(t *testing.T) {
	blue("123%s","123")
}

func TestYellow(t *testing.T) {
	yellow("123%s","123")
}

func TestMagenta(t *testing.T) {
	magenta("123%s111","123")
}