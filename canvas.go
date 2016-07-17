package g5

import (
	"github.com/amortaza/go-adt"
 	//gl "github.com/chsc/gogl/gl33"
)
 
var g_frameBufferMSStack adt.Stack

type Canvas struct {
	Framebuffer *FrameBuffer
	FramebufferMS *FrameBufferMS

	Width, Height int
}

func NewCanvas(width, height int) *Canvas {
	canvas := &Canvas{}

	canvas.Framebuffer = NewFrameBuffer(width, height)
	canvas.FramebufferMS = NewFrameBufferMS(width, height)

	canvas.Width, canvas.Height = width, height

	return canvas
}

func (c *Canvas) Begin() {
	c.FramebufferMS.Begin()

	g_frameBufferMSStack.Push(c.FramebufferMS)

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

	c.FramebufferMS.Transfer(c.Framebuffer)

	if seeThru {
		DrawTextureRectUpsideDown(c.Framebuffer.Texture, left, top, c.Width, c.Height, alphas)
	} else {
		DrawCanvasRect(c, left, top, c.Width, c.Height, alphas)
	}
}

func (c *Canvas) End() {
	PopView()

	c.FramebufferMS.End()

	g_frameBufferMSStack.Pop()

	if g_frameBufferMSStack.Size > 0 {

		frameBufferMS := g_frameBufferMSStack.Top().(*FrameBufferMS)

		frameBufferMS.Begin()
	}
}

func (c *Canvas) Free() {
	c.Framebuffer.Free()
}


