package g5

import (
	gl "github.com/chsc/gogl/gl33"
	"io/ioutil"
)

type _Shader struct {
	shaderId gl.Uint
}

func newVertexShader(filename string) *_Shader {

	return newShader(filename, gl.VERTEX_SHADER)
}

func newFragmentShader(filename string) *_Shader {

	return newShader(filename, gl.FRAGMENT_SHADER)
}

func newShader(filename string, shaderType gl.Enum) *_Shader {

	src, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err.Error())
	}

	shader := &_Shader{}

	shader.shaderId = gl.CreateShader(shaderType)

	glSrc, free := glStrs(string(src) + "\x00")
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

func (s *_Shader) Free() {
	gl.DeleteShader(s.shaderId);
}