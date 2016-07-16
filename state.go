package g5

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/amortaza/go-adt"
)

var g_projection mgl32.Mat4

var g_colorRect *ColorRect
var g_textureRect *TextureRect
var g_stringRect *TextureRect
var g_canvasRect *TextureRect

var g_viewportWidthStack  adt.Stack
var g_viewportHeightStack adt.Stack
var g_orthoStack adt.Stack


