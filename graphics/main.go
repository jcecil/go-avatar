package graphics

import (
	fmt "fmt"
	gl "github.com/go-gl/gl/v4.1-core/gl"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"
	mgl32 "github.com/go-gl/mathgl/mgl32"
	game "github.com/jcecil/avatar/game"
	input "github.com/jcecil/avatar/input"
	player "github.com/jcecil/avatar/player"
)

const (
	Title        = "Avatar"
	WindowWidth  = 1024
	WindowHeight = 768
)

var (
	Window *glfw.Window
)

func Main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(WindowWidth, WindowHeight, "Cube", nil, nil)
	Window = window
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure the vertex and fragment shaders
	program, err := newProgram("SimpleVertexShader.vertexshader", "SimpleFragmentShader.fragmentshader")
	if err != nil {
		panic(err)
	}
	gl.UseProgram(program)

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(WindowWidth)/WindowHeight, 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	angle := 0.0
	previousTime := glfw.GetTime()

	width, height := window.GetSize()
	window.SetCursorPos(float64(width/2), float64(height/2))
	window.SetKeyCallback(input.OnKey)
	window.SetCursorPosCallback(input.OnCursor)
	window.SetMouseButtonCallback(input.OnMouse)

	for !player.ShouldClose {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time

		angle += elapsed
		model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

		// Render
		gl.UseProgram(program)
		// gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
		player.MainPlayer.Draw(program)
		for _, element := range game.Universe {
			(element).Draw(program)
		}

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func CloseWindow() {
	Window.SetShouldClose(true)
}
