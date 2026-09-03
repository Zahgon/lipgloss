package lipgloss

import (
	"image"

	uv "github.com/charmbracelet/ultraviolet"
)

type Layer struct {
	id            string
	content       string
	width, height int
	x, y, z       int
	layers        []*Layer
}

func NewLayer(content string, layers ...*Layer) *Layer { _ = "STUB: not implemented"; return nil }

func (l *Layer) GetContent() string { _ = "STUB: not implemented"; return "" }

func (l *Layer) Width() int { _ = "STUB: not implemented"; return 0 }

func (l *Layer) Height() int { _ = "STUB: not implemented"; return 0 }

func (l *Layer) GetID() string { _ = "STUB: not implemented"; return "" }

func (l *Layer) ID(id string) *Layer { _ = "STUB: not implemented"; return nil }

func (l *Layer) X(x int) *Layer { _ = "STUB: not implemented"; return nil }

func (l *Layer) Y(y int) *Layer { _ = "STUB: not implemented"; return nil }

func (l *Layer) Z(z int) *Layer { _ = "STUB: not implemented"; return nil }

func (l *Layer) GetX() int { _ = "STUB: not implemented"; return 0 }

func (l *Layer) GetY() int { _ = "STUB: not implemented"; return 0 }

func (l *Layer) GetZ() int { _ = "STUB: not implemented"; return 0 }

func (l *Layer) AddLayers(layers ...*Layer) *Layer { _ = "STUB: not implemented"; return nil }

func (l *Layer) GetLayer(id string) *Layer { _ = "STUB: not implemented"; return nil }

func (l *Layer) MaxZ() int { _ = "STUB: not implemented"; return 0 }

func (l *Layer) boundsWithOffset(parentX, parentY int) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

var _ uv.Drawable = (*Layer)(nil)

func (l *Layer) Draw(scr uv.Screen, area uv.Rectangle) { _ = "STUB: not implemented"; return }

type LayerHit struct {
	id     string
	layer  *Layer
	bounds image.Rectangle
}

func (lh LayerHit) Empty() bool { _ = "STUB: not implemented"; return false }

func (lh LayerHit) ID() string { _ = "STUB: not implemented"; return "" }

func (lh LayerHit) Layer() *Layer { _ = "STUB: not implemented"; return nil }

func (lh LayerHit) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

type Compositor struct {
	root   *Layer
	layers []compositeLayer
	index  map[string]*Layer
	bounds image.Rectangle
}

type compositeLayer struct {
	layer  *Layer
	absX   int
	absY   int
	bounds image.Rectangle
}

func NewCompositor(layers ...*Layer) *Compositor { _ = "STUB: not implemented"; return nil }

func (c *Compositor) AddLayers(layers ...*Layer) *Compositor { _ = "STUB: not implemented"; return nil }

func (c *Compositor) flatten() { _ = "STUB: not implemented"; return }

func (c *Compositor) flattenRecursive(layer *Layer, parentX, parentY int) {
	_ = "STUB: not implemented"
	return
}

func (c *Compositor) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (c *Compositor) Draw(scr uv.Screen, area image.Rectangle) { _ = "STUB: not implemented"; return }

func (c *Compositor) Hit(x, y int) LayerHit { _ = "STUB: not implemented"; return *new(LayerHit) }

func (c *Compositor) GetLayer(id string) *Layer { _ = "STUB: not implemented"; return nil }

func (c *Compositor) Refresh() { _ = "STUB: not implemented"; return }

func (c *Compositor) Render() string { _ = "STUB: not implemented"; return "" }
