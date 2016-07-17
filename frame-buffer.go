package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

type FrameBuffer struct {
	Texture *Texture

	FBO gl.Uint
}

func NewFrameBuffer(width, height int) *FrameBuffer {
	f := &FrameBuffer{}

	f.Texture = NewTexture()
	f.Texture.Allocate(width,height)

	gl.GenFramebuffers(1, &f.FBO)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FBO)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, f.Texture.TextureId, 0)

	stdGlSetup()

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return f
}

var a gl.Enum = gl.COLOR_ATTACHMENT0

func (f *FrameBuffer) Begin() {
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, f.FBO);
	gl.DrawBuffers(1, &a)
}

func (f *FrameBuffer) End() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (f *FrameBuffer) Free() {
	f.Texture.Free()

	gl.DeleteFramebuffers(1, &f.FBO);
}
