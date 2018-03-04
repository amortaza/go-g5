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

	g5.Clear(0.93, 0.93, 0.32, 1.0)

	g5.PopView()
}

func main() {

	xel.Init(800, 600)

	xel.SetCallbacks(afterGL, onLoop, onBeforeWindowDelete, nil, nil, nil, nil )

	xel.Loop("G5 - Color")
}

