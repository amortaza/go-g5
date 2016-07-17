package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

func DrawColorRect4f(	left, top, width, height int,
			r,g,b,a float32) {

	color := []float32{r,g,b,a}
	g_colorRect.Draw(left, top, width, height, color, color, color, color, (*gl.Float)(&g_projection[0]))
}

func DrawColorRect3f(	left, top, width, height int,
			r,g,b float32) {

	color := []float32{r,g,b,1}
	g_colorRect.Draw(left, top, width, height, color, color, color, color, (*gl.Float)(&g_projection[0]))
}

func DrawColorRect1v(	left, top, width, height int,
			color []float32) {

	g_colorRect.Draw(left, top, width, height, color, color, color, color, (*gl.Float)(&g_projection[0]))
}

func DrawColorRect4v(	left, top, width, height int,
			leftTopColor []float32,
			rightTopColor []float32,
			rightBottomColor []float32,
			leftBottomColor []float32) {

	g_colorRect.Draw(left, top, width, height, leftTopColor, rightTopColor, rightBottomColor, leftBottomColor, (*gl.Float)(&g_projection[0]))
}

func DrawTextureRect(	texture *Texture,
			left, top, width, height int,
			alphas []float32,) {

	g_textureRect.Draw(texture, left, top, width, height, alphas, &g_projection[0])
}

func DrawTextureRectUpsideDown(	texture *Texture,
			left, top, width, height int,
			alphas []float32,) {

	g_textureRect.DrawUpsideDown(texture, left, top, width, height, alphas, &g_projection[0])
}

func DrawCanvasRect(	canvas *Canvas,
			left, top, width, height int,
			alphas []float32,) {

	g_canvasRect.DrawUpsideDown(canvas.FramebufferSingleSampled.Texture, left, top, width, height, alphas, &g_projection[0])
}

func DrawStringRect(	fontTexture *StringTexture,
			left, top int,
			rgbFg []float32,
			rgbBg []float32,
			alpha float32) {

	g_stringRect.DrawString(fontTexture.Texture, left, top, fontTexture.Texture.Width, fontTexture.Texture.Height, rgbFg, rgbBg, alpha, &g_projection[0])
}

func ClearRect(	width, height int,
		red, green, blue float32 ) {

	g_colorRect.DrawSolid(0, 0, width, height, red, green, blue, (*gl.Float)(&g_projection[0]))
}
