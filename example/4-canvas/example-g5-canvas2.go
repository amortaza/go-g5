package main

import (
	"github.com/amortaza/go-g5"
	"github.com/amortaza/go-xel"
)

var canvas *g5.Canvas

func afterGL() {

	g5.Init()

	canvas = g5.NewCanvas(400,400)
}

func onBeforeWindowDelete() {

	canvas.Free()

	g5.Uninit()
}

var f = 0

func onLoop() {

	f++

	g5.PushView(xel.WinWidth, xel.WinHeight)

	g5.Clear(0.3, 0.3, 0.32)

	if f == 100 {
		canvas.Free()

		canvas = g5.NewCanvas(900,800)
	}

	canvas.Begin()
	{
		//g5.Clear(1, 0, 0)
		canvas.Clear(1,0,0)

		//g5.DrawColorRect4f(100,200, 100,100,0,1,0, 1)
	}
	canvas.End()

	canvas.Paint(false, 100, 100, g5.Const_4Ones)

	g5.PopView()
}

func main() {

	xel.Init(1000, 1000)

	xel.SetCallbacks(afterGL, onLoop, onBeforeWindowDelete, nil, nil, nil, nil )

	xel.Loop("G5 - Canvas")
}

