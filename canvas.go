package g5

import "github.com/amortaza/go-adt"

var g_frameBufferMultiSampledStack adt.Stack

type Canvas struct {
	Width, Height            int

	framebufferSingleSampled *_FrameBufferSingleSampled
	framebufferMultiSampled  *_FrameBufferMultiSampled
}

func NewCanvas(width, height int) *Canvas {

	canvas := &Canvas{}

	canvas.framebufferSingleSampled = newFrameBufferSingleSampled(width, height)
	canvas.framebufferMultiSampled = newFrameBufferMultiSampled(width, height)

	canvas.Width, canvas.Height = width, height

	return canvas
}

func (c *Canvas) GetWidth() int {

	return c.Width
}

func (c *Canvas) GetHeight() int {

	return c.Height
}

func (c *Canvas) Begin() {

	c.framebufferMultiSampled.Begin()

	g_frameBufferMultiSampledStack.Push(c.framebufferMultiSampled)

	texture := c.framebufferSingleSampled.Texture

	PushView(texture.Width, texture.Height)
}

func (c *Canvas) Clear(red, green, blue float32) {

	ClearRect(c.Width, c.Height, red, green, blue)
}

func (c *Canvas) Paint(seeThru bool, left, top int, alphas []float32) {

	var flippedAlphas []float32

	if alphas == nil {
		flippedAlphas = Const_4Ones

	} else {

		// need to flip alphas
		flippedAlphas = []float32{alphas[3], alphas[2], alphas[1], alphas[0]}
	}

	if seeThru {
		DrawTextureRectUpsideDown(c.framebufferSingleSampled.Texture, left, top, c.Width, c.Height, flippedAlphas)

	} else {
		DrawCanvasRect(c, left, top, c.Width, c.Height, flippedAlphas)
	}
}

func (c *Canvas) End() {

	c.framebufferMultiSampled.Transfer(c.framebufferSingleSampled)

	PopView()

	c.framebufferMultiSampled.End()

	g_frameBufferMultiSampledStack.Pop()

	if g_frameBufferMultiSampledStack.Size > 0 {

		frameBufferMultiSampled := g_frameBufferMultiSampledStack.Top().(*_FrameBufferMultiSampled)

		frameBufferMultiSampled.Begin()
	}
}

func (c *Canvas) Free() {

	c.framebufferSingleSampled.Free()
	c.framebufferMultiSampled.Free()
}


