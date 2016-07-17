package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

type FrameBufferMS struct {
	TextureMS *TextureMS

	FBO gl.Uint
	RBO gl.Uint
}

func NewFrameBufferMS(width, height int) *FrameBufferMS {
	f := &FrameBufferMS{}

	f.TextureMS = NewTextureMS(width,height)

	//

	gl.GenFramebuffers(1, &f.FBO)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FBO)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D_MULTISAMPLE, f.TextureMS.TextureId, 0)

	//

	gl.GenRenderbuffers(1, &f.RBO)
	gl.BindRenderbuffer(gl.RENDERBUFFER, f.RBO)

	// multi-sampled
	gl.RenderbufferStorageMultisample(gl.RENDERBUFFER, 4, gl.DEPTH24_STENCIL8, gl.Sizei(width), gl.Sizei(height))

	gl.BindRenderbuffer(gl.RENDERBUFFER, 0) //?
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, f.RBO)

	stdGlSetup()

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return f
}

var ams gl.Enum = gl.COLOR_ATTACHMENT0

func (f *FrameBufferMS) Begin() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FBO);
	gl.DrawBuffers(1, &ams)
	stdGlSetup()
}

func (f *FrameBufferMS) End() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (ms *FrameBufferMS) Transfer(ss *FrameBuffer) {
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, ms.FBO)
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, ss.FBO)
	gl.BlitFramebuffer(0, 0, gl.Int(ms.TextureMS.Width), gl.Int(ms.TextureMS.Height), 0, 0, gl.Int(ms.TextureMS.Width), gl.Int(ms.TextureMS.Height), gl.COLOR_BUFFER_BIT, gl.NEAREST);
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (f *FrameBufferMS) Free() {
	f.TextureMS.Free()

	gl.DeleteRenderbuffers(1, &f.RBO)
	gl.DeleteFramebuffers(1, &f.FBO);
}
