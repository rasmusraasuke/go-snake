package main

type Player struct {
	game      SnakeGame
	inputType InputType
}

func NewPlayer(input InputType) *Player {
	game := NewSnakeGame(input)

	player := Player{game: *game}
	return &player
}
