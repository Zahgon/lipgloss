package main

import (
	"cmp"
	"fmt"
	"image/color"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

var gradients = [][]color.Color{
	{
		lipgloss.Color("#FF6B6B"),
		lipgloss.Color("#FFB74D"),
		lipgloss.Color("#FFDFBA"),
	},
	{
		lipgloss.Color("#0077B6"),
		lipgloss.Color("#48CAE4"),
		lipgloss.Color("#ADE8F4"),
	},
	{
		lipgloss.Color("#228B22"),
		lipgloss.Color("#90EE90"),
		lipgloss.Color("#FFFFE0"),
	},
	{
		lipgloss.Color("#9370DB"),
		lipgloss.Color("#DDA0DD"),
		lipgloss.Color("#FFB6C1"),
	},
	{
		lipgloss.Color("#9900FF"),
		lipgloss.Color("#00FA68"),
		lipgloss.Color("#ED5353"),
	},
}

func main() {
	m := model{
		boxWidth:         20,
		boxHeight:        10,
		angle:            45,
		selectedGradient: 0,

		infoStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			MarginTop(1),
		controlsStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			MarginTop(1),
		gradientBoxStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForegroundBlend(
				charmtone.Cherry,
				charmtone.Charple,
				charmtone.Guac,
				charmtone.Charple,
				charmtone.Sriracha,
			),
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	windowWidth      int
	windowHeight     int
	boxWidth         int
	boxHeight        int
	angle            float64
	selectedGradient int

	infoStyle        lipgloss.Style
	controlsStyle    lipgloss.Style
	gradientBoxStyle lipgloss.Style
	gradients        []color.Color
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m *model) updateGradient() { _ = "STUB: not implemented"; return }

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func clamp[T cmp.Ordered](v, low, high T) T { _ = "STUB: not implemented"; return *new(T) }
