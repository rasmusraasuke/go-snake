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
var cherry *ebiten.Image
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

	cherry, _, err = ebitenutil.NewImageFromFile("assets/cherry.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	snake snake.Snake
	food  map[Coordinate]FoodType
}

func (g *Game) Update() error {
	g.snake.UpdatePendingOrientation()

	if len(g.food) == 0 {
		randX := rand.IntN(GRID_SIZE)
		randY := rand.IntN(GRID_SIZE)
		g.food[Coordinate{randX, randY}] = Cherry
		log.Printf("Cherry now at [%d; %d]\n", randX, randY)
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

	coordiate := Coordinate{newX, newY}
	switch g.food[coordiate] {
	case 0:
		g.snake.Move(newX, newY)
	case Cherry:
		g.snake.EatCherry(newX, newY)
		delete(g.food, coordiate)
	}

	return nil
}

func DrawBoard(boardImage *ebiten.Image) {

}

func (g *Game) Draw(screen *ebiten.Image) {
	backgroundSize := background.Bounds().Size()
	cherrySize := cherry.Bounds().Size()

	tileX := screen.Bounds().Dx() / GRID_SIZE
	tileY := screen.Bounds().Dy() / GRID_SIZE

	tileXScale := float64(tileX) / float64(backgroundSize.X)
	tileYScale := float64(tileY) / float64(backgroundSize.Y)
	cherryXScale := float64(tileX) / float64(cherrySize.X)
	cherryYScale := float64(tileY) / float64(cherrySize.Y)

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

	for cords, _ := range g.food {
		op := &ebiten.DrawImageOptions{}
		x := tileX * cords.X
		y := tileY * cords.Y
		op.GeoM.Scale(cherryXScale, cherryYScale)
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(cherry, op)

	}

	for _, bodyElement := range g.snake.Body {
		op := &ebiten.DrawImageOptions{}
		x := tileX * bodyElement.XPos
		y := tileY * bodyElement.YPos
		op.GeoM.Scale(tileXScale, tileYScale)
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(square, op)
	}
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
