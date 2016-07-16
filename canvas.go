package g5

import (
	"github.com/amortaza/go-adt"
)
 
var g_canvasStack adt.Stack

type Canvas struct {
	Framebuffer *FrameBuffer
	Width, Height int
}

func NewCanvas(width, height int) *Canvas {
	canvas := &Canvas{}

	canvas.Framebuffer = NewFrameBuffer(width, height)

	canvas.Width, canvas.Height = width, height

	return canvas
}

func (c *Canvas) Begin() {
	c.Framebuffer.Begin()

	g_canvasStack.Push(c.Framebuffer)

	texture := c.Framebuffer.Texture

	PushView(texture.Width, texture.Height)
}

func (c *Canvas) Clear(red, green, blue float32) {
	ClearRect(c.Width, c.Height, red, green, blue)
}

var allOnes = []float32{1,1,1,1}

func (c *Canvas) Paint(seeThru bool, left, top int, alphas []float32) {
	if alphas == nil {
		alphas = allOnes
	}

	if seeThru {
		DrawTextureRectUpsideDown(c.Framebuffer.Texture, left, top, c.Width, c.Height, alphas)
	} else {
		DrawCanvasRect(c, left, top, c.Width, c.Height, alphas)
	}
}

func (c *Canvas) End() {
	PopView()

	c.Framebuffer.End()

	g_canvasStack.Pop()

	if g_canvasStack.Size > 0 {

		frameBuffer := g_canvasStack.Top().(*FrameBuffer)

		frameBuffer.Begin()
	}
}

func (c *Canvas) Free() {
	c.Framebuffer.Free()
}


