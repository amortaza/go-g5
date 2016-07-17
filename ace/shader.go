package ace

import (
	gl "github.com/chsc/gogl/gl33"
	"io/ioutil"
	"github.com/amortaza/go-g5/util"
)

type Shader struct {
	shaderId gl.Uint
}

func NewVertexShader(filename string) *Shader {
	return newShader(filename, gl.VERTEX_SHADER)
}

func NewFragmentShader(filename string) *Shader {
	return newShader(filename, gl.FRAGMENT_SHADER)
}

func newShader(filename string, shaderType gl.Enum) *Shader {
	src, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err.Error())
	}

	shader := &Shader{}

	shader.shaderId = gl.CreateShader(shaderType)

	glSrc, free := util.GLstrs(string(src) + "\x00")
	defer free()

	gl.ShaderSource(shader.shaderId, 1, glSrc, nil)

	gl.CompileShader(shader.shaderId)

	var status gl.Int
	gl.GetShaderiv(shader.shaderId, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		panic("There was a problem with compiling the shader")
	}

	return shader
}

func (s *Shader) Free() {
	gl.DeleteShader(s.shaderId);
}