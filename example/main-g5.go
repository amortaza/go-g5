package main

import (
	"github.com/amortaza/go-xel-glfw-goxjs"
	"github.com/shibukawa/nanovgo"
	"fmt"
	"github.com/shibukawa/nanovgo/sample/demo"
	"log"
	gl3 "github.com/chsc/gogl/gl33"
	g5 "github.com/amortaza/go-g5-chsc-gogl"
)

var ctx *nanovgo.Context
var demoData *demo.DemoData

var arial18 *g5.G5font
var str1 *g5.StringTexture
var canvas *g5.Canvas

func afterGL() {
	var err error

	ctx, err = nanovgo.NewContext(nanovgo.AntiAlias | nanovgo.StencilStrokes /*| nanovgo.Debug*/)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("(+) Created nanovgo context")

	demoData = LoadDemo(ctx)

	e := gl3.Init()

	if e != nil {
		panic("ok")
	}

	g5.Init()

	f := g5.LoadTrueTypeFromFile("github.com/amortaza/go-bellina-examples/assets/fonts/arial.ttf")
	arial18 = g5.NewG4Font(f, 8)
	str1 = g5.NewStringTexture("Welcome to Clown World!", arial18)
	canvas = g5.NewCanvas(640,480)
}

func onDelete() {
	arial18.Free()
	str1.Free()

	g5.Uninit()
	demoData.FreeData(ctx)

	fmt.Println("(-) Deleting nanovg context")
	ctx.Delete()
}

var t float32
var w, h int = 0,0

func onLoop() {
	t+=0.175
	g5.PushView(w,h)

	g5.Clear(0.3, 0.3, 0.32, 1.0)

	//ctx.BeginFrame(w, h, 1)
	//demo.RenderDemo(ctx, float32(1), float32(1), float32(1024), float32(768), t, false, demoData)
	//ctx.EndFrame()

	canvas.Begin()
	g5.Clear(0.51, 0.51, 0.51, 1.0)
	ctx.BeginFrame(640, 480, 1)
	demo.RenderDemo(ctx, float32(1), float32(1), float32(640), float32(480), t, false, demoData)
	ctx.EndFrame()
	canvas.End()
	canvas.Paint(false, 400, 200, []float32{.5,.5,.5,.5})

	//g5.DrawColorRect3f(0,0,200,200,.5,.1,0)
	g5.DrawStringRect(str1,10,10, g5.ThreeOnesFloat32, g5.ThreeZeroesFloat32, 1)

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

func LoadDemo(ctx *nanovgo.Context) *demo.DemoData {
	d := &demo.DemoData{}
	for i := 0; i < 12; i++ {
		path := fmt.Sprintf("github.com/shibukawa/nanovgo/sample/images/image%d.jpg", i+1)
		d.Images = append(d.Images, ctx.CreateImage(path, 0))
		if d.Images[i] == 0 {
			log.Fatalf("Could not load %s", path)
		}
	}

	d.FontIcons = ctx.CreateFont("icons", "github.com/shibukawa/nanovgo/sample/entypo.ttf")
	if d.FontIcons == -1 {
		log.Fatalln("Could not add font icons.")
	}
	d.FontNormal = ctx.CreateFont("sans", "github.com/shibukawa/nanovgo/sample/Roboto-Regular.ttf")
	if d.FontNormal == -1 {
		log.Fatalln("Could not add font italic.")
	}
	d.FontBold = ctx.CreateFont("sans-bold", "github.com/shibukawa/nanovgo/sample/Roboto-Bold.ttf")
	if d.FontBold == -1 {
		log.Fatalln("Could not add font bold.")
	}
	return d
}
