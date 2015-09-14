package graphics

import (
	"fmt"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"
)

type CursorPosition struct {
	xPos float64
	yPos float64
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

func onCursor(window *glfw.Window, xPos float64, yPos float64) {
	select {
	case CursorInput <- &CursorPosition{xPos, yPos}:
	default:
	}
	window.SetCursorPos(600, 600)
}

func onMouse(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	select {
	case MouseInput <- &MouseButton{button, action, mod}:
	default:
	}
}

func onKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	select {
	case KeyInput <- &KeyButton{k, s, action, mods}:
	default:
	}
}
