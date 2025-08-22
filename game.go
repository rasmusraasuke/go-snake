package main

import (
	"errors"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Coordinate struct {
	x, y int
}

type Game struct {
	snakes []*Snake
	boards []*Board
}

func NewGame(players [2]*Player) *Game {
	game := &Game{}
	snakes := *new([]*Snake)
	boards := *new([]*Board)

	for _, player := range players {
		coordinate := Coordinate{int(GRID_SIZE / 2), int(GRID_SIZE / 2)}
		startOrient := rand.IntN(4)
		snake := NewSnake(player.name, coordinate, Direction(startOrient), player.input)

		board := NewBoard(game, snake)

		snakes = append(snakes, snake)
		boards = append(boards, board)
	}

	game.boards = boards
	game.snakes = snakes
	return game
}

func (g *Game) FeedOtherSnake(ownName string) {
	for _, snake := range g.snakes {
		if snake.playerName == ownName {
			continue
		}

		snake.EatCherry()
	}
}

func (g *Game) Update() error {
	gameErrors := *new([]error)
	for _, board := range g.boards {
		error := board.Update()

		if error != nil {
			gameErrors = append(gameErrors, error)
		}
	}

	if len(gameErrors) == 2 {
		return errors.New("Draw")
	} else if len(gameErrors) == 1 {
		return gameErrors[0]
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	size := screen.Bounds().Size()
	marginX := float64(size.X-GRID_SIZE*TILE_SIZE*BOARD_COUNT) / float64(BOARD_COUNT+1)
	marginY := float64(size.Y-GRID_SIZE*TILE_SIZE) / 2

	vector.DrawFilledRect(screen, 0, 0, float32(size.X), float32(size.Y), colornames.Darkolivegreen, true)
	for i, board := range g.boards {
		board := board.GetBoard()

		x := float64(i)*float64(board.Bounds().Dx()) + float64(i+1)*marginX
		y := marginY

		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)

		screen.DrawImage(board, &op)
	}
}
