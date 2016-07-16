package g5

import (
	"golang.org/x/image/font"
	"image"
	"golang.org/x/image/math/fixed"
)

type StringTexture struct {
	String string
	Texture *Texture
}

func NewStringTexture(str string, aceFont *G5font) *StringTexture {
	width := aceFont.Width(str)
	height := aceFont.Height
	descent := aceFont.Descent

	var rgba *image.RGBA

	rgba = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	d := &font.Drawer{	Dst: rgba,
				Src: image.Black,
				Face: *aceFont.Face }

	d.Dot = fixed.P(0, int(height - descent))

	d.DrawString(str)

	ft := &StringTexture{}

	ft.String = str

	texture := NewTexture()

	texture.LoadBytes_RGBA(width, height, rgba.Pix)

	ft.Texture = texture

	return ft
}

func (s *StringTexture) Free() {
	s.Texture.Free()
}