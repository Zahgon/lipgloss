package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type styles struct {
	frame,
	paragraph,
	text,
	keyword,
	activeButton,
	inactiveButton lipgloss.Style
}

func newStyles(backgroundIsDark bool) (s *styles) { _ = "STUB: not implemented"; return nil }

type model struct {
	styles  *styles
	yes     bool
	chosen  bool
	aborted bool
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	m, err := tea.NewProgram(model{}).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Uh oh: %v", err)
		os.Exit(1)
	}

	if m := m.(model); m.chosen {
		if m.yes {
			fmt.Println("Are you sure? It's not ripe yet.")
		} else {
			fmt.Println("Well, alright. It was probably good, though.")
		}
	}
}
