package opengl

import (
	"bufio"
	"bytes"
	"fmt"
	gl "github.com/go-gl/gl"
	"os"
	"os/exec"
)

func readFile(filename string) string {
	var fileText bytes.Buffer
	fullname, _ := exec.LookPath("avatar")
	fullname = fullname[:len(fullname)-6] + "/shaders/" + filename
	fmt.Println(fullname)
	file, _ := os.Open(fullname)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileText.WriteString(scanner.Text())
		fileText.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error: ", err)
	}

	file.Close()
	return fileText.String()
}

func LoadShaders(vertexPath, fragmentPath string) gl.Program {
	// Create the shaders
	vertexShaderID := gl.CreateShader(gl.VERTEX_SHADER)
	fragmentShaderID := gl.CreateShader(gl.FRAGMENT_SHADER)

	// Read the Vertex Shader code from the file
	vertexShaderCode := readFile(vertexPath)
	vertexShaderID.Source(vertexShaderCode)
	vertexShaderID.Compile()
	defer vertexShaderID.Delete()

	// Read the Fragment Shader code from the file
	fragmentShaderCode := readFile(fragmentPath)
	fragmentShaderID.Source(fragmentShaderCode)
	fragmentShaderID.Compile()
	defer fragmentShaderID.Delete()

	program := gl.CreateProgram()
	program.AttachShader(vertexShaderID)
	program.AttachShader(fragmentShaderID)
	program.Link()
	return program
}
