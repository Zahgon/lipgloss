package main

import (
	"image/color"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

func newField(rows, cols int, color color.Color) string { _ = "STUB: not implemented"; return "" }

func newCard(darkMode bool, text string) string { _ = "STUB: not implemented"; return "" }

func main() {
	darkMode := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)
	lightDark := lipgloss.LightDark(darkMode)

	lighterField := newField(17, 43, lightDark(charmtone.Smoke, charmtone.Pepper))
	darkerField := newField(17, 43, lightDark(charmtone.Squid, charmtone.Charcoal))

	pickles := lipgloss.NewLayer(newCard(darkMode, "Pickles"))
	melon := lipgloss.NewLayer(newCard(darkMode, "Bitter Melon"))
	sriracha := lipgloss.NewLayer(newCard(darkMode, "Sriracha"))

	layers := []*lipgloss.Layer{

		lipgloss.NewLayer(lighterField).X(5).Y(2),

		lipgloss.NewLayer(darkerField).AddLayers(
			pickles.X(4).Y(2).Z(1),
			melon.X(22).Y(1),
			sriracha.X(11).Y(7),
		),
	}

	comp := lipgloss.NewCompositor(layers...)

	lipgloss.Println(comp.Render())
}
