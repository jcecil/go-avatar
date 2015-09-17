package input

import (
	"fmt"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"
	runtime "runtime"
)

func init() {
	runtime.LockOSThread()
	KeyInput = make(chan *KeyButton, 100)
	MouseInput = make(chan *MouseButton, 100)
	CursorInput = make(chan *CursorPosition, 100)
	//WindowConfigure = make(chan bool)
}

type CursorPosition struct {
	Xpos float32
	Ypos float32
}

type KeyButton struct {
	Button   glfw.Key
	scancode int
	action   glfw.Action
	mod      glfw.ModifierKey
}

type MouseButton struct {
	button glfw.MouseButton
	action glfw.Action
	mod    glfw.ModifierKey
}

var (
	MouseInput  chan *MouseButton
	KeyInput    chan *KeyButton
	CursorInput chan *CursorPosition
)

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}

func OnCursor(window *glfw.Window, xPos float64, yPos float64) {
	select {
	case CursorInput <- &CursorPosition{float32(xPos), float32(yPos)}:
	default:
	}
	window.SetCursorPos(1024/2, 768/2)
}

func OnMouse(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	select {
	case MouseInput <- &MouseButton{button, action, mod}:
	default:
	}
}

func OnKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	select {
	case KeyInput <- &KeyButton{k, s, action, mods}:
	default:
	}
}
