package main

import (
	game "github.com/jcecil/avatar/game"
	opengl "github.com/jcecil/avatar/opengl"
	views "github.com/jcecil/avatar/views"
	"fmt"
	"time"
)

var (
	exit bool
)

func main() {
	fmt.Println("Avatar!")
	exit = false
	game.Initialize()
	defer game.TearDown()

	views.Initialize()
	defer views.TearDown()

	opengl.Initialize()
	defer opengl.TearDown()

	// OpenGL code has to all occur on the main thread
	// Create a go channel to run the other parts of our code
	go Loop()
	for !exit {
		opengl.Loop()
		time.Sleep(10 * time.Millisecond)

	}
}

func Loop() {
	for !exit {
		exit = views.Loop()
		game.Loop()
		time.Sleep(10 * time.Millisecond)
	}
}
