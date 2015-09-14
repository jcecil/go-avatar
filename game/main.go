package game

import (
	"fmt"
	actor "github.com/jcecil/avatar/actor"
)

var (
	Universe [](actor.RenderableActor)
	Players  []*(actor.Player)
	players  chan int
)

func Initialize() {
	players = make(chan int, 10)
	Universe = append(Universe, &actor.Cube{})
	fmt.Println("Initializing game")
}

func TearDown() {
}

func Loop() bool {
	// Process player input

	// Update game state
	return false
}

func AddPPlayer() {
	p := actor.Player{}
	Players = append(Players, &p)
}
