package actor

import (
	mgl32 "github.com/go-gl/mathgl/mgl32"
)

// All items in Avatar
type Actor struct {
	direction       mgl32.Vec3
	right           mgl32.Vec3
	up              mgl32.Vec3
	position        mgl32.Vec3
	velocity        mgl32.Vec3
	acceleration    mgl32.Vec3
	renderableActor RenderableActor
}

// Characters, Elements, Weapons, etc...
type RenderableActor interface {
	Initialize()
	Draw(uint32)
}

func (a *Actor) Draw(Program uint32) {
	a.renderableActor.Draw(Program)
}
