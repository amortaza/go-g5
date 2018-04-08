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

var first = 0
func onLoop() {

	if first < 2 {
		first++
	} else {
		return
	}

	//g5.PushView(xel.WinWidth*2, xel.WinHeight/2)

	g5.Clear(0.93, 0.93, 0.32)

	//g5.PopView()
}

func main() {

	xel.Init(800, 600)

	xel.SetCallbacks(afterGL, onLoop, onBeforeWindowDelete, nil, nil, nil, nil )

	xel.Loop("G5 - Color")
}

