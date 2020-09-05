package main

import (
	"github.com/amortaza/go-g5"
	"github.com/amortaza/go-xel"
)

var canvas *g5.Canvas
var canvas2 *g5.Canvas

var troll *g5.Texture

func afterGL() {

	g5.Init()

	canvas = g5.NewCanvas(320,240)
	canvas2 = g5.NewCanvas(160,120)

	troll = g5.NewTexture()
	troll.Allocate(400, 365)
	troll.LoadImage("github.com/amortaza/go-g5/example/troll.png")
}

func onBeforeWindowDelete() {

	troll.Free()

	canvas.Free()
	canvas2.Free()

	g5.Uninit()
}

func onLoop() {

	g5.PushView(xel.WinWidth, xel.WinHeight)

	g5.Clear(0.3, 0.3, 0.32)

	g5.DrawTextureRect(troll, 20, 20, 400, 365, g5.Const_4Ones)
	g5.DrawTextureRect(troll, 20, 20 + 240+10, 400, 365, g5.Const_4Ones)

	//var alphas = g5.Const_4Ones
	//var alphas = []float32{1.0, 0.75, 0.75, 0.75}
	//var alphas = []float32{0.0, 1.0, 0.0, 0.0}
	//var alphas = []float32{0.0, 0.0, 0.0, 1.0}
	var alphas = []float32{0.5, 0.5, 0.5, 0.5}

	ones := g5.Const_4Ones
	var sixes = []float32{0.36, 0.6, 0.6, 0.6}
	var nines = []float32{0.9, 0.9, 0.9, 0.9}

	var alphas2 = []float32{0.0, 0.0, 0.0, 1.0}

	canvas2.Begin()
	{
		g5.Clear(1, 0.5, 0.1)
		g5.DrawTextureRect(troll, 20, 20 , 120, 100, sixes)
	}
	canvas2.End()

// -------------------------------------------------------------------------------------------------

	canvas.Begin()
	{
		g5.Clear(0.1, 0.1, 0.1)

		g5.DrawColorRect4f(100,100, 250,250,0.6,0,0, 0.5)

		g5.DrawTextureRect(troll, 30, 30 , 180, 150, sixes)

		canvas2.Paint(true, 50, 100, ones)
	}
	canvas.End()

	canvas.Paint(true, 100, 150, ones)
	canvas.Paint(false, 100, 150 + 240 + 10, ones)

	// -------------------------------------------------------------------------------------------------

	if false {
		alphas = alphas
		alphas2 = alphas2
		nines = nines
		ones = ones
		sixes = sixes
	}

	g5.PopView()
}

func main() {

	// I have an ultra-wide screen :)
	xel.Init(1200, 50, 1300, 1000)

	xel.SetCallbacks(afterGL, onLoop, onBeforeWindowDelete, nil, nil, nil, nil )

	xel.Loop("Bellina")
}

