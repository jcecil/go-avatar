package main

import (
	fmt "fmt"
	game "github.com/jcecil/avatar/game"
	graphics "github.com/jcecil/avatar/graphics"
	player "github.com/jcecil/avatar/player"
	//	time "time"
)

func main() {
	fmt.Println("Avatar!")

	game.Initialize()
	defer game.TearDown()

	player.Initialize()
	defer player.TearDown()

	go Loop()

	graphics.Main()
}

func Loop() {
	exit := false
	for !exit {
		exit = player.Loop()
		exit = exit || game.Loop()
		//time.Sleep(10 * time.Millisecond)
	}
}
