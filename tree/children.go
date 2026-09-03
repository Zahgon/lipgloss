package tree

type Children interface {
	At(index int) Node

	Length() int
}

type NodeChildren []Node

func (n NodeChildren) Append(child Node) NodeChildren {
	_ = "STUB: not implemented"
	return *new(NodeChildren)
}

func (n NodeChildren) Remove(index int) NodeChildren {
	_ = "STUB: not implemented"
	return *new(NodeChildren)
}

func (n NodeChildren) Length() int { _ = "STUB: not implemented"; return 0 }

func (n NodeChildren) At(i int) Node { _ = "STUB: not implemented"; return *new(Node) }

func NewStringData(data ...string) Children { _ = "STUB: not implemented"; return *new(Children) }

var _ Children = NewFilter(nil)

type Filter struct {
	data   Children
	filter func(index int) bool
}

func NewFilter(data Children) *Filter { _ = "STUB: not implemented"; return nil }

func (m *Filter) At(index int) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *Filter) Filter(f func(index int) bool) *Filter { _ = "STUB: not implemented"; return nil }

func (m *Filter) Length() int { _ = "STUB: not implemented"; return 0 }
