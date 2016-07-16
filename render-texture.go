package g5

import (
	gl "github.com/chsc/gogl/gl33"
	_ "image/png"
	_ "image/jpeg"
)

type RenderTexture struct {
	TextureId     gl.Uint

	Width, Height int

	textureUnit   gl.Enum
}

func NewRenderTexture(width, height int) *RenderTexture {
	t := &RenderTexture{}

	//t.Width  = width
	//t.Height = height
	//
	//gl.ActiveTexture(gl.TEXTURE0)
	//gl.GenTextures(1, &t.TextureId)
	//gl.BindTexture(gl.TEXTURE_2D, t.TextureId)
	//gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.Sizei(t.Width), gl.Sizei(t.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(nil))
	//
	//gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	//gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);
	//
	//t.Deactivate()



	t.Width  = width
	t.Height = height

	gl.GenTextures(1, &t.TextureId)

	t.Activate(gl.TEXTURE0)

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.Sizei(t.Width), gl.Sizei(t.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(nil))

	t.Deactivate()

	return t
}

func (t *RenderTexture) Activate(texUnit gl.Enum) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(gl.TEXTURE_2D, t.TextureId)
	t.textureUnit = texUnit
}

func (t *RenderTexture) Deactivate() {
	t.textureUnit = 0
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *RenderTexture) Free() {
	t.Deactivate()
	gl.DeleteTextures(1, &t.TextureId);
	t.TextureId = 0
}

