package table

import (
	"charm.land/lipgloss/v2"
)

const HeaderRow int = -1

type StyleFunc func(row, col int) lipgloss.Style

func DefaultStyles(_, _ int) lipgloss.Style { _ = "STUB: not implemented"; return *new(lipgloss.Style) }

type Table struct {
	baseStyle lipgloss.Style
	styleFunc StyleFunc
	border    lipgloss.Border

	borderTop    bool
	borderBottom bool
	borderLeft   bool
	borderRight  bool
	borderHeader bool
	borderColumn bool
	borderRow    bool

	borderStyle lipgloss.Style
	headers     []string
	data        Data

	width           int
	height          int
	useManualHeight bool
	yOffset         int
	wrap            bool

	widths  []int
	heights []int

	firstVisibleRowIndex int
	lastVisibleRowIndex  int
	overflowHeight       int
}

func New() *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) ClearRows() *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BaseStyle(baseStyle lipgloss.Style) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) StyleFunc(style StyleFunc) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) style(row, col int) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (t *Table) Data(data Data) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) GetData() Data { _ = "STUB: not implemented"; return *new(Data) }

func (t *Table) Rows(rows ...[]string) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) Row(row ...string) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) Headers(headers ...string) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) GetHeaders() []string { _ = "STUB: not implemented"; return nil }

func (t *Table) Border(border lipgloss.Border) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BorderTop(v bool) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BorderBottom(v bool) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BorderLeft(v bool) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BorderRight(v bool) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BorderHeader(v bool) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BorderColumn(v bool) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BorderRow(v bool) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) BorderStyle(style lipgloss.Style) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) GetBorderTop() bool { _ = "STUB: not implemented"; return false }

func (t *Table) GetBorderBottom() bool { _ = "STUB: not implemented"; return false }

func (t *Table) GetBorderLeft() bool { _ = "STUB: not implemented"; return false }

func (t *Table) GetBorderRight() bool { _ = "STUB: not implemented"; return false }

func (t *Table) GetBorderHeader() bool { _ = "STUB: not implemented"; return false }

func (t *Table) GetBorderColumn() bool { _ = "STUB: not implemented"; return false }

func (t *Table) GetBorderRow() bool { _ = "STUB: not implemented"; return false }

func (t *Table) Width(w int) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) Height(h int) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) GetHeight() int { _ = "STUB: not implemented"; return 0 }

func (t *Table) YOffset(o int) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) GetYOffset() int { _ = "STUB: not implemented"; return 0 }

func (t *Table) FirstVisibleRowIndex() int { _ = "STUB: not implemented"; return 0 }

func (t *Table) LastVisibleRowIndex() int { _ = "STUB: not implemented"; return 0 }

func (t *Table) VisibleRows() int { _ = "STUB: not implemented"; return 0 }

func (t *Table) Wrap(w bool) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) String() string { _ = "STUB: not implemented"; return "" }

func (t *Table) computeHeight() int { _ = "STUB: not implemented"; return 0 }

func (t *Table) Render() string { _ = "STUB: not implemented"; return "" }

func (t *Table) constructTopBorder() string { _ = "STUB: not implemented"; return "" }

func (t *Table) constructBottomBorder() string { _ = "STUB: not implemented"; return "" }

func (t *Table) constructHeaders() string { _ = "STUB: not implemented"; return "" }

func (t *Table) constructRow(index int, isOverflow bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *Table) truncateCell(cell string, rowIndex, colIndex int) string {
	_ = "STUB: not implemented"
	return ""
}
