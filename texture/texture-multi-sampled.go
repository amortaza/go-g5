package g5

import (
	gl "github.com/chsc/gogl/gl33"
	_ "image/png"
	_ "image/jpeg"
)

type TextureMultiSampled struct {
	TextureId     gl.Uint

	Width, Height int

	textureUnit   gl.Enum
}

func NewTextureMultiSampled(width, height int) *TextureMultiSampled {
	t := &TextureMultiSampled{}

	gl.GenTextures(1, &t.TextureId)

	t.Width  = width
	t.Height = height

	t.Activate(gl.TEXTURE0)

	gl.TexImage2DMultisample(gl.TEXTURE_2D_MULTISAMPLE, 4, gl.RGBA8, gl.Sizei(t.Width), gl.Sizei(t.Height), gl.Boolean(1))

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);

	t.Deactivate()

	return t
}

func (t *TextureMultiSampled) Activate(texUnit gl.Enum) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(gl.TEXTURE_2D_MULTISAMPLE, t.TextureId)
	t.textureUnit = texUnit
}

func (t *TextureMultiSampled) Deactivate() {
	t.textureUnit = 0
	gl.BindTexture(gl.TEXTURE_2D_MULTISAMPLE, 0)
}

func (t *TextureMultiSampled) Free() {
	t.Deactivate()
	gl.DeleteTextures(1, &t.TextureId);
	t.TextureId = 0
}

