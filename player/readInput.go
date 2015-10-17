package player

import (
	fmt "fmt"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"
	mgl32 "github.com/go-gl/mathgl/mgl32"
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
				//graphics.Window.SetShouldClose(true)
				return true
			case glfw.KeyW:
				p.position = p.position.Add(p.direction.Mul(p.movementSpeed))
			case glfw.KeyA:
				p.position = p.position.Sub(p.right.Mul(p.movementSpeed))
			case glfw.KeyS:
				p.position = p.position.Sub(p.direction.Mul(p.movementSpeed))
			case glfw.KeyD:
				p.position = p.position.Add(p.right.Mul(p.movementSpeed))
			}
		case m = <-input.MouseInput:
			fmt.Println("Mouse input:", *m)

		case c = <-input.CursorInput:
			fmt.Println("Cursor movement:", *c)
			fmt.Println("Mouse speed:", p.mouseSpeed)
			p.horizontalAngle += p.mouseSpeed * deltaTime * (1024/2 - c.Xpos)
			p.verticalAngle += p.mouseSpeed * deltaTime * (768/2 - c.Ypos)
			p.direction = mgl32.Vec3{cosine(p.verticalAngle) * sine(p.horizontalAngle), sine(p.verticalAngle), cosine(p.verticalAngle) * cosine(p.horizontalAngle)}
			fmt.Println("direction:", p.direction)
			fmt.Println("direction2:", MainPlayer.direction)
			p.right = mgl32.Vec3{sine(p.horizontalAngle - 2.14/2.0), 0, cosine(p.horizontalAngle - 3.14/2.0)}
			p.up = p.right.Cross(p.direction)

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
