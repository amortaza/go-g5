package g5

import (
	"github.com/go-gl/mathgl/mgl32"
	gl "github.com/chsc/gogl/gl33"
	"fmt"
)

var ThreeZeroesFloat32 = []float32{0,0,0}
var ThreeOnesFloat32 = []float32{1,1,1}
var FourOnesFloat32 = []float32{1,1,1,1}

func Init() {
	gl.Init()

	gl.ClearColor(0.1, 0.4, 0.4, 1.0)

	stdGlSetup()

	g_colorRect = NewColorRect()
	g_textureRect = NewTextureRect("github.com/amortaza/go-g5/shader/texture.vertex.txt", "github.com/amortaza/go-g5/shader/texture.fragment.txt")

	fmt.Println("(+) Initialized G5")
}

func Clear(red,green,blue,alpha float32) {
	gl.ClearColor(gl.Float(red),gl.Float(green),gl.Float(blue),gl.Float(alpha))
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	//gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Uninit() {
	g_textureRect.Free()
	g_colorRect.Free()

	fmt.Println("(-) Uninitialized G5")
}

func PushView(width, height int) {
	PushViewport(width, height)
	PushOrtho(width,height)
}

func PopView() {
	PopViewport()
	PopOrtho()
}

func PushViewport(width, height int) {
	g_viewportWidthStack.Push(width)
	g_viewportHeightStack.Push(height)

	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height));
}

func PopViewport() {
	g_viewportWidthStack.Pop()
	g_viewportHeightStack.Pop()

	if g_viewportWidthStack.Size != 0 {
		width, _ := g_viewportWidthStack.Top().(int)
		height, _ := g_viewportHeightStack.Top().(int)

		gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height));
	}
}

func PushOrtho(width, height int) {
	g_projection = mgl32.Ortho2D(0, float32(width), float32(height), 0)
	g_orthoStack.Push(g_projection)
}

func PopOrtho() {
	g_orthoStack.Pop()

	if g_orthoStack.Size != 0 {
		g_projection = g_orthoStack.Top().(mgl32.Mat4)
	}
}
