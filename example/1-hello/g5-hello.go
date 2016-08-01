package main

import (
	"github.com/amortaza/go-xel2"
	"fmt"
	gl3 "github.com/chsc/gogl/gl33"
	"github.com/amortaza/go-g5"
)

func afterGL() {

	e := gl3.Init()

	if e != nil {
		panic("ok")
	}

	g5.Init()
}

func onDelete() {

	g5.Uninit()
}

var w, h int = 0,0

func onLoop() {
	g5.PushView(w,h)

	g5.Clear(0.3, 0.3, 0.32, 1.0)

	//g5.DrawColorRect3f(10, 10, 200, 200, .2, .5, 1)

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

