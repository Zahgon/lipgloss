package lipgloss

import (
	"image/color"
)

type Border struct {
	Top          string
	Bottom       string
	Left         string
	Right        string
	TopLeft      string
	TopRight     string
	BottomLeft   string
	BottomRight  string
	MiddleLeft   string
	MiddleRight  string
	Middle       string
	MiddleTop    string
	MiddleBottom string
}

func (b Border) GetTopSize() int { _ = "STUB: not implemented"; return 0 }

func (b Border) GetRightSize() int { _ = "STUB: not implemented"; return 0 }

func (b Border) GetBottomSize() int { _ = "STUB: not implemented"; return 0 }

func (b Border) GetLeftSize() int { _ = "STUB: not implemented"; return 0 }

func getBorderEdgeWidth(borderParts ...string) (maxWidth int) { _ = "STUB: not implemented"; return 0 }

var (
	noBorder = Border{}

	normalBorder = Border{
		Top:          "─",
		Bottom:       "─",
		Left:         "│",
		Right:        "│",
		TopLeft:      "┌",
		TopRight:     "┐",
		BottomLeft:   "└",
		BottomRight:  "┘",
		MiddleLeft:   "├",
		MiddleRight:  "┤",
		Middle:       "┼",
		MiddleTop:    "┬",
		MiddleBottom: "┴",
	}

	roundedBorder = Border{
		Top:          "─",
		Bottom:       "─",
		Left:         "│",
		Right:        "│",
		TopLeft:      "╭",
		TopRight:     "╮",
		BottomLeft:   "╰",
		BottomRight:  "╯",
		MiddleLeft:   "├",
		MiddleRight:  "┤",
		Middle:       "┼",
		MiddleTop:    "┬",
		MiddleBottom: "┴",
	}

	blockBorder = Border{
		Top:          "█",
		Bottom:       "█",
		Left:         "█",
		Right:        "█",
		TopLeft:      "█",
		TopRight:     "█",
		BottomLeft:   "█",
		BottomRight:  "█",
		MiddleLeft:   "█",
		MiddleRight:  "█",
		Middle:       "█",
		MiddleTop:    "█",
		MiddleBottom: "█",
	}

	outerHalfBlockBorder = Border{
		Top:         "▀",
		Bottom:      "▄",
		Left:        "▌",
		Right:       "▐",
		TopLeft:     "▛",
		TopRight:    "▜",
		BottomLeft:  "▙",
		BottomRight: "▟",
	}

	innerHalfBlockBorder = Border{
		Top:         "▄",
		Bottom:      "▀",
		Left:        "▐",
		Right:       "▌",
		TopLeft:     "▗",
		TopRight:    "▖",
		BottomLeft:  "▝",
		BottomRight: "▘",
	}

	thickBorder = Border{
		Top:          "━",
		Bottom:       "━",
		Left:         "┃",
		Right:        "┃",
		TopLeft:      "┏",
		TopRight:     "┓",
		BottomLeft:   "┗",
		BottomRight:  "┛",
		MiddleLeft:   "┣",
		MiddleRight:  "┫",
		Middle:       "╋",
		MiddleTop:    "┳",
		MiddleBottom: "┻",
	}

	doubleBorder = Border{
		Top:          "═",
		Bottom:       "═",
		Left:         "║",
		Right:        "║",
		TopLeft:      "╔",
		TopRight:     "╗",
		BottomLeft:   "╚",
		BottomRight:  "╝",
		MiddleLeft:   "╠",
		MiddleRight:  "╣",
		Middle:       "╬",
		MiddleTop:    "╦",
		MiddleBottom: "╩",
	}

	hiddenBorder = Border{
		Top:          " ",
		Bottom:       " ",
		Left:         " ",
		Right:        " ",
		TopLeft:      " ",
		TopRight:     " ",
		BottomLeft:   " ",
		BottomRight:  " ",
		MiddleLeft:   " ",
		MiddleRight:  " ",
		Middle:       " ",
		MiddleTop:    " ",
		MiddleBottom: " ",
	}

	markdownBorder = Border{
		Top:          "-",
		Bottom:       "-",
		Left:         "|",
		Right:        "|",
		TopLeft:      "|",
		TopRight:     "|",
		BottomLeft:   "|",
		BottomRight:  "|",
		MiddleLeft:   "|",
		MiddleRight:  "|",
		Middle:       "|",
		MiddleTop:    "|",
		MiddleBottom: "|",
	}

	asciiBorder = Border{
		Top:          "-",
		Bottom:       "-",
		Left:         "|",
		Right:        "|",
		TopLeft:      "+",
		TopRight:     "+",
		BottomLeft:   "+",
		BottomRight:  "+",
		MiddleLeft:   "+",
		MiddleRight:  "+",
		Middle:       "+",
		MiddleTop:    "+",
		MiddleBottom: "+",
	}
)

func NormalBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func RoundedBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func BlockBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func OuterHalfBlockBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func InnerHalfBlockBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func ThickBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func DoubleBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func HiddenBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func MarkdownBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

func ASCIIBorder() Border { _ = "STUB: not implemented"; return *new(Border) }

type borderBlend struct {
	topGradient    []color.Color
	rightGradient  []color.Color
	bottomGradient []color.Color
	leftGradient   []color.Color
}

func (s Style) borderBlend(width, height int, colors ...color.Color) *borderBlend {
	_ = "STUB: not implemented"
	return nil
}

func (s Style) applyBorder(str string) string { _ = "STUB: not implemented"; return "" }

func renderHorizontalEdge(left, middle, right string, width int) string {
	_ = "STUB: not implemented"
	return ""
}

func (s Style) styleBorder(border string, fg, bg color.Color) string {
	_ = "STUB: not implemented"
	return ""
}

func (s Style) styleBorderBlend(border string, fg []color.Color, bg color.Color) string {
	_ = "STUB: not implemented"
	return ""
}

func maxRuneWidth(str string) int { _ = "STUB: not implemented"; return 0 }

func getFirstRuneAsString(str string) string { _ = "STUB: not implemented"; return "" }
