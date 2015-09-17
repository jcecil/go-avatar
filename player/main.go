package player

import (
	fmt "fmt"
)

var (
	MainPlayer Player
)

func Initialize() {
	MainPlayer.Initialize()
	fmt.Println("Initializing player")
}

func TearDown() {
	MainPlayer.TearDown()
}

func Loop() bool {
	return MainPlayer.Loop()
}
