package main

type Player struct {
	name      string
	bestScore int
}

func NewPlayer(name string) *Player {
	return &Player{name: name, bestScore: 0}
}
