package main

type State int

const (
	MAIN_MENU State = iota
	GAME
	GAME_OVER
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type InputType int

const (
	WASD InputType = iota
	ARROWS
	VIM
	COMPUTER
)

type FoodType int

const (
	Cherry FoodType = iota + 1
)
