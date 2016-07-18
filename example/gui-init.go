package gui

import "github.com/shibukawa/nanovgo"

var g_style Style
var Icon IconTable
var ctx *nanovgo.Context

func Init(ctx2 *nanovgo.Context) {
	ctx = ctx2

	g_style.HoverColor = nanovgo.RGBA(210, 210, 10, 168)

	Icon.Trash = IconToStr(0xE729)
}

