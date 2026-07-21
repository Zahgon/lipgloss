package main

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
)

func addBranches(root *tree.Tree, path string) error { _ = "STUB: not implemented"; return nil }

func main() {
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingRight(1)
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true).PaddingRight(1)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	t := tree.Root(pwd).
		IndenterStyle(enumeratorStyle).
		EnumeratorStyle(enumeratorStyle).
		RootStyle(itemStyle).
		ItemStyle(itemStyle)

	if err := addBranches(t, "."); err != nil {
		fmt.Fprintf(os.Stderr, "Error building tree: %v\n", err)
		os.Exit(1)
	}

	lipgloss.Println(t)
}
