package main

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/list"
)

var purchased = []string{
	"Bananas",
	"Barley",
	"Cashews",
	"Coconut Milk",
	"Dill",
	"Eggs",
	"Fish Cake",
	"Leeks",
	"Papaya",
}

func groceryEnumerator(items list.Items, i int) string { _ = "STUB: not implemented"; return "" }

var dimEnumStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240")).
	MarginRight(1)

var highlightedEnumStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10")).
	MarginRight(1)

func enumStyleFunc(items list.Items, i int) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func itemStyleFunc(items list.Items, i int) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func main() {
	l := list.New(
		"Artichoke",
		"Baking Flour", "Bananas", "Barley", "Bean Sprouts",
		"Cashew Apple", "Cashews", "Coconut Milk", "Curry Paste", "Currywurst",
		"Dill", "Dragonfruit", "Dried Shrimp",
		"Eggs",
		"Fish Cake", "Furikake",
		"Jicama",
		"Kohlrabi",
		"Leeks", "Lentils", "Licorice Root",
	).
		Enumerator(groceryEnumerator).
		EnumeratorStyleFunc(enumStyleFunc).
		ItemStyleFunc(itemStyleFunc)

	lipgloss.Println(l)
}
