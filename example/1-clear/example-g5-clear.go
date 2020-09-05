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

	g5.Clear(0.93, 0.93, 0.32)
}

func main() {

	// I have an ultra-wide screen :)
	xel.Init(1200, 100, 800, 600)

	xel.SetCallbacks(afterGL, onLoop, onBeforeWindowDelete, nil, nil, nil, nil )

	xel.Loop("Bellina")
}

