package player

import (
	//	"fmt"
	//gl "github.com/go-gl/gl/v4.1-core/gl"
	//mgl32 "github.com/go-gl/mathgl/mgl32"
	actor "github.com/jcecil/avatar/actor"
	input "github.com/jcecil/avatar/input"
)

var (
	k *input.KeyButton
	m *input.MouseButton
	c *input.CursorPosition
)

type Player struct {
	PlayerActor  *actor.Actor
	PlayerCamera *actor.Camera
	PressedW     bool
	PressedA     bool
	PressedS     bool
	PressedD     bool
	CursorXpos   float32
	CursorYpos   float32
	//position        mgl32.Vec3
	//direction       mgl32.Vec3
	//right           mgl32.Vec3
	//up              mgl32.Vec3
	horizontalAngle float32
	verticalAngle   float32
	mouseSpeed      float32
	movementSpeed   float32
}

func NewPlayer() *Player {
	p := new(Player)

	p.playerActor = new(actor.Actor)
	//p.playerActor.position = mgl32.Vec3{0, 0, 0}

	p.playerCamera = new(actor.Camera)
	//p.playerCamera.position = mgl32.Vec3{0, 0, 5}
	//p.playerCamera.horizontalAngle = 3.14
	//p.playerCamera.verticalAngle = 0

	p.mouseSpeed = 0.005
	p.movementSpeed = 0.05

	return p
}

func (*Player) TearDown() {
}

func (p *Player) Draw(program uint32) {

	//fmt.Println(p.direction)
	//projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(768)/1024, 0.1, 10.0)
	p.playerActor.Draw(program)
	p.playerCamera.Draw(program)
	//camera := mgl32.LookAtV(p.position, p.position.Add(p.direction), p.up)
	//cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	//gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

}

func (p *Player) Loop() bool {
	return p.ReadInput()
}
