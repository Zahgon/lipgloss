package lipgloss

import (
	"image/color"
	"io"
	"time"

	"github.com/charmbracelet/x/ansi"
)

func queryBackgroundColor(in io.Reader, out io.Writer) (c color.Color, err error) {
	_ = "STUB: not implemented"
	return *new(color.Color), nil
}

const defaultQueryTimeout = time.Second * 2

type queryTerminalFilter func(seq string, pa *ansi.Parser) bool

func queryTerminal(
	in io.Reader,
	out io.Writer,
	timeout time.Duration,
	filter queryTerminalFilter,
	query string,
) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint: errcheck
