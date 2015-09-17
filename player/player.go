package player

import (
	gl "github.com/go-gl/gl/v4.1-core/gl"
	mgl32 "github.com/go-gl/mathgl/mgl32"
	input "github.com/jcecil/avatar/input"
	math "math"
)

var (
	k *input.KeyButton
	m *input.MouseButton
	c *input.CursorPosition
)

type Player struct {
	position        mgl32.Vec3
	direction       mgl32.Vec3
	right           mgl32.Vec3
	up              mgl32.Vec3
	horizontalAngle float32
	verticalAngle   float32
}

func (p Player) Initialize() {
	p.position = mgl32.Vec3{3, 3, 3}
	p.direction = mgl32.Vec3{0, 0, 0}
	p.right = mgl32.Vec3{0, 0, 0}
	p.up = mgl32.Vec3{0, 1, 0}
	p.horizontalAngle = 3.14
	p.verticalAngle = 0
}

func (Player) TearDown() {
}

func (p Player) Draw(program uint32) {

	//projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(768)/1024, 0.1, 10.0)
	camera := mgl32.LookAtV(p.position, p.position.Add(p.direction), p.up)
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

}

func (Player) Loop() bool {
	return ReadInput()
}

func cosine(f float32) float32 {
	return float32(math.Cos(float64(f)))
}
func sine(f float32) float32 {
	return float32(math.Sin(float64(f)))
}
