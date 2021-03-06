package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

type _FrameBufferMultiSampled struct {
	TextureMS *TextureMultiSampled

	FBO gl.Uint
	RBO gl.Uint
}

func newFrameBufferMultiSampled(width, height int) *_FrameBufferMultiSampled {

	f := &_FrameBufferMultiSampled{}

	f.TextureMS = NewTextureMultiSampled(width,height)

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

func (f *_FrameBufferMultiSampled) Begin() {

	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FBO);

	var colorAttachment gl.Enum = gl.COLOR_ATTACHMENT0
	gl.DrawBuffers(1, &colorAttachment)
}

func (f *_FrameBufferMultiSampled) End() {

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (ms *_FrameBufferMultiSampled) Transfer(ss *_FrameBufferSingleSampled) {

	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, ms.FBO)
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, ss.FBO)
	gl.BlitFramebuffer(0, 0, gl.Int(ms.TextureMS.Width), gl.Int(ms.TextureMS.Height), 0, 0, gl.Int(ms.TextureMS.Width), gl.Int(ms.TextureMS.Height), gl.COLOR_BUFFER_BIT, gl.NEAREST);
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (f *_FrameBufferMultiSampled) Free() {

	f.TextureMS.Free()

	gl.DeleteRenderbuffers(1, &f.RBO)
	gl.DeleteFramebuffers(1, &f.FBO);
}
