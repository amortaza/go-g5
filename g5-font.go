package g5

import (
	"io/ioutil"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"fmt"
)

type G5font struct {
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

func NewG4Font(f *truetype.Font, fontSize int) *G5font {
	g4font := &G5font{}

	face := truetype.NewFace(f, &truetype.Options{
		Size:    float64(fontSize),
		DPI:     196.0,
		Hinting: font.HintingNone })

	metrics := face.Metrics()

	g4font.Height = metrics.Height.Ceil()
	g4font.Descent = metrics.Descent.Ceil()
	g4font.Ascent = metrics.Ascent.Ceil()
	g4font.Face = &face

	return g4font
}

func (f *G5font) Width(str string) int {
	return int(font.MeasureString(*f.Face, str)>>6)
}

func (f *G5font) Free() {
	fmt.Println("Free has not been implemented for font.AceFont")
}