package opengl

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
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
	CursorInput <- &CursorPosition{xPos, yPos}
	window.SetCursorPosition(600, 600)
}

func onMouse(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	MouseInput <- &MouseButton{button, action, mod}
}

func onKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	KeyInput <- &KeyButton{k, s, action, mods}
}
