package main

import (
	"errors"
	_ "image/png"

	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rasmusraasuke/gogame/snake"

	"log"
)

const GRID_SIZE = 20

type FoodType int

const (
	Cherry FoodType = iota + 1
)

type Coordinate struct {
	X, Y int
}

var square *ebiten.Image
var background *ebiten.Image
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
}

type Game struct {
	snake snake.Snake
	food  map[Coordinate]FoodType
}

func (g *Game) Update() error {
	g.snake.UpdateOrientation()

	if len(g.food) == 0 {
		g.food[Coordinate{3, 3}] = Cherry
	}

	if wait < 10 {
		wait++
		return nil
	}
	wait = 0

	newX, newY := g.snake.CalculateNextTile()
	if newX < 0 || newY < 0 || newX >= GRID_SIZE || newY >= GRID_SIZE {
		return errors.New("Snake hit it's head against the wall!")
	}

	g.snake.Move(newX, newY)

	return nil
}

func DrawBoard(boardImage *ebiten.Image) {

}

func (g *Game) Draw(screen *ebiten.Image) {
	imageSize := background.Bounds().Size()

	tileX := screen.Bounds().Dx() / GRID_SIZE
	tileY := screen.Bounds().Dy() / GRID_SIZE

	tileXScale := float64(tileX) / float64(imageSize.X)
	tileYScale := float64(tileY) / float64(imageSize.Y)

	for j := range GRID_SIZE {
		for i := range GRID_SIZE {
			op := &ebiten.DrawImageOptions{}
			x := i * tileX
			y := j * tileY
			op.GeoM.Scale(tileXScale, tileYScale)
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(background, op)
		}
	}

	newXCord := float64((tileX) * g.snake.Body[0].XPos)
	newYCord := float64((tileY) * g.snake.Body[0].YPos)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(tileXScale, tileYScale)
	op.GeoM.Translate(newXCord, newYCord)

	screen.DrawImage(square, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Go Snake")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	startX := int(GRID_SIZE / 2)
	startY := int(GRID_SIZE / 2)
	startOrient := rand.IntN(4)

	snake := snake.New(startX, startY, snake.Direction(startOrient))
	food := make(map[Coordinate]FoodType)

	if err := ebiten.RunGame(&Game{*snake, food}); err != nil {
		log.Fatal(err)
	}
}
