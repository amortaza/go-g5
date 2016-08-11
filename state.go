package g5

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/amortaza/go-adt"
	gl "github.com/chsc/gogl/gl33"
)

var g_projection mgl32.Mat4

var g_colorRect *ColorRect
var g_textureRect *TextureRect
var g_stringRect *TextureRect
var g_canvasRect *TextureRect

var g_viewportWidthStack  adt.Stack
var g_viewportHeightStack adt.Stack
var g_orthoStack adt.Stack

//var g_devicePixelRatio int

func stdGlSetup() {
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.STENCIL_TEST)

	// blending is required to be able to render text
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Enable( gl.MULTISAMPLE )
	gl.Enable( gl.SCISSOR_TEST)
}


