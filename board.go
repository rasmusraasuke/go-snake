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

func (b *Board) getGrid() [GRID_SIZE][GRID_SIZE]int {
	grid := new([GRID_SIZE][GRID_SIZE]int)

	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = 0
		}
	}

	for _, bodyElement := range b.snake.body {
		grid[bodyElement.yPos][bodyElement.xPos] = 1
	}

	return *grid
}

func (b *Board) getComputerMove() (int, int) {
	snakeHead := b.snake.body[0]
	headPos := Coordinate{x: snakeHead.xPos, y: snakeHead.yPos}
	var foodPos Coordinate
	for cord, food := range b.food {
		if food.Type == Cherry {
			foodPos = cord
			break
		}
	}

	grid := b.getGrid()

	path := FindPath(grid, headPos, foodPos)
	var nextMove Coordinate
	if len(path) != 0 {
		nextMove = path[1]
	} else {
		nextMove = GetBestNeighbour(&grid, headPos)
	}

	return nextMove.x, nextMove.y
}

func (b *Board) createNewFood(foodType FoodType) {
	openPos := *new([]Coordinate)
	grid := b.getGrid()
	for y := range GRID_SIZE {
		for x := range GRID_SIZE {
			if grid[y][x] != 1 {
				openPos = append(openPos, Coordinate{x: x, y: y})
			}
		}
	}

	randIndex := rand.IntN(len(openPos))
	randCord := openPos[randIndex]
	food := NewFood(foodType, randCord.x, randCord.y)
	b.food[randCord] = *food
}

func (b *Board) Update() error {
	b.snake.UpdateMovementQueue()

	if len(b.food) == 0 {
		b.createNewFood(Cherry)
	}

	if b.wait < WAIT_TIME {
		b.wait++
		return nil
	}
	b.wait = 0

	var newX, newY int
	switch b.snake.playerName {
	case "Computer":
		newX, newY = b.getComputerMove()
	default:
		newX, newY = b.snake.CalculateNextPos()
	}

	if newX < 0 || newY < 0 || newX >= GRID_SIZE || newY >= GRID_SIZE {
		return errors.New(b.snake.playerName + " hit the wall.")
	}

	for i, bodyElement := range b.snake.body {
		if i != len(b.snake.body)-1 && bodyElement.xPos == newX && bodyElement.yPos == newY {
			return errors.New(b.snake.playerName + " hit itself!")
		}
	}

	coordiate := Coordinate{int(newX), int(newY)}

	b.snake.Move(newX, newY)
	switch b.food[coordiate].Type {
	case Cherry:
		b.game.scores[b.snake.playerName] = b.game.scores[b.snake.playerName] + 1
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
