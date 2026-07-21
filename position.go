package lipgloss

type Position float64

func (p Position) value() float64 { _ = "STUB: not implemented"; return 0 }

const (
	Top    Position = 0.0
	Bottom Position = 1.0
	Center Position = 0.5
	Left   Position = 0.0
	Right  Position = 1.0
)

func Place(width, height int, hPos, vPos Position, str string, opts ...WhitespaceOption) string {
	_ = "STUB: not implemented"
	return ""
}

func PlaceHorizontal(width int, pos Position, str string, opts ...WhitespaceOption) string {
	_ = "STUB: not implemented"
	return ""
}

func PlaceVertical(height int, pos Position, str string, opts ...WhitespaceOption) string {
	_ = "STUB: not implemented"
	return ""
}
