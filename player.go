package main

type Player struct {
	inputType InputType
}

func NewPlayer(input InputType) *Player {
	player := Player{input}
	return &player
}
