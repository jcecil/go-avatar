package actor

import (
	gl "github.com/go-gl/gl/v4.1-core/gl"
	mgl32 "github.com/go-gl/mathgl/mgl32"
	"math"
)

type Camera struct {
	position        mgl32.Vec3
	direction       mgl32.Vec3
	right           mgl32.Vec3
	up              mgl32.Vec3
	HorizontalAngle float32
	VerticalAngle   float32
}

func (c *Camera) UpdateCamera() {
	c.direction = mgl32.Vec3{cosine(c.verticalAngle) * sine(c.horizontalAngle), sine(c.verticalAngle), cosine(c.verticalAngle) * cosine(c.horizontalAngle)}
	c.right = mgl32.Vec3{sine(c.horizontalAngle - 2.14/2.0), 0, cosine(c.horizontalAngle - 3.14/2.0)}
	c.up = c.right.Cross(c.direction)
}

func (c *Camera) Draw(program uint32) {
	//c.UpdateCamera()

	camera := mgl32.LookAtV(c.position, c.position.Add(c.direction), c.up)
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

}

func cosine(f float32) float32 {
	return float32(math.Cos(float64(f)))
}
func sine(f float32) float32 {
	return float32(math.Sin(float64(f)))
}
