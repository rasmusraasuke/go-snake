package main

type Player struct {
	name      string
	input     InputType
	bestScore int
}

func NewPlayer(name string, input InputType) *Player {
	return &Player{name: name, input: input, bestScore: 0}
}

func NewComputer() *Player {
	return &Player{name: "Computer", bestScore: 0, input: COMPUTER}
}
