package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
)

const selected = "/Users/bash/.config/doom-emacs"

type styles struct {
	base,
	container,
	dir,
	selected,
	dimmed,
	toggle lipgloss.Style
}

func defaultStyles() styles { _ = "STUB: not implemented"; return *new(styles) }

type dir struct {
	name   string
	open   bool
	styles styles
}

func (d dir) String() string { _ = "STUB: not implemented"; return "" }

type file struct {
	name   string
	styles styles
}

func (s file) String() string { _ = "STUB: not implemented"; return "" }

func (s file) Hidden() bool { _ = "STUB: not implemented"; return false }

func (s file) Children() tree.Children { _ = "STUB: not implemented"; return *new(tree.Children) }

func (s file) Value() string { _ = "STUB: not implemented"; return "" }

func (s file) SetValue(val any) { _ = "STUB: not implemented"; return }

func (s file) SetHidden(val bool) { _ = "STUB: not implemented"; return }

func isItemSelected(children tree.Children, index int) bool {
	_ = "STUB: not implemented"
	return false
}

func itemStyle(children tree.Children, index int) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func indenterStyle(children tree.Children, index int) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func main() {
	s := defaultStyles()

	t := tree.Root(dir{"~/charm", true, s}).
		Child(
			dir{"ayman", false, s},
			tree.Root(dir{"bash", true, s}).
				Child(
					file{"/Users/bash/.config/doom-emacs", s},
				),
			tree.Root(dir{"carlos", true, s}).
				Child(
					tree.Root(dir{"emotes", true, s}).
						Child(
							file{"/home/caarlos0/Pictures/chefkiss.png", s},
							file{"/home/caarlos0/Pictures/kekw.png", s},
						),
				),
			dir{"maas", false, s},
		).
		Width(30).
		Indenter(Indenter).
		Enumerator(Enumerator).
		EnumeratorStyleFunc(indenterStyle).
		IndenterStyleFunc(indenterStyle).
		ItemStyleFunc(itemStyle)

	fmt.Println(s.container.Render(t.String()))
}

func Enumerator(children tree.Children, index int) string { _ = "STUB: not implemented"; return "" }

func Indenter(children tree.Children, index int) string { _ = "STUB: not implemented"; return "" }
