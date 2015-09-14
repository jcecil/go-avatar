package actor

import ()

// All items in Avatar
type Actor interface {
}

// Characters, Elements, Weapons, etc...
type RenderableActor interface {
	Initialize()
	Draw(uint32)
}

type Player struct {
	//	mesh Mesh
}

func (Player) Draw() {
}
