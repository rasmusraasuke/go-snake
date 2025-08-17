package main

import (
	"errors"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Coordinate struct {
	X, Y int
}

type SnakeGame struct {
	snakeGrid SnakeGrid
	snake     Snake
	food      map[Coordinate]Food
	wait      int
}

func NewSnakeGame() *SnakeGame {
	startX := int(GRID_SIZE / 2)
	startY := int(GRID_SIZE / 2)
	startOrient := rand.IntN(4)

	grid := NewSnakeGrid()
	snake := NewSnake(startX, startY, Direction(startOrient))
	food := make(map[Coordinate]Food)
	wait := 0

	game := SnakeGame{*grid, *snake, food, wait}
	return &game
}

func (g *SnakeGame) Update() error {
	g.snake.UpdatePendingOrientation()

	if len(g.food) == 0 {
		randX := rand.IntN(GRID_SIZE)
		randY := rand.IntN(GRID_SIZE)
		g.food[Coordinate{randX, randY}] = *NewFood(Cherry, randX, randY)
	}

	if g.wait < 10 {
		g.wait++
		return nil
	}
	g.wait = 0

	newX, newY := g.snake.CalculateNextTile()
	if newX < 0 || newY < 0 || newX >= GRID_SIZE || newY >= GRID_SIZE {
		return errors.New("Snake hit it's head against the wall!")
	}

	for i, bodyElement := range g.snake.Body {
		if bodyElement.XPos == newX && bodyElement.YPos == newY && i != len(g.snake.Body)-1 {
			return errors.New("Snake hit itself!")
		}
	}

	coordiate := Coordinate{newX, newY}
	switch g.food[coordiate].Type {
	case 0:
		g.snake.Move(newX, newY)
	case Cherry:
		g.snake.EatCherry(newX, newY)
		delete(g.food, coordiate)
	}

	return nil
}

func (g *SnakeGame) Draw(screen *ebiten.Image) {
	g.snakeGrid.Draw(screen)

	for _, food := range g.food {
		food.Draw(screen)
	}

	g.snake.Draw(screen)
}
