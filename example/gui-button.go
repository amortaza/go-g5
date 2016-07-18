package gui

import vgo "github.com/shibukawa/nanovgo"

type IconTable struct {
	Trash string
}

type Style struct {
	HoverColor vgo.Color
}

type ButtonState int

const (
	Default ButtonState = iota
	Pressed
	Hover
)

type Button struct {
	Label string
	Icon  int
	W, H  float32
	Color vgo.Color
	State ButtonState
}

func NewButton() *Button {
	return &Button{State : Default, Color : vgo.RGBA(255,0,0,255), Label : "Ok", W : 100, H : 30}
}

func DrawButton(button *Button, x, y float32) {
	var cornerRadius float32 = 6.0
	var iw float32

	var alpha uint8 = 32

	var bg vgo.Paint

	if button.State == Default {
		bg = vgo.LinearGradient(x, y, x, y + button.H, vgo.RGBA(255, 255, 255, alpha), vgo.RGBA(0, 0, 0, alpha))
	} else {
		bg = vgo.LinearGradient(x, y + button.H, x, y, vgo.RGBA(255, 255, 255, alpha), vgo.RGBA(0, 0, 0, alpha))
	}

	ctx.BeginPath()
	ctx.RoundedRect(x, y, button.W, button.H, cornerRadius)
	ctx.SetFillColor(button.Color)
	ctx.Fill()
	ctx.SetFillPaint(bg)
	ctx.Fill()

	// shadow outline
	ctx.BeginPath()
	ctx.RoundedRect(x, y, button.W, button.H, cornerRadius)

	if button.State == Default || button.State == Pressed  {
		ctx.SetStrokeColor(vgo.RGBA(0, 0, 0, 48))
	} else {
		ctx.SetStrokeColor(g_style.HoverColor)
	}

	ctx.SetStrokeWidth(3)
	ctx.Stroke()

	ctx.SetFontSize(20.0)
	ctx.SetFontFace("sans-bold")
	tw, _ := ctx.TextBounds(0, 0, button.Label)
	if button.Icon != 0 {
		ctx.SetFontSize(button.H * 1.3)
		ctx.SetFontFace("icons")
		iw, _ = ctx.TextBounds(0, 0, IconToStr(button.Icon))
		iw += button.H * 0.15

		ctx.SetFillColor(vgo.RGBA(255, 255, 255, 96))
		ctx.SetTextAlign(vgo.AlignLeft | vgo.AlignMiddle)
		ctx.Text(x + button.W * 0.5 - tw * 0.5 - iw * 0.8, y + button.H * 0.5, IconToStr(button.Icon))
	}

	btnx := x + button.W * 0.5 - tw * 0.5 + iw * 0.25
	btny := y + button.H * 0.5

	ctx.SetFontSize(20.0)
	ctx.SetFontFace("sans-bold")
	ctx.SetTextAlign(vgo.AlignLeft | vgo.AlignMiddle)
	ctx.SetFillColor(vgo.RGBA(0, 0, 0, 160))

	if button.State == Default || button.State == Hover {
		ctx.Text(btnx, btny + 1, button.Label)
		ctx.SetFillColor(vgo.RGBA(255, 255, 255, 160))
		ctx.Text(btnx, btny, button.Label)
	} else {
		ctx.Text(btnx, btny, button.Label)
		ctx.SetFillColor(vgo.RGBA(255, 255, 255, 160))
		ctx.Text(btnx, btny + 1, button.Label)
	}
}

func IconToStr(cp int) string {
	return string([]rune{rune(cp)})
}
