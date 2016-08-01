package g5

import (
	gl "github.com/chsc/gogl/gl33"
	"github.com/amortaza/go-g5/util"
	"github.com/amortaza/go-g5/shader"
)

type ColorRect struct {
	program *shader.Program

	vao gl.Uint
	vbo gl.Uint
}

func NewColorRect() *ColorRect {
	r := &ColorRect{}

	r.program = shader.NewProgram("github.com/amortaza/go-g5/shader/rgb.vertex.txt", "github.com/amortaza/go-g5/shader/rgb.fragment.txt")

	gl.GenVertexArrays(1, &r.vao)
	gl.GenBuffers(1, &r.vbo)

	return r
}

func (r *ColorRect) Draw(	left, top, width, height int,
				leftTopColor []float32,
				rightTopColor []float32,
				rightBottomColor []float32,
				leftBottomColor []float32,
				projection *gl.Float ) {

	r.program.Activate()

	gl.UniformMatrix4fv(r.program.GetUniformLocation("project"), 1, 0, projection)

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	right := left + width
	bottom := top + height

	vertices := []float32{
		float32(left), float32(top), leftTopColor[0], leftTopColor[1], leftTopColor[2], leftTopColor[3],
			float32(right), float32(top), rightTopColor[0], rightTopColor[1], rightTopColor[2], rightTopColor[3],
				float32(right), float32(bottom), rightBottomColor[0], rightBottomColor[1], rightBottomColor[2], rightBottomColor[3],
					float32(left), float32(bottom), leftBottomColor[0], leftBottomColor[1], leftBottomColor[2], leftBottomColor[3] }

	colorRect_setVertexData(vertices)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.BindVertexArray(0)
}

func (r *ColorRect) DrawSolid(	left, top, width, height int,
				red, green, blue float32,
				projection *gl.Float ) {

	r.program.Activate()

	gl.UniformMatrix4fv(r.program.GetUniformLocation("project"), 1, 0, projection)

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	right := left + width
	bottom := top + height

	vertices := []float32{
		float32(left), float32(top), red, green, blue, 1,
			float32(right), float32(top), red, green, blue, 1,
				float32(right), float32(bottom), red, green, blue, 1,
					float32(left), float32(bottom), red, green, blue, 1 }

	colorRect_setVertexData(vertices)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.BindVertexArray(0)
}

func (r *ColorRect) Free() {
	gl.DeleteVertexArrays(1, &r.vao)
	gl.DeleteBuffers(1, &r.vbo)

	r.program.Free()
}

func colorRect_setVertexData(data []float32) {

	// copy vertices data into VBO (it needs to be bound first)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(data)*4), gl.Pointer(util.GLptr(data)), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 2 /*posPartCount*/ *4 + 4 /*colorPartCount*/ *4
	var offset int = 0

	// position
	gl.VertexAttribPointer(0, 2 /*posPartCount*/, gl.FLOAT, 0, gl.Sizei(stride), gl.Pointer(util.GLptrOffset(offset)))
	gl.EnableVertexAttribArray(0)
	offset += 2 /*posPartCount*/ * 4

	// color
	gl.VertexAttribPointer(1, 4 /*colorPartCount*/, gl.FLOAT, 0, gl.Sizei(stride), gl.Pointer(util.GLptrOffset(offset)))
	gl.EnableVertexAttribArray(1)
}
