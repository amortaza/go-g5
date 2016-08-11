package g5

import (
	"github.com/amortaza/go-adt"
)
 
var g_frameBufferMSStack adt.Stack

type Canvas struct {
	FramebufferSingleSampled *FrameBufferSingleSampled
	FramebufferMS            *FrameBufferMultiSampled

	Width, Height            int
}

func NewCanvas(width, height int) *Canvas {
	canvas := &Canvas{}

	canvas.FramebufferSingleSampled = NewFrameBufferSingleSampled(width, height)
	canvas.FramebufferMS = NewFrameBufferMultiSampled(width, height)

	canvas.Width, canvas.Height = width, height

	return canvas
}

func (c *Canvas) Begin() {
	c.FramebufferMS.Begin()

	g_frameBufferMSStack.Push(c.FramebufferMS)

	texture := c.FramebufferSingleSampled.Texture

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
		DrawTextureRectUpsideDown(c.FramebufferSingleSampled.Texture, left, top, c.Width, c.Height, alphas)
	} else {
		DrawCanvasRect(c, left, top, c.Width, c.Height, alphas)
	}
}

func (c *Canvas) End() {
	c.FramebufferMS.Transfer(c.FramebufferSingleSampled)

	PopView()

	c.FramebufferMS.End()

	g_frameBufferMSStack.Pop()

	if g_frameBufferMSStack.Size > 0 {

		frameBufferMS := g_frameBufferMSStack.Top().(*FrameBufferMultiSampled)

		frameBufferMS.Begin()
	}
}

func (c *Canvas) Free() {
	c.FramebufferSingleSampled.Free()
	c.FramebufferMS.Free()
}


