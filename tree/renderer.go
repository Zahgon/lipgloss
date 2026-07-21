package tree

import (
	"charm.land/lipgloss/v2"
)

type StyleFunc func(children Children, i int) lipgloss.Style

type Style struct {
	enumeratorFunc StyleFunc
	indenterFunc   StyleFunc
	itemFunc       StyleFunc
	root           lipgloss.Style
}

func newRenderer() *renderer { _ = "STUB: not implemented"; return nil }

type renderer struct {
	style      Style
	enumerator Enumerator
	indenter   Indenter
	width      int
}

func (r *renderer) render(node Node, root bool, prefix string) string {
	_ = "STUB: not implemented"
	return ""
}
