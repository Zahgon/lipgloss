package lipgloss

import (
	"io"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
)

func Wrap(s string, width int, breakpoints string) string { _ = "STUB: not implemented"; return "" }

//nolint:errcheck

type WrapWriter struct {
	w     io.Writer
	p     *ansi.Parser
	style uv.Style
	link  uv.Link
}

func NewWrapWriter(w io.Writer) *WrapWriter { _ = "STUB: not implemented"; return nil }

func (w *WrapWriter) Style() uv.Style { _ = "STUB: not implemented"; return *new(uv.Style) }

func (w *WrapWriter) Link() uv.Link { _ = "STUB: not implemented"; return *new(uv.Link) }

func (w *WrapWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *WrapWriter) Close() error { _ = "STUB: not implemented"; return nil }
