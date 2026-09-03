package main

import (
	"fmt"
	"image/color"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var gradients = []gradientData{
	{
		name: "Sunset",
		stops: []color.Color{
			lipgloss.Color("#FF6B6B"),
			lipgloss.Color("#FFB74D"),
			lipgloss.Color("#FFDFBA"),
		},
	},
	{
		name: "Ocean",
		stops: []color.Color{
			lipgloss.Color("#0077B6"),
			lipgloss.Color("#48CAE4"),
			lipgloss.Color("#ADE8F4"),
		},
	},
	{
		name: "Forest",
		stops: []color.Color{
			lipgloss.Color("#228B22"),
			lipgloss.Color("#90EE90"),
			lipgloss.Color("#FFFFE0"),
		},
	},
	{
		name: "Purple Dream",
		stops: []color.Color{
			lipgloss.Color("#9370DB"),
			lipgloss.Color("#DDA0DD"),
			lipgloss.Color("#FFB6C1"),
		},
	},
	{
		name: "Fire",
		stops: []color.Color{
			lipgloss.Color("#FF0000"),
			lipgloss.Color("#FFA500"),
			lipgloss.Color("#FFFF00"),
		},
	},
}

type gradientData struct {
	name  string
	stops []color.Color
}

type styles struct {
	title        lipgloss.Style
	gradientName lipgloss.Style
	info         lipgloss.Style
}

func newStyles(dark bool) (s *styles) { _ = "STUB: not implemented"; return nil }

type model struct {
	width  int
	height int
	styles *styles
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	_, err := tea.NewProgram(model{styles: newStyles(true)}).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Uh oh: %v", err)
		os.Exit(1)
	}
}
