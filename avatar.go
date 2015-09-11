package main

import (
	fmt "fmt"
	game "github.com/jcecil/avatar/game"
	graphics "github.com/jcecil/avatar/graphics"
	views "github.com/jcecil/avatar/views"
	time "time"
)

func main() {
	fmt.Println("Avatar!")

	game.Initialize()
	defer game.TearDown()

	views.Initialize()
	defer views.TearDown()

	go Loop()

	graphics.Main()

}

func Loop() {
	exit := false
	for !exit {
		//exit = views.Loop()
		//game.Loop()
		time.Sleep(10 * time.Millisecond)
	}
}
