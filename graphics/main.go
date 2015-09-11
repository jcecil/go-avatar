package graphics

import (
	//	game "bitbucket.org/jcecil/avatar/game"
	fmt "fmt"
	//	"github.com/go-gl/gl"
	gl "github.com/go-gl/gl/v4.1-core/gl"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"

	//	"github.com/go-gl/mathgl/mgl32"
	"runtime"
)

const (
	Title        = "Avatar"
	WindowWidth  = 1024
	WindowHeight = 768
)

//var (
//	window       *glfw.Window
//	err          error
//	rotx, roty   float32
//	vertexBuffer gl.Buffer
//	colorBuffer  gl.Buffer
//	programID    gl.Program
//	matrixID     gl.UniformLocation
//	mvp          mgl32.Mat4
//	colorBufferData []gl.GLfloat
//	loopCount       int        = 1
//	position        mgl32.Vec3 = [...]float32{0, 0, 5}
//	horizontalAngle float32    = 3.14
//	verticalAngle   float32    = 0
//	fieldOfView     float32    = 45
//)

func init() {
	runtime.LockOSThread()
	//	KeyInput = make(chan *KeyButton, 100)
	//	MouseInput = make(chan *MouseButton, 100)
	//	CursorInput = make(chan *CursorPosition, 100)
}

func Main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	//glfw.WindowHint(glfw.ContextVersionMajor, 4)
	//glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(WindowWidth, WindowHeight, "Cube", nil, nil)
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

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

//	if !glfw.Init() {
//		panic("Can't init glfw!")
//	}

// http://www.opengl-tutorial.org/beginners-tutorials/tutorial-1-opening-a-window/
// 4x antialiasing
//	glfw.WindowHint(glfw.Samples, 4)
//	glfw.WindowHint(glfw.ContextVersionMajor, 3)
//	glfw.WindowHint(glfw.ContextVersionMinor, 3)
//	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

//	window, err = glfw.CreateWindow(Width, Height, Title, nil, nil)
//	if err != nil {
//		panic(err)
//	}

//	window.MakeContextCurrent()

//	glfw.SwapInterval(1)

//	gl.Init()

// Triangle draw
//	vertexArray := gl.GenVertexArray()
//	vertexArray.Bind()

// Triangle
//vertexBufferData := []gl.GLfloat{-1.0, -1.0, 0.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0}
// Cube
//	vertexBufferData := []gl.GLfloat{-1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, 1.0, -1.0, 1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, -1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, -1.0, 1.0}
//	colorBufferData = []gl.GLfloat{-1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, 1.0, -1.0, 1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, -1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, -1.0, 1.0}

//	vertexBuffer = gl.GenBuffer()
//	vertexBuffer.Bind(gl.ARRAY_BUFFER)
//	gl.BufferData(gl.ARRAY_BUFFER, len(vertexBufferData)*4, vertexBufferData, gl.STATIC_DRAW)

//	colorBuffer = gl.GenBuffer()
//	colorBuffer.Bind(gl.ARRAY_BUFFER)
//	gl.BufferData(gl.ARRAY_BUFFER, len(colorBufferData)*4, colorBufferData, gl.STATIC_DRAW)

//	gl.ClearColor(0.0, 0.0, 0.3, 0.0)

//	programID = LoadShaders("SimpleVertexShader.vertexshader", "SimpleFragmentShader.fragmentshader")
//	matrixID = programID.GetUniformLocation("MVP")
//	matrixID.UniformMatrix4f(false, &mvp)

//Callbacks
//	window.SetKeyCallback(onKey)
//	window.SetCursorPositionCallback(onCursor)
//	window.SetMouseButtonCallback(onMouse)
//	fmt.Println("done!!!!!")
//}

//func Loop() bool {
//	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
//	gl.Enable(gl.DEPTH_TEST)
//	gl.DepthFunc(gl.LESS)

// Model transformation
//	model := mgl32.Ident4()

// View transformation
//	var eye mgl32.Vec3 = [...]float32{4, 3, 3}
//	var center mgl32.Vec3 = [...]float32{0, 0, 0}
//	var up mgl32.Vec3 = [...]float32{0, 1, 0}
//	view := mgl32.LookAtV(eye, center, up)

// Projection transformation
//	projection := mgl32.Perspective(45, 4.0/3.0, .1, 100)
//
//	mvp = projection.Mul4(view).Mul4(model)

//	programID.Use()
//	matrixID.UniformMatrix4fv(false, mvp)

//	var vertexAttrib gl.AttribLocation = 0
//	vertexAttrib.EnableArray()
//	vertexBuffer.Bind(gl.ARRAY_BUFFER)
//	vertexAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

//	colorBuffer.Bind(gl.ARRAY_BUFFER)
//	loopCount = (loopCount + 1) % len(colorBufferData)
//	for i := range colorBufferData {
//		colorBufferData[i] = gl.GLfloat((float32((i + loopCount) % len(colorBufferData))) / float32(len(colorBufferData)))
//	}
//	gl.BufferData(gl.ARRAY_BUFFER, len(colorBufferData)*4, colorBufferData, gl.STATIC_DRAW)

//	var colorAttrib gl.AttribLocation = 1
//	colorAttrib.EnableArray()
//	colorBuffer.Bind(gl.ARRAY_BUFFER)
//	colorAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

//gl.DrawArrays(gl.TRIANGLES, 0, 3)
//	gl.DrawArrays(gl.TRIANGLES, 0, 36)
//	vertexAttrib.DisableArray()

//	glfw.PollEvents()
//	window.SwapBuffers()
//	return window.ShouldClose()
//}

//func GetWindow() *glfw.Window {
//	return window
//}
