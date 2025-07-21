package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rasmusraasuke/gogame/character"
	_ "image/jpeg"
	_ "image/png"

	"log"
)

var snake *ebiten.Image
var background *ebiten.Image

func init() {
	var err error
	snake, _, err = ebitenutil.NewImageFromFile("assets/snake.png")
	if err != nil {
		log.Fatal(err)
	}
	background, _, err = ebitenutil.NewImageFromFile("assets/jungle.jpg")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	character character.Character
}

func (g *Game) Update() error {
	g.character.MoveCharacter()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	screen.DrawImage(background, op)

	op.GeoM.Translate(g.character.YPos, g.character.XPos)
	op.GeoM.Scale(0.4, 0.3)
	screen.DrawImage(snake, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
