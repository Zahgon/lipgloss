package lipgloss

import (
	"cmp"
	"errors"
	"image/color"

	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
)

func clamp[T cmp.Ordered](v, low, high T) T { _ = "STUB: not implemented"; return *new(T) }

const (
	Black ansi.BasicColor = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

var noColor = NoColor{}

type NoColor struct{}

func (n NoColor) RGBA() (r, g, b, a uint32) {
	_ = "STUB: not implemented"
	return 0,
		//nolint:mnd
		0, 0, 0
}

func Color(s string) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

//nolint:gosec

//nolint:gosec

//nolint:gosec

var errInvalidFormat = errors.New("invalid hex format")

func parseHex(s string) (c color.RGBA, err error) {
	_ = "STUB: not implemented"
	return *new(color.RGBA), nil
}

type RGBColor struct {
	R uint8
	G uint8
	B uint8
}

func (c RGBColor) RGBA() (r, g, b, a uint32) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

type ANSIColor = ansi.IndexedColor

type LightDarkFunc func(light, dark color.Color) color.Color

func LightDark(isDark bool) LightDarkFunc { _ = "STUB: not implemented"; return *new(LightDarkFunc) }

func isDarkColor(c color.Color) bool { _ = "STUB: not implemented"; return false }

//nolint:mnd

type CompleteFunc func(ansi, ansi256, truecolor color.Color) color.Color

func Complete(p colorprofile.Profile) CompleteFunc {
	_ = "STUB: not implemented"
	return *new(CompleteFunc)
}

//nolint:exhaustive

func ensureNotTransparent(c color.Color) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func Alpha(c color.Color, alpha float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func Complementary(c color.Color) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func Darken(c color.Color, percent float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func Lighten(c color.Color, percent float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}
