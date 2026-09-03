package table

type Data interface {
	At(row, cell int) string

	Rows() int

	Columns() int
}

type StringData struct {
	rows    [][]string
	columns int
}

func NewStringData(rows ...[]string) *StringData { _ = "STUB: not implemented"; return nil }

func (m *StringData) Append(row []string) { _ = "STUB: not implemented"; return }

func (m *StringData) At(row, cell int) string { _ = "STUB: not implemented"; return "" }

func (m *StringData) Columns() int { _ = "STUB: not implemented"; return 0 }

func (m *StringData) Item(rows ...string) *StringData { _ = "STUB: not implemented"; return nil }

func (m *StringData) Rows() int { _ = "STUB: not implemented"; return 0 }

type Filter struct {
	data   Data
	filter func(row int) bool
}

func NewFilter(data Data) *Filter { _ = "STUB: not implemented"; return nil }

func (m *Filter) Filter(f func(row int) bool) *Filter { _ = "STUB: not implemented"; return nil }

func (m *Filter) At(row, cell int) string { _ = "STUB: not implemented"; return "" }

func (m *Filter) Columns() int { _ = "STUB: not implemented"; return 0 }

func (m *Filter) Rows() int { _ = "STUB: not implemented"; return 0 }

func DataToMatrix(data Data) (rows [][]string) { _ = "STUB: not implemented"; return nil }
