package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
)

var (
	frameColor      = compat.AdaptiveColor{Light: lipgloss.Color("#C5ADF9"), Dark: lipgloss.Color("#864EFF")}
	textColor       = compat.AdaptiveColor{Light: lipgloss.Color("#696969"), Dark: lipgloss.Color("#bdbdbd")}
	keywordColor    = compat.AdaptiveColor{Light: lipgloss.Color("#37CD96"), Dark: lipgloss.Color("#22C78A")}
	inactiveBgColor = compat.AdaptiveColor{Light: lipgloss.Color("#988F95"), Dark: lipgloss.Color("#978692")}
	inactiveFgColor = compat.AdaptiveColor{Light: lipgloss.Color("#FDFCE3"), Dark: lipgloss.Color("#FBFAE7")}
)

type styles struct {
	frame,
	paragraph,
	text,
	keyword,
	activeButton,
	inactiveButton lipgloss.Style
}

func newStyles() (s styles) { _ = "STUB: not implemented"; return *new(styles) }

type model struct {
	styles  styles
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
	initialModel := model{
		yes:    true,
		styles: newStyles(),
	}
	m, err := tea.NewProgram(initialModel).Run()
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
