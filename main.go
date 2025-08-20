package main

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

var snakeGrid *ebiten.Image

func init() {
	snakeGrid = GetGrid()
}

type Main struct {
	ui      *ebitenui.UI
	players []*Player
	games   []*SnakeGame
}

func NewMain() *Main {
	main := Main{
		players: *new([]*Player),
		games:   *new([]*SnakeGame),
		ui:      &ebitenui.UI{},
	}

	main.ui.Container = CreateMenu(&main)
	return &main

}

func (m *Main) StartSinglePlayer() {
	player := NewPlayer(WASD)
	game := NewSnakeGame(*player)

	m.players = append(m.players, player)
	m.games = append(m.games, game)
}

func (m *Main) Update() error {
	m.ui.Update()
	for _, game := range m.games {
		error := game.Update()

		if error != nil {
			return error
		}
	}
	return nil
}

func (m *Main) Draw(screen *ebiten.Image) {
	m.ui.Draw(screen)

	for _, game := range m.games {
		game.Draw(screen)
	}
}

func (m *Main) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("GoSnake")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewMain()); err != nil {
		log.Fatal(err)
	}
}
