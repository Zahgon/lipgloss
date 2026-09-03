package lipgloss

import (
	"image/color"
)

func (s *Style) set(key propKey, value any) { _ = "STUB: not implemented"; return }

//nolint:nestif

func (s *Style) setFrom(key propKey, i Style) { _ = "STUB: not implemented"; return }

func colorOrNil(c any) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (s Style) Bold(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Italic(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Underline(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) UnderlineStyle(u Underline) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) UnderlineColor(c color.Color) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Strikethrough(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Reverse(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Blink(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Faint(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Foreground(c color.Color) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Background(c color.Color) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Width(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Height(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Align(p ...Position) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) AlignHorizontal(p Position) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) AlignVertical(p Position) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Padding(i ...int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) PaddingLeft(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) PaddingRight(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) PaddingTop(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) PaddingBottom(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) PaddingChar(r rune) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) ColorWhitespace(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Margin(i ...int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) MarginLeft(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) MarginRight(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) MarginTop(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) MarginBottom(i int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) MarginBackground(c color.Color) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) MarginChar(r rune) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Border(b Border, sides ...bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) BorderStyle(b Border) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) BorderTop(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) BorderRight(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) BorderBottom(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) BorderLeft(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) BorderForeground(c ...color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderTopForeground(c color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderRightForeground(c color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderBottomForeground(c color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderLeftForeground(c color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderForegroundBlend(c ...color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderForegroundBlendOffset(v int) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderBackground(c ...color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderTopBackground(c color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderRightBackground(c color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderBottomBackground(c color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) BorderLeftBackground(c color.Color) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) Inline(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) MaxWidth(n int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) MaxHeight(n int) Style { _ = "STUB: not implemented"; return *new(Style) }

const NoTabConversion = -1

func (s Style) TabWidth(n int) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) UnderlineSpaces(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) StrikethroughSpaces(v bool) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Transform(fn func(string) string) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func (s Style) Hyperlink(link string, params ...string) Style {
	_ = "STUB: not implemented"
	return *new(Style)
}

func whichSidesInt(i ...int) (top, right, bottom, left int, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, false
}

//nolint:mnd

//nolint:mnd

//nolint:mnd

func whichSidesBool(i ...bool) (top, right, bottom, left bool, ok bool) {
	_ = "STUB: not implemented"
	return false, false, false, false, false
}

//nolint:mnd

//nolint:mnd

//nolint:mnd

func whichSidesColor(i ...color.Color) (top, right, bottom, left color.Color, ok bool) {
	_ = "STUB: not implemented"
	return *new(color.Color), *new(color.Color), *new(color.Color), *new(color.Color), false
}

//nolint:mnd

//nolint:mnd

//nolint:mnd
