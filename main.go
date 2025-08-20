package main

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	_ "image/png"
	"log"
)

type Main struct {
	ui      *ebitenui.UI
	players []Player
	games   []SnakeGame
}

func NewMain() *Main {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Steelblue),
		),
	)
	return &Main{
		players: *new([]Player),
		games:   *new([]SnakeGame),
		ui:      &ebitenui.UI{Container: root},
	}

}

func (m *Main) Update() error {
	m.ui.Update()
	//for _, game := range m.games {
	//	error := game.Update()

	//	if error != nil {
	//		return error
	//	}
	//}
	return nil
}

func (m *Main) Draw(screen *ebiten.Image) {
	m.ui.Draw(screen)
	//for _, game := range m.games {
	//	game.Draw(screen)
	//}
}

func (m *Main) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Go Snake")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewMain()); err != nil {
		log.Fatal(err)
	}
}
