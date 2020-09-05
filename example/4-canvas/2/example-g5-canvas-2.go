package main

import (
	"github.com/amortaza/go-g5"
	"github.com/amortaza/go-xel"
)

var canvas *g5.Canvas

var troll *g5.Texture

func afterGL() {

	g5.Init()

	canvas = g5.NewCanvas(320,240)

	troll = g5.NewTexture()
	troll.Allocate(400, 365)
	troll.LoadImage("github.com/amortaza/go-g5/example/troll.png")
}

func onBeforeWindowDelete() {

	troll.Free()

	canvas.Free()

	g5.Uninit()
}

func onLoop() {

	g5.PushView(xel.WinWidth, xel.WinHeight)

	g5.Clear(0.3, 0.3, 0.3)

	g5.DrawTextureRect(troll, 20, 20, 400, 365, g5.Const_4Ones)

	ones := g5.Const_4Ones
	//var sixes = []float32{0.36, 0.6, 0.6, 0.6}

// -------------------------------------------------------------------------------------------------

	canvas.Begin()
	{
		g5.Clear(0.1, 0.1, 0.1)

		g5.DrawColorRect4f(100,100, 250,250,0.6,0,0, 0.9)

		//g5.DrawTextureRect(troll, 30, 30 , 180, 150, sixes)
	}
	canvas.End()

	canvas.Paint(true, 100, 150, ones)

	g5.PopView()
}

func main() {

	// I have an ultra-wide screen :)
	xel.Init(1200, 50, 1300, 1000)

	xel.SetCallbacks(afterGL, onLoop, onBeforeWindowDelete, nil, nil, nil, nil )

	xel.Loop("G5 - Canvas")
}

