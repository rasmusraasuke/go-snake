package main

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
)

type Color int

const (
	Green Color = iota
	Yellow
)

type FoodType int

const (
	Cherry FoodType = iota + 1
)
