package player

import (
	fmt "fmt"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"
	mgl32 "github.com/go-gl/mathgl/mgl32"
	input "github.com/jcecil/avatar/input"
)

func ReadInput() bool {
	var mouseSpeed float32 = .005
	var deltaTime float32 = 1
	doSomething := true

	for doSomething {
		select {
		case k = <-input.KeyInput:
			fmt.Println("Keyboard input", *k)

			switch glfw.Key((*k).Button) {
			case glfw.KeyEscape:
				fmt.Println("Escape key pressed. Goodbye!!")
				//graphics.Window.SetShouldClose(true)
				return true
			case glfw.KeyW:
				MainPlayer.position.Add(MainPlayer.direction)
			case glfw.KeyA:
			case glfw.KeyS:
			case glfw.KeyD:
			}
		case m = <-input.MouseInput:
			fmt.Println("Mouse input:", *m)

		case c = <-input.CursorInput:
			fmt.Println("Cursor movement:", *c)
			MainPlayer.horizontalAngle += mouseSpeed * deltaTime * (1024/2 - c.Xpos)
			MainPlayer.verticalAngle += mouseSpeed * deltaTime * (768/2 - c.Ypos)
			MainPlayer.direction = mgl32.Vec3{cosine(MainPlayer.verticalAngle) * sine(MainPlayer.horizontalAngle), sine(MainPlayer.verticalAngle), cosine(MainPlayer.verticalAngle) * cosine(MainPlayer.horizontalAngle)}
			MainPlayer.right = mgl32.Vec3{sine(MainPlayer.horizontalAngle - 2.14/2.0), 0, cosine(MainPlayer.horizontalAngle - 3.14/2.0)}
			MainPlayer.up = MainPlayer.right.Cross(MainPlayer.direction)

		default:
			doSomething = false
		}
	}
	return false
}
