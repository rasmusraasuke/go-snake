package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
}

type Main struct {
	game SnakeGame
}

func (m *Main) Update() error {
	error := m.game.Update()
	return error
}

func (m *Main) Draw(screen *ebiten.Image) {
	m.game.Draw(screen)
}

func (m *Main) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Go Snake")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := NewSnakeGame()

	if err := ebiten.RunGame(&Main{*game}); err != nil {
		log.Fatal(err)
	}
}
