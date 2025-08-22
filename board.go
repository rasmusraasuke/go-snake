package main

import (
	"errors"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	game   *Game
	screen *ebiten.Image
	snake  *Snake
	food   map[Coordinate]Food
	wait   int
}

func NewBoard(game *Game, snake *Snake) *Board {
	screen := ebiten.NewImage(GRID_SIZE*TILE_SIZE, GRID_SIZE*TILE_SIZE)
	food := make(map[Coordinate]Food)
	wait := 0

	board := &Board{game, screen, snake, food, wait}
	return board
}

func (b *Board) Update() error {
	b.snake.UpdateMovementQueue()

	if len(b.food) == 0 {
		randX := rand.IntN(GRID_SIZE)
		randY := rand.IntN(GRID_SIZE)
		b.food[Coordinate{randX, randY}] = *NewFood(Cherry, randX, randY)
	}

	if b.wait < WAIT_TIME {
		b.wait++
		return nil
	}
	b.wait = 0

	newX, newY := b.snake.CalculateNextPos()
	if newX < 0 || newY < 0 || newX >= GRID_SIZE || newY >= GRID_SIZE {
		return errors.New("Snake hit it's head against the wall!")
	}

	for i, bodyElement := range b.snake.body {
		if i != len(b.snake.body)-1 && bodyElement.xPos == newX && bodyElement.yPos == newY {
			return errors.New("Snake hit itself!")
		}
	}

	coordiate := Coordinate{int(newX), int(newY)}

	b.snake.Move(newX, newY)
	switch b.food[coordiate].Type {
	case 0:

	case Cherry:
		b.game.FeedOtherSnake(b.snake.playerName)
		delete(b.food, coordiate)
	}

	return nil
}

func (b *Board) GetBoard() *ebiten.Image {
	b.screen.DrawImage(snakeGrid, &ebiten.DrawImageOptions{})

	for _, food := range b.food {
		food.Draw(b.screen)
	}
	b.snake.Draw(b.screen)

	return b.screen
}
