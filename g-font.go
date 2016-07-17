package g5

import (
	"io/ioutil"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"fmt"
)

type Gfont struct {
	Face *font.Face
	Height int
	Ascent int
	Descent int
}

func LoadTrueTypeFromFile(fontFilename string) *truetype.Font {
	fmt.Println("Reading font file ", fontFilename)
	fontBytes, _ := ioutil.ReadFile(fontFilename)

	f, _ := truetype.Parse(fontBytes)

	return f
}

func NewGfont(f *truetype.Font, fontSize int) *Gfont {
	g5font := &Gfont{}

	face := truetype.NewFace(f, &truetype.Options{
		Size:    float64(fontSize),
		DPI:     196.0,
		Hinting: font.HintingNone })

	metrics := face.Metrics()

	g5font.Height = metrics.Height.Ceil()
	g5font.Descent = metrics.Descent.Ceil()
	g5font.Ascent = metrics.Ascent.Ceil()
	g5font.Face = &face

	return g5font
}

func (f *Gfont) Width(str string) int {
	return int(font.MeasureString(*f.Face, str)>>6)
}

func (f *Gfont) Free() {
	fmt.Println("Free has not been implemented for font.AceFont")
}