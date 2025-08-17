package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
}

type Main struct {
	player Player
}

func (m *Main) Update() error {
	error := m.player.game.Update()
	return error
}

func (m *Main) Draw(screen *ebiten.Image) {
	m.player.game.Draw(screen)
}

func (m *Main) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Go Snake")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	player := NewPlayer(ARROWS)

	if err := ebiten.RunGame(&Main{player: *player}); err != nil {
		log.Fatal(err)
	}
}
