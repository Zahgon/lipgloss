package main

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
)

type styles struct {
	base,
	block,
	pink,
	dir,
	toggle,
	file lipgloss.Style
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

func main() {
	s := defaultStyles()

	t := tree.Root(dir{"~/charm", true, s}).
		Enumerator(tree.RoundedEnumerator).
		IndenterStyle(s.pink).
		EnumeratorStyle(s.pink).
		Child(
			dir{"ayman", false, s},
			tree.Root(dir{"bash", true, s}).
				Child(
					tree.Root(dir{"tools", true, s}).
						Child(
							file{"zsh", s},
							file{"doom-emacs", s},
						),
				),
			tree.Root(dir{"carlos", true, s}).
				Child(
					tree.Root(dir{"emotes", true, s}).
						Child(
							file{"chefkiss.png", s},
							file{"kekw.png", s},
						),
				),
			dir{"maas", false, s},
		)

	lipgloss.Println(s.block.Render(t.String()))
}
