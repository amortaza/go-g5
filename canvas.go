package g5

import (
	"github.com/amortaza/go-adt"
 	//gl "github.com/chsc/gogl/gl33"
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
	//gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0);   // Make sure no FBO is set as the draw framebuffer
	//gl.BindFramebuffer(gl.READ_FRAMEBUFFER, c.Framebuffer.FBO); // Make sure your multisampled FBO is the read framebuffer
	//gl.DrawBuffer(gl.FRONT);                       // Set the back buffer as the draw buffer
	//gl.BlitFramebuffer(0, 0, gl.Int(c.Width), gl.Int(c.Height), 0, 0, gl.Int(c.Width), gl.Int(c.Height), gl.COLOR_BUFFER_BIT, gl.NEAREST);
	//
	//
	//if true {
	//	return
	//}

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


