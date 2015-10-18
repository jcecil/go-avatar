package game

import (
	"fmt"
	actor "github.com/jcecil/avatar/actor"
	player "github.com/jcecil/avatar/player"
)

var (
	Universe [](actor.RenderableActor)
	Players  []*(player.Player)
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
	p := player.NewPlayer()
	Players = append(Players, &p)
}
