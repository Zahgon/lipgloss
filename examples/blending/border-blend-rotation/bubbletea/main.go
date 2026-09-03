package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

const (
	borderRotationFPS   = 15
	borderRotationSteps = 5
)

type borderRotationTickMsg struct {
	Value int
}

func borderRotationTick(current int) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

type model struct {
	borderRotation int
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	_, err := tea.NewProgram(model{}).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Uh oh: %v", err)
		os.Exit(1)
	}
}
