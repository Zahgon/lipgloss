package compat

import (
	"image/color"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
)

var (
	HasDarkBackground = lipgloss.HasDarkBackground(os.Stdin, os.Stdout)

	Profile = colorprofile.Detect(os.Stdout, os.Environ())
)

type AdaptiveColor struct {
	Light color.Color
	Dark  color.Color
}

func (c AdaptiveColor) RGBA() (uint32, uint32, uint32, uint32) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

type CompleteColor struct {
	TrueColor color.Color
	ANSI256   color.Color
	ANSI      color.Color
}

func (c CompleteColor) RGBA() (uint32, uint32, uint32, uint32) {
	_ = "STUB: not implemented"
	//nolint:exhaustive
	return 0, 0, 0, 0
}

type CompleteAdaptiveColor struct {
	Light CompleteColor
	Dark  CompleteColor
}

func (c CompleteAdaptiveColor) RGBA() (uint32, uint32, uint32, uint32) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}
