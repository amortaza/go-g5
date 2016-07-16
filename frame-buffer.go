package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

type FrameBuffer struct {
	Texture *Texture

	FramebufferId gl.Uint
}

func NewFrameBuffer(width, height int) *FrameBuffer {
	f := &FrameBuffer{}

	f.Texture = NewTexture()

	f.Texture.Allocate(width, height)

	gl.GenFramebuffers(1, &f.FramebufferId)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FramebufferId)
	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, f.Texture.TextureId, 0)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return f
}

func (f *FrameBuffer) Begin() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FramebufferId);
}

func (f *FrameBuffer) End() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (f *FrameBuffer) Free() {
	f.Texture.Free()

	gl.DeleteFramebuffers(1, &f.FramebufferId);
}
