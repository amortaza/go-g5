package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

type TextureRect struct {
	program *Program

	vao gl.Uint
	vbo gl.Uint
}

func NewTextureRect(vertexShaderFilename, fragmentShaderFilename string) *TextureRect {
	r := &TextureRect{}

	r.program = NewProgram(vertexShaderFilename, fragmentShaderFilename)

	gl.GenVertexArrays(1, &r.vao)
	gl.GenBuffers(1, &r.vbo)

	return r
}

func (r *TextureRect) Draw(	texture *Texture,
				left, top, width, height int,
				leftTopRightBottomAlphas []float32,
				projection *float32 ) {

	r.program.Activate()

	texture.Activate(gl.TEXTURE0)

	gl.Uniform1i(r.program.GetUniformLocation("Sampler"), 0)
	gl.UniformMatrix4fv(r.program.GetUniformLocation("Projection"), 1, gl.FALSE, (*gl.Float)(projection))
	gl.Uniform4f(r.program.GetUniformLocation("Alphas"),
		gl.Float(leftTopRightBottomAlphas[0]),
		gl.Float(leftTopRightBottomAlphas[1]),
		gl.Float(leftTopRightBottomAlphas[2]),
		gl.Float(leftTopRightBottomAlphas[3]));

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	right := left + width
	bottom := top + height

	vertices := []float32{
		float32(left), float32(top), 0.0, 0.0,
		float32(right), float32(top), 1.0, 0.0 ,
		float32(right), float32(bottom), 1.0, 1.0,
		float32(left), float32(bottom), 0.0, 1.0 }

	setVertexData2(vertices)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.BindVertexArray(0)

	texture.Deactivate()
}

func (r *TextureRect) DrawUpsideDown(	texture *Texture,
					left, top, width, height int,
					leftTopRightBottomAlphas []float32,
					projection *float32 ) {

	r.program.Activate()

	texture.Activate(gl.TEXTURE0)

	gl.Uniform1i(r.program.GetUniformLocation("Sampler"), 0)
	gl.UniformMatrix4fv(r.program.GetUniformLocation("Projection"), 1, gl.FALSE, (*gl.Float)(projection))
	gl.Uniform4f(r.program.GetUniformLocation("Alphas"),
		gl.Float(leftTopRightBottomAlphas[0]),
		gl.Float(leftTopRightBottomAlphas[1]),
		gl.Float(leftTopRightBottomAlphas[2]),
		gl.Float(leftTopRightBottomAlphas[3]));

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	right := left + width
	bottom := top + height

	vertices := []float32{
		float32(left), float32(top), 0.0, 1.0,
		float32(right), float32(top), 1.0, 1.0 ,
		float32(right), float32(bottom), 1.0, 0.0,
		float32(left), float32(bottom), 0.0, 0.0 }

	setVertexData2(vertices)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.BindVertexArray(0)

	texture.Deactivate()
}

func (r *TextureRect) Free() {
	gl.DeleteVertexArrays(1, &r.vao)
	gl.DeleteBuffers(1, &r.vbo)

	r.program.Free()
}

func setVertexData2(data []float32) {

	// copy vertices data into VBO (it needs to be bound first)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(data)*4), gl.Pointer(GLptr(data)), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride gl.Sizei = 2 /*posPartCount*/ *4 + 2 /*texPartCount*/ *4
	var offset int = 0

	// position
	gl.VertexAttribPointer(0, 2 /*posPartCount*/, gl.FLOAT, gl.FALSE, stride, gl.Pointer(GLptrOffset(offset)))
	gl.EnableVertexAttribArray(0)
	offset += 2 /*posPartCount*/ * 4

	// texture
	gl.VertexAttribPointer(1, 2 /*texPartCount*/, gl.FLOAT, gl.FALSE, stride, gl.Pointer(GLptrOffset(offset)))
	gl.EnableVertexAttribArray(1)
}
