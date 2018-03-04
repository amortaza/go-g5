package g5

import (
	gl "github.com/chsc/gogl/gl33"
)

type _Program struct {

	programId gl.Uint

	vShader, fShader *_Shader
}

func newProgram(vertexFilename, fragmentFilename string) *_Program {

	p := &_Program{}

	p.vShader = newVertexShader(vertexFilename)
	p.fShader = newFragmentShader(fragmentFilename)

	p.programId = gl.CreateProgram()

	gl.AttachShader(p.programId, p.vShader.shaderId)
	gl.AttachShader(p.programId, p.fShader.shaderId)

	gl.LinkProgram(p.programId)

	return p
}

func (p *_Program) Activate() {

	gl.UseProgram(p.programId);
}

func (p *_Program) Free() {

	gl.DetachShader(p.programId, p.vShader.shaderId)
	gl.DetachShader(p.programId, p.fShader.shaderId)

	p.vShader.Free()
	p.fShader.Free()

	gl.DeleteProgram(p.programId)
}

func (p *_Program) GetUniformLocation(name string) gl.Int {

	return gl.GetUniformLocation(p.programId, glStr(name))
}


