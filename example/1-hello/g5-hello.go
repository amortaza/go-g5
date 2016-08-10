package main

import (
	"github.com/amortaza/go-xel2"
	"fmt"
	gl3 "github.com/chsc/gogl/gl33"
	"github.com/amortaza/go-g5"
)

var arial18 *g5.Gfont
var str1 *g5.StringTexture
var str2 *g5.StringTexture
var canvas *g5.Canvas
var canvas2 *g5.Canvas

func afterGL() {

	e := gl3.Init()

	if e != nil {
		panic("ok")
	}

	g5.Init()

	f := g5.LoadTrueTypeFromFile("github.com/amortaza/go-g5/assets/fonts/arial.ttf")
	arial18 = g5.NewGfont(f, 8)
	str1 = g5.NewStringTexture("Welcome to Clown World!", arial18)
	str2 = g5.NewStringTexture("TEST", arial18)
	canvas = g5.NewCanvas(640,480)
	canvas2 = g5.NewCanvas(320,240)
}

func onDelete() {
	arial18.Free()
	str1.Free()

	g5.Uninit()
}

var t float32
var w, h int = 0,0

func onLoop() {
	t+=0.175
	g5.PushView(w,h)

	g5.Clear(0.3, 0.3, 0.32, 1.0)

	canvas2.Begin()
	{
		g5.Clear(0.1, 0.5, 0.1, 1.0)
		g5.DrawStringRect(str2, 10, 10, g5.ThreeOnesFloat32, g5.ThreeZeroesFloat32, 1)
	}
	canvas2.End()

	//canvas2.Paint(false, 100, 200, g5.FourOnesFloat32)

	canvas.Begin()
	{
		g5.Clear(0.51, 0.1, 0.51, 1.0)
		g5.DrawStringRect(str1,10,10, g5.ThreeOnesFloat32, g5.ThreeZeroesFloat32, 1)
		canvas2.Paint(false, 10, 60, g5.FourOnesFloat32)
	}
	canvas.End()

	canvas.Paint(false, 400, 400, g5.FourOnesFloat32)



	g5.PopView()
}

func onResize(a,b int) {
	w,h=a,b
	fmt.Println(w, " ", h)
}

func main() {
	xel.Init("Welcome, home!", 1280, 1024)

	xel.SetCallbacks(afterGL, onLoop, onDelete, onResize, nil, nil, nil )
	xel.Loop()
	xel.Uninit()
}

