package lipgloss

import (
	uv "github.com/charmbracelet/ultraviolet"
)

type Canvas struct {
	scr uv.ScreenBuffer
}

var _ uv.Screen = (*Canvas)(nil)

func NewCanvas(width, height int) *Canvas { _ = "STUB: not implemented"; return nil }

func (c *Canvas) Resize(width, height int) { _ = "STUB: not implemented"; return }

func (c *Canvas) Clear() { _ = "STUB: not implemented"; return }

func (c *Canvas) Bounds() uv.Rectangle { _ = "STUB: not implemented"; return *new(uv.Rectangle) }

func (c *Canvas) Width() int { _ = "STUB: not implemented"; return 0 }

func (c *Canvas) Height() int { _ = "STUB: not implemented"; return 0 }

func (c *Canvas) CellAt(x int, y int) *uv.Cell { _ = "STUB: not implemented"; return nil }

func (c *Canvas) SetCell(x int, y int, cell *uv.Cell) { _ = "STUB: not implemented"; return }

func (c *Canvas) WidthMethod() uv.WidthMethod {
	_ = "STUB: not implemented"
	return *new(uv.WidthMethod)
}

func (c *Canvas) Compose(drawer uv.Drawable) *Canvas { _ = "STUB: not implemented"; return nil }

func (c *Canvas) Draw(scr uv.Screen, area uv.Rectangle) { _ = "STUB: not implemented"; return }

func (c *Canvas) Render() string { _ = "STUB: not implemented"; return "" }
