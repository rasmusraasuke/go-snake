package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rasmusraasuke/gogame/snake"

	"log"
)

type GridValue int

const (
	Empty GridValue = iota
	Snake
	Wall
	Food
)

const GRID_SIZE = 102

var square *ebiten.Image
var background *ebiten.Image
var grid = make([][]GridValue, GRID_SIZE)
var wait = 0

func init() {
	var err error

	square, _, err = ebitenutil.NewImageFromFile("assets/square.png")
	if err != nil {
		log.Fatal(err)
	}

	background, _, err = ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	for i := range grid {
		grid[i] = make([]GridValue, GRID_SIZE)
		for j := range grid[i] {
			if i == 0 || j == 0 || i == GRID_SIZE-1 || j == GRID_SIZE-1 {
				grid[i][j] = Wall
			} else {
				grid[i][j] = Empty
			}
		}
	}
	grid[1][1] = Snake
}

type Game struct {
	snake snake.Snake
}

func (g *Game) Update() error {
	g.snake.UpdateOrientation()

	//if wait < 60 {
	//	wait++
	//	return nil
	//}
	//wait = 0

	grid[g.snake.YPos][g.snake.XPos] = Empty
	g.snake.Move()
	grid[g.snake.YPos][g.snake.XPos] = Snake

	log.Println("Snake orientation:", g.snake.Orientation)
	log.Println("Snake position:", g.snake.XPos, g.snake.YPos)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	newXCord := float64((screen.Bounds().Dx() / GRID_SIZE) * g.snake.XPos)
	newYCord := float64((screen.Bounds().Dy() / GRID_SIZE) * g.snake.YPos)
	//log.Println("Drawing snake at", newXCord, newYCord)

	op.GeoM.Scale(0.2, 0.2)
	op.GeoM.Translate(newXCord, newYCord)

	screen.DrawImage(square, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{snake: snake.Snake{XPos: 1, YPos: 1, Orientation: snake.Right}}); err != nil {
		log.Fatal(err)
	}
}
