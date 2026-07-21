package lipgloss

type whitespace struct {
	chars string
	style Style
}

func newWhitespace(opts ...WhitespaceOption) *whitespace { _ = "STUB: not implemented"; return nil }

func (w whitespace) render(width int) string { _ = "STUB: not implemented"; return "" }

type WhitespaceOption func(*whitespace)

func WithWhitespaceStyle(s Style) WhitespaceOption {
	_ = "STUB: not implemented"
	return *new(WhitespaceOption)
}

func WithWhitespaceChars(s string) WhitespaceOption {
	_ = "STUB: not implemented"
	return *new(WhitespaceOption)
}
