package lipgloss

import (
	"image/color"
)

func (s Style) GetBold() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetItalic() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetUnderline() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetUnderlineStyle() Underline { _ = "STUB: not implemented"; return *new(Underline) }

func (s Style) GetUnderlineColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (s Style) GetStrikethrough() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetReverse() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetBlink() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetFaint() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetForeground() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (s Style) GetBackground() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (s Style) GetWidth() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetHeight() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetAlign() Position { _ = "STUB: not implemented"; return *new(Position) }

func (s Style) GetAlignHorizontal() Position { _ = "STUB: not implemented"; return *new(Position) }

func (s Style) GetAlignVertical() Position { _ = "STUB: not implemented"; return *new(Position) }

func (s Style) GetPadding() (top, right, bottom, left int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func (s Style) GetPaddingTop() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetPaddingRight() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetPaddingBottom() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetPaddingLeft() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetPaddingChar() rune { _ = "STUB: not implemented"; return 0 }

func (s Style) GetHorizontalPadding() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetVerticalPadding() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetColorWhitespace() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetMargin() (top, right, bottom, left int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func (s Style) GetMarginTop() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetMarginRight() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetMarginBottom() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetMarginLeft() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetMarginChar() rune { _ = "STUB: not implemented"; return 0 }

func (s Style) GetHorizontalMargins() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetVerticalMargins() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetBorder() (b Border, top, right, bottom, left bool) {
	_ = "STUB: not implemented"
	return *new(Border), false, false, false, false
}

func (s Style) GetBorderStyle() Border { _ = "STUB: not implemented"; return *new(Border) }

func (s Style) GetBorderTop() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetBorderRight() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetBorderBottom() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetBorderLeft() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetBorderTopForeground() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s Style) GetBorderRightForeground() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s Style) GetBorderBottomForeground() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s Style) GetBorderLeftForeground() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s Style) GetBorderForegroundBlend() []color.Color { _ = "STUB: not implemented"; return nil }

func (s Style) GetBorderForegroundBlendOffset() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetBorderTopBackground() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s Style) GetBorderRightBackground() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s Style) GetBorderBottomBackground() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s Style) GetBorderLeftBackground() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s Style) GetBorderTopWidth() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetBorderTopSize() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetBorderLeftSize() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetBorderBottomSize() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetBorderRightSize() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetHorizontalBorderSize() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetVerticalBorderSize() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetInline() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetMaxWidth() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetMaxHeight() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetTabWidth() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetUnderlineSpaces() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetStrikethroughSpaces() bool { _ = "STUB: not implemented"; return false }

func (s Style) GetHorizontalFrameSize() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetVerticalFrameSize() int { _ = "STUB: not implemented"; return 0 }

func (s Style) GetFrameSize() (x, y int) { _ = "STUB: not implemented"; return 0, 0 }

func (s Style) GetTransform() func(string) string { _ = "STUB: not implemented"; return nil }

func (s Style) GetHyperlink() (link, params string) { _ = "STUB: not implemented"; return "", "" }

func (s Style) isSet(k propKey) bool { _ = "STUB: not implemented"; return false }

func (s Style) getAsRune(k propKey) rune { _ = "STUB: not implemented"; return 0 }

//nolint:exhaustive

func (s Style) getAsBool(k propKey, defaultVal bool) bool { _ = "STUB: not implemented"; return false }

func (s Style) getAsColors(k propKey) (colors []color.Color) { _ = "STUB: not implemented"; return nil }

//nolint:exhaustive

func (s Style) getAsColor(k propKey) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

//nolint:exhaustive

func (s Style) getAsInt(k propKey) int { _ = "STUB: not implemented"; return 0 }

//nolint:exhaustive

func (s Style) getAsPosition(k propKey) Position { _ = "STUB: not implemented"; return *new(Position) }

//nolint:exhaustive

func (s Style) getBorderStyle() Border { _ = "STUB: not implemented"; return *new(Border) }

func (s Style) getAsTransform(propKey) func(string) string { _ = "STUB: not implemented"; return nil }

func getLines(s string) (lines []string, widest int) { _ = "STUB: not implemented"; return nil, 0 }

func (s Style) isBorderStyleSetWithoutSides() bool { _ = "STUB: not implemented"; return false }

//nolint:staticcheck
