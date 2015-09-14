package views

import (
	fmt "fmt"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"
	graphics "github.com/jcecil/avatar/graphics"
)

var (
	k *graphics.KeyButton
	m *graphics.MouseButton
	c *graphics.CursorPosition
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
	select {
	case k = <-graphics.KeyInput:
		fmt.Println("Keyboard input", *k)

		switch glfw.Key((*k).Button) {
		case glfw.KeyEscape:
			fmt.Println("Escape key pressed. Goodbye!!")
			graphics.Window.SetShouldClose(true)
			return true
		}
	case m = <-graphics.MouseInput:
		fmt.Println("Mouse input:", *m)
	case c = <-graphics.CursorInput:
		fmt.Println("Cursor movement:", *c)
	default:

	}
	return false
}
