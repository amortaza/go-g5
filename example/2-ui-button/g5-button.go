package main

import (
	"github.com/shibukawa/nanovgo"
	"fmt"
	"github.com/amortaza/go-xel2"
	"github.com/amortaza/go-g5"
	gl3 "github.com/chsc/gogl/gl33"
	"github.com/amortaza/go-g5/example"
)

var ctx *nanovgo.Context

func afterGL() {
	var err error

	ctx, err = nanovgo.NewContext(nanovgo.AntiAlias | nanovgo.StencilStrokes /*| nanovgo.Debug*/)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("(+) Created nanovgo context")

	e := gl3.Init()

	if e != nil {
		panic("ok")
	}

	g5.Init()

	//canvas = g5.NewCanvas(640,480)

	ctx.CreateFont("icons", "github.com/shibukawa/nanovgo/sample/entypo.ttf")
	ctx.CreateFont("sans", "github.com/shibukawa/nanovgo/sample/Roboto-Regular.ttf")
	ctx.CreateFont("sans-bold", "github.com/shibukawa/nanovgo/sample/Roboto-Bold.ttf")

	gui.Init(ctx)

	b.Color = nanovgo.RGBA(128,0,0,255)
}

func onDelete() {
	g5.Uninit()
	ctx.Delete()
}

var w, h int = 0,0
var frame, flip int = 0, -1
var b = gui.NewButton()

func onLoop() {
	g5.PushView(w,h)

	g5.Clear(0.3, 0.3, 0.32, 1.0)

	g5.Clear(0.51, 0.51, 0.51, 1.0)
	ctx.BeginFrame(640, 480, 4)

	if frame % 30 == 0 {
		flip++
	}

	if flip % 3 == 0 {
		b.State = gui.Default
	} else if flip % 3 == 1 {
		b.State = gui.Hover
	} else {
		b.State = gui.Pressed
	}

	gui.DrawButton(b, 10,10)

	ctx.EndFrame()

	g5.PopView()

	frame++
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
