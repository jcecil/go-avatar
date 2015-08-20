package actor

// All items in Avatar
type Actor interface {
}

// Characters, Elements, Weapons, etc...
type RenderableActor interface {
	Draw()
}

type Player struct {
	mesh Mesh
}

func (Player) Draw() {
}
