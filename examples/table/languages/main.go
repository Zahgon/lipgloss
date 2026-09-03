package main

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
)

const (
	purple    = "99"
	gray      = "245"
	lightGray = "241"
)

func main() {
	var (
		HeaderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(purple)).Bold(true).Align(lipgloss.Center)

		CellStyle = lipgloss.NewStyle().Padding(0, 1).Width(14)

		OddRowStyle = CellStyle.Foreground(lipgloss.Color(gray))

		EvenRowStyle = CellStyle.Foreground(lipgloss.Color(lightGray))

		BorderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(purple))
	)

	rows := [][]string{
		{"Chinese", "您好", "你好"},
		{"Japanese", "こんにちは", "やあ"},
		{"Arabic", "أهلين", "أهلا"},
		{"Russian", "Здравствуйте", "Привет"},
		{"Spanish", "Hola", "¿Qué tal?"},
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == table.HeaderRow:
				return HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
			}

			if col == 1 {
				style = style.Width(22)
			}

			if row < len(rows) && rows[row][0] == "Arabic" && col != 0 {
				style = style.Align(lipgloss.Right)
			}

			return style
		}).
		Headers("LANGUAGE", "FORMAL", "INFORMAL").
		Rows(rows...)

	t.Row("English", "You look absolutely fabulous.", "How's it going?")

	lipgloss.Println(t)
}
