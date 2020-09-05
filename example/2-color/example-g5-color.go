package main

import (
	"github.com/amortaza/go-g5"
	"github.com/amortaza/go-xel"
)

func afterGL() {

	g5.Init()
}

func onBeforeWindowDelete() {

	g5.Uninit()
}

func onLoop() {

	g5.PushView(xel.WinWidth, xel.WinHeight)

	g5.Clear(0.93, 0.93, 0.32)

	g5.DrawColorRect4f(10, 10, 100, 100, 1, 0, 0, 1)
	g5.DrawColorRect4f(20, 120, 100, 100, 1, 0, 0, 0.5)

	leftTop := []float32{1.0, 1.0, 1.0, 1.0}
	rightTop := []float32{1.0, 1.0, 1.0, 0.0}

	g5.DrawColorRect4v(30, 230, 100, 100, leftTop, rightTop, rightTop, leftTop, )
	g5.DrawColorRect4v(40, 340, 100, 100, leftTop, leftTop, rightTop, rightTop, )

	g5.PopView()
}

func main() {

	// I have an ultra-wide screen :)
	xel.Init(1200, 100, 800, 600)

	xel.SetCallbacks(afterGL, onLoop, onBeforeWindowDelete, nil, nil, nil, nil )

	xel.Loop("Bellina")
}

