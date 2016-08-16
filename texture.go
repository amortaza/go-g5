package g5

import (
	gl "github.com/chsc/gogl/gl33"
	_ "image/png"
	_ "image/jpeg"
	"os"
	"image"
	"image/draw"
)

type Texture struct {
	TextureId     gl.Uint

	Width, Height int

	textureUnit   gl.Enum
}

func NewTexture() *Texture {
	t := &Texture{}

	gl.GenTextures(1, &t.TextureId)

	return t
}

func (t *Texture) Allocate(width, height int) {

	t.Width  = width
	t.Height = height

	t.Activate(gl.TEXTURE0)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.Sizei(t.Width), gl.Sizei(t.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(nil))
	//gl.TexImage2DMultisample(gl.TEXTURE_2D_MULTISAMPLE, 4, gl.RGBA8, gl.Sizei(t.Width), gl.Sizei(t.Height), gl.Boolean(0))

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);

	t.Deactivate()
}

func (t *Texture) Activate(texUnit gl.Enum) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(gl.TEXTURE_2D, t.TextureId)
	t.textureUnit = texUnit
}

func (t *Texture) Deactivate() {
	t.textureUnit = 0
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *Texture) Free() {
	t.Deactivate()
	gl.DeleteTextures(1, &t.TextureId);
	t.TextureId = 0
}

func (t *Texture) LoadImage(filename string) {

	rgba, img := loadRGBA(filename)

	draw.Draw(rgba, rgba.Bounds(), *img, image.Pt(0, 0), draw.Src)

	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("wrong stride")
	}

	t.Width  = rgba.Rect.Size().X
	t.Height = rgba.Rect.Size().Y

	dataPtr := gl.Pointer(GLptr(rgba.Pix))

	t.Activate(gl.TEXTURE0)

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, gl.Sizei(t.Width), gl.Sizei(t.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, dataPtr)

	t.Deactivate()
}

func (t *Texture) LoadBytes_RGBA(width, height int, bytes []uint8) {

	t.Width  = width
	t.Height = height

	dataPtr := gl.Pointer(GLptr(bytes))

	t.Activate(gl.TEXTURE0)

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, gl.Sizei(t.Width), gl.Sizei(t.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, dataPtr)

	t.Deactivate()
}

func loadRGBA(filename string) (*image.RGBA, *image.Image) {
	imgFile, err := os.Open(filename)

	if err != nil {
		panic(err.Error())
	}

	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)

	if err != nil {
		panic(err.Error())
	}

	rgba := image.NewRGBA(img.Bounds())

	return rgba, &img
}
