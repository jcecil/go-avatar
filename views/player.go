package views

import (
	opengl "github.com/jcecil/avatar/opengl"
	"fmt"
	glfw "github.com/go-gl/glfw3"
)

type PlayerView struct {
	id int
}

func (PlayerView) Initialize() {
}

func (PlayerView) TearDown() {
}

func (PlayerView) Loop() bool {
	return ReadInput()
}

func ReadInput() bool {
	var k *opengl.KeyButton
	var m *opengl.MouseButton
	var c *opengl.CursorPosition
	select {
	case k = <-opengl.KeyInput:
		fmt.Println("Keyboard input", *k)
		switch glfw.Key((*k).Button) {
		case glfw.KeyEscape:
			fmt.Println("Escape key pressed. Goodbye!!")
			return true
		}
	case m = <-opengl.MouseInput:
		fmt.Println("Mouse input:", *m)
	case c = <-opengl.CursorInput:
		fmt.Println("Cursor movement:", *c)
	default:

	}
	return false
}
