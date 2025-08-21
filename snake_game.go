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
	player Player
	screen *ebiten.Image
	snake  Snake
	food   map[Coordinate]Food
	wait   int
}

func NewSnakeGame(player Player) *SnakeGame {
	startX := int(GRID_SIZE / 2)
	startY := int(GRID_SIZE / 2)
	startOrient := rand.IntN(4)

	screen := ebiten.NewImage(GRID_SIZE*TILE_SIZE, GRID_SIZE*TILE_SIZE)
	snake := NewSnake(startX, startY, Direction(startOrient), player.inputType)
	food := make(map[Coordinate]Food)
	wait := 0

	game := SnakeGame{player, screen, *snake, food, wait}
	return &game
}

func (g *SnakeGame) Update() error {
	g.snake.UpdateMovementQueue()

	if len(g.food) == 0 {
		randX := rand.IntN(GRID_SIZE)
		randY := rand.IntN(GRID_SIZE)
		g.food[Coordinate{randX, randY}] = *NewFood(Cherry, randX, randY)
	}

	if g.wait < WAIT_TIME {
		g.wait++
		return nil
	}
	g.wait = 0

	newX, newY := g.snake.CalculateNextPos()
	if newX < 0 || newY < 0 || newX >= GRID_SIZE || newY >= GRID_SIZE {
		return errors.New("Snake hit it's head against the wall!")
	}

	for i, bodyElement := range g.snake.Body {
		if i != len(g.snake.Body)-1 && bodyElement.XPos == newX && bodyElement.YPos == newY {
			return errors.New("Snake hit itself!")
		}
	}

	coordiate := Coordinate{int(newX), int(newY)}
	switch g.food[coordiate].Type {
	case 0:
		g.snake.Move(newX, newY)
	case Cherry:
		g.snake.EatCherry(newX, newY)
		delete(g.food, coordiate)
	}

	return nil
}

func (g *SnakeGame) GetImage() *ebiten.Image {
	g.screen.DrawImage(snakeGrid, &ebiten.DrawImageOptions{})

	for _, food := range g.food {
		food.Draw(g.screen)
	}
	g.snake.Draw(g.screen)

	return g.screen
}
