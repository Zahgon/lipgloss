package table

func (t *Table) resize() { _ = "STUB: not implemented"; return }

type resizerColumn struct {
	index      int
	min        int
	max        int
	median     int
	rows       [][]string
	xPadding   int
	fixedWidth int
}

type resizer struct {
	tableWidth  int
	tableHeight int
	headers     []string
	allRows     [][]string
	rowHeights  []int
	columns     []resizerColumn

	wrap         bool
	borderColumn bool
	yPaddings    [][]int

	yOffset         int
	useManualHeight bool
	borderTop       bool
	borderBottom    bool
	borderLeft      bool
	borderRight     bool
	borderHeader    bool
	borderRow       bool
}

func newResizer(tableWidth, tableHeight int, headers []string, rows [][]string) *resizer {
	_ = "STUB: not implemented"
	return nil
}

func (r *resizer) optimizedWidths() (colWidths, rowHeights []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *resizer) detectTableWidth() int { _ = "STUB: not implemented"; return 0 }

func (r *resizer) expandTableWidth() (colWidths []int) { _ = "STUB: not implemented"; return nil }

func (r *resizer) minWidth(j int) int { _ = "STUB: not implemented"; return 0 }

func (r *resizer) shrinkTableWidth() (colWidths []int) { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func (r *resizer) expandRowHeights(colWidths []int) { _ = "STUB: not implemented"; return }

func (r *resizer) defaultRowHeights() (rowHeights []int) { _ = "STUB: not implemented"; return nil }

func (r *resizer) maxColumnWidths() []int { _ = "STUB: not implemented"; return nil }

func (r *resizer) columnCount() int { _ = "STUB: not implemented"; return 0 }

func (r *resizer) maxCharCount() int { _ = "STUB: not implemented"; return 0 }

func (r *resizer) maxTotal() (maxTotal int) { _ = "STUB: not implemented"; return 0 }

func (r *resizer) totalHorizontalPadding() (totalHorizontalPadding int) {
	_ = "STUB: not implemented"
	return 0
}

func (r *resizer) xPaddingForCol(j int) int { _ = "STUB: not implemented"; return 0 }

func (r *resizer) yPaddingForCell(i, j int) int { _ = "STUB: not implemented"; return 0 }

func (r *resizer) totalHorizontalBorder() int { _ = "STUB: not implemented"; return 0 }

func (r *resizer) detectContentHeight(content string, width int) (height int) {
	_ = "STUB: not implemented"
	return 0
}

func (r *resizer) visibleRowIndexes() (firstVisibleRowIndex, lastVisibleRowIndex, overflowHeight int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}
