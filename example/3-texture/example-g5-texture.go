package main

import (
	"github.com/amortaza/go-g5"
	"github.com/amortaza/go-xel"
)

var troll *g5.Texture

func afterGL() {

	g5.Init()

	troll = g5.NewTexture()
	troll.Allocate(400, 365)
	troll.LoadImage("github.com/amortaza/go-g5/example/troll.png")
}

func onBeforeWindowDelete() {

	troll.Free()

	g5.Uninit()
}

func onLoop() {

	g5.PushView(xel.WinWidth, xel.WinHeight)

	g5.Clear(0.3, 0.3, 0.32)

	g5.DrawTextureRectUpsideDown(troll, 10, 10, 400/2, 365/2, g5.Const_4Ones)
	g5.DrawTextureRect(troll, 100, 185, 400, 365, g5.Const_4Ones)

	a := []float32{0.0,0.0,1.0,0.0}
	g5.DrawTextureRect(troll, 200, 15, 400, 365, a)

	g5.PopView()
}

func main() {

	// I have an ultra-wide screen :)
	xel.Init(1200, 100, 800, 600)

	xel.SetCallbacks(afterGL, onLoop, onBeforeWindowDelete, nil, nil, nil, nil )

	xel.Loop("Bellina")
}

