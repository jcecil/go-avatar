package player

import (
	"fmt"
	gl "github.com/go-gl/gl/v4.1-core/gl"
	mgl32 "github.com/go-gl/mathgl/mgl32"
	input "github.com/jcecil/avatar/input"
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
	mouseSpeed      float32
	movementSpeed   float32
}

func (p *Player) Printme() {
	fmt.Println(p.direction)
}

func NewPlayer() *Player {
	p := new(Player)
	p.position = mgl32.Vec3{0, 0, 5}
	p.horizontalAngle = 3.14
	p.verticalAngle = 0
	p.mouseSpeed = 0.005
	p.movementSpeed = 0.05

	p.updateCamera()
	return p
}

func (*Player) TearDown() {
}

func (p *Player) updateCamera() {
	p.direction = mgl32.Vec3{cosine(p.verticalAngle) * sine(p.horizontalAngle), sine(p.verticalAngle), cosine(p.verticalAngle) * cosine(p.horizontalAngle)}
	p.right = mgl32.Vec3{sine(p.horizontalAngle - 2.14/2.0), 0, cosine(p.horizontalAngle - 3.14/2.0)}
	p.up = p.right.Cross(p.direction)
}

func (p *Player) Draw(program uint32) {

	//fmt.Println(p.direction)
	//projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(768)/1024, 0.1, 10.0)
	camera := mgl32.LookAtV(p.position, p.position.Add(p.direction), p.up)
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

}

func (p *Player) Loop() bool {
	return p.ReadInput()
}
