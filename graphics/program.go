package graphics

import (
	//	"bytes"
	errors "errors"
	fmt "fmt"
	gl "github.com/go-gl/gl/v4.1-core/gl"
	shaders "github.com/jcecil/avatar/graphics/shaders"
	strings "strings"
)

var vertexShaderSource string = `
#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;
void main() {
	    fragTexCoord = vertTexCoord;
	        gl_Position = projection * camera * model * vec4(vert, 1);
	}
	` + "\x00"

var fragmentShaderSource = `
#version 330
uniform sampler2D tex;
in vec2 fragTexCoord;
out vec4 outputColor;
void main() {
	    outputColor = texture(tex, fragTexCoord);
    }
    ` + "\x00"

func newProgram(vertexShaderFilename, fragmentShaderFilename string) (uint32, error) {
	program := gl.CreateProgram()

	//vertexShaderSource, err := shaders.Asset(vertexShaderFilename)
	//if err != nil {
	//	return 0, err
	//}

	vertexShader, err := shaders.Compile(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	//fragmentShaderSource, err := shaders.Asset(fragmentShaderFilename)
	//if err != nil {
	//	return 0, err
	//}
	fragmentShader, err := shaders.Compile(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, errors.New(fmt.Sprintf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}
