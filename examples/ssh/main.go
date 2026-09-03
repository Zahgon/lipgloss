package main

import (
	"fmt"
	"log"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/v2"
)

type styles struct {
	bold          lipgloss.Style
	faint         lipgloss.Style
	italic        lipgloss.Style
	underline     lipgloss.Style
	strikethrough lipgloss.Style
	red           lipgloss.Style
	green         lipgloss.Style
	yellow        lipgloss.Style
	blue          lipgloss.Style
	magenta       lipgloss.Style
	cyan          lipgloss.Style
	gray          lipgloss.Style
}

func makeStyles() styles { _ = "STUB: not implemented"; return *new(styles) }

func handler(next ssh.Handler) ssh.Handler { _ = "STUB: not implemented"; return *new(ssh.Handler) }

func main() {
	port := 3456
	s, err := wish.NewServer(
		ssh.AllocatePty(),
		wish.WithAddress(fmt.Sprintf(":%d", port)),
		wish.WithHostKeyPath("ssh_example"),
		wish.WithMiddleware(handler),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SSH server listening on port %d", port)
	log.Printf("To connect from your local machine run: ssh localhost -p %d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
