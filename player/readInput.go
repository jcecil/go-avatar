package player

import (
	fmt "fmt"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"
	input "github.com/jcecil/avatar/input"
	math "math"
)

func (p *Player) ReadInput() bool {
	var deltaTime float32 = 1
	doSomething := true

	for doSomething {
		select {
		case k = <-input.KeyInput:
			fmt.Println("Keyboard input", *k)

			switch glfw.Key((*k).Button) {
			case glfw.KeyEscape:
				fmt.Println("Escape key pressed. Goodbye!!")
				ShouldClose = true
				return true
			case glfw.KeyW:
				p.PressedW = true
				//p.position = p.position.Add(p.direction.Mul(p.movementSpeed))
			case glfw.KeyA:
				p.PressedA = true
				//p.position = p.position.Sub(p.right.Mul(p.movementSpeed))
			case glfw.KeyS:
				p.PressedS = true
				//p.position = p.position.Sub(p.direction.Mul(p.movementSpeed))
			case glfw.KeyD:
				p.PressedD = true
				//p.position = p.position.Add(p.right.Mul(p.movementSpeed))
			}
		case m = <-input.MouseInput:
			fmt.Println("Mouse input:", *m)

		case c = <-input.CursorInput:
			fmt.Println("Cursor movement:", *c)
			//			p.PlayerCamera.horizontalAngle += p.mouseSpeed * deltaTime * (1023/2 - c.Xpos)
			//			p.PlayerCamera.verticalAngle += p.mouseSpeed * deltaTime * (768/2 - c.Ypos)
			//p.CursorXpos = c.Xpos
			//p.CursorYpos = c.Ypos

			//p.updateCamera()

		default:
			doSomething = false
		}
	}
	return false
}

func cosine(f float32) float32 {
	return float32(math.Cos(float64(f)))
}
func sine(f float32) float32 {
	return float32(math.Sin(float64(f)))
}
