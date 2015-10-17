package player

import ()

var (
	MainPlayer *Player
)

func Initialize() {
	MainPlayer = NewPlayer()
}

func TearDown() {
	MainPlayer.TearDown()
}

func Loop() bool {
	return MainPlayer.Loop()
}
