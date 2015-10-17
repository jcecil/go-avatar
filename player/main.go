package player

import ()

var (
	MainPlayer  *Player
	ShouldClose bool
)

func Initialize() {
	ShouldClose = false
	MainPlayer = NewPlayer()
}

func TearDown() {
	MainPlayer.TearDown()
}

func Loop() bool {
	return MainPlayer.Loop()
}
