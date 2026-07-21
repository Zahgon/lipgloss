package tree

import (
	"fmt"
	"sync"

	"charm.land/lipgloss/v2"
)

type Node interface {
	fmt.Stringer
	Value() string
	Children() Children
	Hidden() bool
	SetHidden(bool)
	SetValue(any)
}

type Leaf struct {
	value  string
	hidden bool
}

func NewLeaf(value any, hidden bool) *Leaf { _ = "STUB: not implemented"; return nil }

func (Leaf) Children() Children { _ = "STUB: not implemented"; return *new(Children) }

func (s Leaf) Value() string { _ = "STUB: not implemented"; return "" }

func (s *Leaf) SetValue(value any) { _ = "STUB: not implemented"; return }

func (s Leaf) Hidden() bool { _ = "STUB: not implemented"; return false }

func (s *Leaf) SetHidden(hidden bool) { _ = "STUB: not implemented"; return }

func (s Leaf) String() string { _ = "STUB: not implemented"; return "" }

type Tree struct {
	value    string
	hidden   bool
	offset   [2]int
	children Children

	r     *renderer
	ronce sync.Once
}

func (t *Tree) Hidden() bool { _ = "STUB: not implemented"; return false }

func (t *Tree) Hide(hide bool) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) SetHidden(hidden bool) { _ = "STUB: not implemented"; return }

func (t *Tree) Offset(start, end int) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) Value() string { _ = "STUB: not implemented"; return "" }

func (t *Tree) SetValue(value any) { _ = "STUB: not implemented"; return }

func (t *Tree) String() string { _ = "STUB: not implemented"; return "" }

func (t *Tree) Child(children ...any) *Tree { _ = "STUB: not implemented"; return nil }

func ensureParent(nodes Children, item *Tree) (*Tree, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (t *Tree) ensureRenderer() *renderer { _ = "STUB: not implemented"; return nil }

func (t *Tree) EnumeratorStyle(style lipgloss.Style) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) EnumeratorStyleFunc(fn StyleFunc) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) IndenterStyle(style lipgloss.Style) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) IndenterStyleFunc(fn StyleFunc) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) RootStyle(style lipgloss.Style) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) ItemStyle(style lipgloss.Style) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) ItemStyleFunc(fn StyleFunc) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) Enumerator(enum Enumerator) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) Indenter(indenter Indenter) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) Width(width int) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) Children() Children { _ = "STUB: not implemented"; return *new(Children) }

func Root(root any) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) Root(root any) *Tree { _ = "STUB: not implemented"; return nil }

func New() *Tree { _ = "STUB: not implemented"; return nil }
