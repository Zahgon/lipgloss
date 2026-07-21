package lipgloss

func StyleRanges(s string, ranges ...Range) string { _ = "STUB: not implemented"; return "" }

func NewRange(start, end int, style Style) Range { _ = "STUB: not implemented"; return *new(Range) }

type Range struct {
	Start, End int
	Style      Style
}
