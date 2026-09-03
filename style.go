package lipgloss

import (
	"image/color"

	"github.com/charmbracelet/x/ansi"
)

const (
	NBSP            = '\u00A0'
	tabWidthDefault = 4
)

type propKey int64

const (
	boldKey propKey = 1 << iota
	italicKey
	strikethroughKey
	reverseKey
	blinkKey
	faintKey
	underlineSpacesKey
	strikethroughSpacesKey
	colorWhitespaceKey

	underlineKey
	foregroundKey
	backgroundKey
	underlineColorKey
	widthKey
	heightKey
	alignHorizontalKey
	alignVerticalKey

	paddingTopKey
	paddingRightKey
	paddingBottomKey
	paddingLeftKey
	paddingCharKey

	marginTopKey
	marginRightKey
	marginBottomKey
	marginLeftKey
	marginBackgroundKey
	marginCharKey

	borderStyleKey

	borderTopKey
	borderRightKey
	borderBottomKey
	borderLeftKey

	borderTopForegroundKey
	borderRightForegroundKey
	borderBottomForegroundKey
	borderLeftForegroundKey
	borderForegroundBlendKey
	borderForegroundBlendOffsetKey

	borderTopBackgroundKey
	borderRightBackgroundKey
	borderBottomBackgroundKey
	borderLeftBackgroundKey

	inlineKey
	maxWidthKey
	maxHeightKey
	tabWidthKey

	transformKey

	linkKey
	linkParamsKey
)

type props int64

func (p props) set(k propKey) props { _ = "STUB: not implemented"; return *new(props) }

func (p props) unset(k propKey) props { _ = "STUB: not implemented"; return *new(props) }

func (p props) has(k propKey) bool { _ = "STUB: not implemented"; return false }

type Underline = ansi.Underline

const (
	UnderlineNone = ansi.UnderlineNone

	UnderlineSingle = ansi.UnderlineSingle

	UnderlineDouble = ansi.UnderlineDouble

	UnderlineCurly = ansi.UnderlineCurly

	UnderlineDotted = ansi.UnderlineDotted

	UnderlineDashed = ansi.UnderlineDashed
)

func NewStyle() Style { _ = "STUB: not implemented"; return *new(Style) }

type Style struct {
	props props
	value string

	link, linkParams string

	attrs int

	fgColor color.Color
	bgColor color.Color
	ulColor color.Color

	ul Underline

	width  int
	height int

	alignHorizontal Position
	alignVertical   Position

	paddingTop    int
	paddingRight  int
	paddingBottom int
	paddingLeft   int
	paddingChar   rune

	marginTop     int
	marginRight   int
	marginBottom  int
	marginLeft    int
	marginBgColor color.Color
	marginChar    rune

	borderStyle                 Border
	borderTopFgColor            color.Color
	borderRightFgColor          color.Color
	borderBottomFgColor         color.Color
	borderLeftFgColor           color.Color
	borderBlendFgColor          []color.Color
	borderForegroundBlendOffset int
	borderTopBgColor            color.Color
	borderRightBgColor          color.Color
	borderBottomBgColor         color.Color
	borderLeftBgColor           color.Color

	maxWidth  int
	maxHeight int
	tabWidth  int

	transform func(string) string
}

func joinString(strs ...string) string { _ = "STUB: not implemented"; return "" }

func (s Style) SetString(strs ...string) Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Value() string { _ = "STUB: not implemented"; return "" }

func (s Style) String() string { _ = "STUB: not implemented"; return "" }

func (s Style) Copy() Style { _ = "STUB: not implemented"; return *new(Style) }

func (s Style) Inherit(i Style) Style { _ = "STUB: not implemented"; return *new(Style) }

//nolint:exhaustive

func (s Style) Render(strs ...string) string { _ = "STUB: not implemented"; return "" }

//nolint:nestif

func (s Style) maybeConvertTabs(str string) string { _ = "STUB: not implemented"; return "" }

func (s Style) applyMargins(str string, inline bool) string { _ = "STUB: not implemented"; return "" }

func padLeft(str string, n int, style *ansi.Style, r rune) string {
	_ = "STUB: not implemented"
	return ""
}

func padRight(str string, n int, style *ansi.Style, r rune) string {
	_ = "STUB: not implemented"
	return ""
}

func pad(str string, n int, style *ansi.Style, r rune) string { _ = "STUB: not implemented"; return "" }

func abs(a int) int { _ = "STUB: not implemented"; return 0 }
