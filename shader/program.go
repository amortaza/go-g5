package shader

import (
	gl "github.com/chsc/gogl/gl33"
	"github.com/amortaza/go-g5/util"
)

type Program struct {
	programId gl.Uint

	vShader, fShader *Shader
}

func NewProgram(vertexFilename, fragmentFilename string) *Program {
	p := &Program{}

	p.vShader = NewVertexShader(vertexFilename)
	p.fShader = NewFragmentShader(fragmentFilename)

	p.programId = gl.CreateProgram()

	gl.AttachShader(p.programId, p.vShader.shaderId)
	gl.AttachShader(p.programId, p.fShader.shaderId)

	gl.LinkProgram(p.programId)

	return p
}

func (p *Program) Activate() {
	gl.UseProgram(p.programId);
}

func (p *Program) Free() {
	gl.DetachShader(p.programId, p.vShader.shaderId)
	gl.DetachShader(p.programId, p.fShader.shaderId)

	p.vShader.Free()
	p.fShader.Free()

	gl.DeleteProgram(p.programId)
}

func (p *Program) GetUniformLocation(name string) gl.Int {
	return gl.GetUniformLocation(p.programId, util.GLstr(name))
}


