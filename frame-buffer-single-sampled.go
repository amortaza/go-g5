package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

type FrameBufferSingleSampled struct {
	Texture *Texture

	FBO gl.Uint
}

func NewFrameBufferSingleSampled(width, height int) *FrameBufferSingleSampled {
	f := &FrameBufferSingleSampled{}

	f.Texture = NewTexture()
	f.Texture.Allocate(width,height)

	gl.GenFramebuffers(1, &f.FBO)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FBO)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, f.Texture.TextureId, 0)

	stdGlSetup()

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return f
}

func (f *FrameBufferSingleSampled) Begin() {
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, f.FBO);

	var colorAttachment gl.Enum = gl.COLOR_ATTACHMENT0
	gl.DrawBuffers(1, &colorAttachment)
}

func (f *FrameBufferSingleSampled) End() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (f *FrameBufferSingleSampled) Free() {
	f.Texture.Free()

	gl.DeleteFramebuffers(1, &f.FBO);
}



