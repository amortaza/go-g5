package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

type FrameBuffer struct {
	Texture *Texture

	FBO gl.Uint
	RBO gl.Uint
}

func NewFrameBuffer(width, height int) *FrameBuffer {
	f := &FrameBuffer{}

	f.Texture = NewTexture()
	f.Texture.Allocate(width,height)

	//

	gl.GenFramebuffers(1, &f.FBO)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FBO)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, f.Texture.TextureId, 0)
	//gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D_MULTISAMPLE, f.Texture.TextureId, 0)

	//

	gl.GenRenderbuffers(1, &f.RBO)
	gl.BindRenderbuffer(gl.RENDERBUFFER, f.RBO)

	// old
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, gl.Sizei(width), gl.Sizei(height))

	// new
	//gl.RenderbufferStorageMultisample(gl.RENDERBUFFER, 4, gl.DEPTH24_STENCIL8, gl.Sizei(width), gl.Sizei(height))

	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, f.RBO)

	stdGlSetup()

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return f
}

var a gl.Enum = gl.COLOR_ATTACHMENT0

func (f *FrameBuffer) Begin() {
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, f.FBO);
	gl.DrawBuffers(1, &a)
	stdGlSetup()
}

func (f *FrameBuffer) End() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (f *FrameBuffer) Free() {
	f.Texture.Free()

	gl.DeleteRenderbuffers(1, &f.RBO)
	gl.DeleteFramebuffers(1, &f.FBO);
}
