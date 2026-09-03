package list

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
)

type List struct{ tree *tree.Tree }

func New(items ...any) *List { _ = "STUB: not implemented"; return nil }

type Items tree.Children

type StyleFunc func(items Items, index int) lipgloss.Style

func (l *List) Hidden() bool { _ = "STUB: not implemented"; return false }

func (l *List) Hide(hide bool) *List { _ = "STUB: not implemented"; return nil }

func (l *List) Offset(start, end int) *List { _ = "STUB: not implemented"; return nil }

func (l *List) Value() string { _ = "STUB: not implemented"; return "" }

func (l *List) String() string { _ = "STUB: not implemented"; return "" }

func (l *List) EnumeratorStyle(style lipgloss.Style) *List { _ = "STUB: not implemented"; return nil }

func (l *List) EnumeratorStyleFunc(f StyleFunc) *List { _ = "STUB: not implemented"; return nil }

func (l *List) IndenterStyle(style lipgloss.Style) *List { _ = "STUB: not implemented"; return nil }

func (l *List) IndenterStyleFunc(f StyleFunc) *List { _ = "STUB: not implemented"; return nil }

func (l *List) Indenter(indenter Indenter) *List { _ = "STUB: not implemented"; return nil }

func (l *List) ItemStyle(style lipgloss.Style) *List { _ = "STUB: not implemented"; return nil }

func (l *List) ItemStyleFunc(f StyleFunc) *List { _ = "STUB: not implemented"; return nil }

func (l *List) Item(item any) *List { _ = "STUB: not implemented"; return nil }

func (l *List) Items(items ...any) *List { _ = "STUB: not implemented"; return nil }

func (l *List) Enumerator(enumerator Enumerator) *List { _ = "STUB: not implemented"; return nil }
