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
	state   State
	players []*Player
	game    *Game
	ui      *ebitenui.UI
}

func NewMain() *Main {
	main := Main{
		state:   MAIN_MENU,
		players: *new([]*Player),
		game:    &Game{},
		ui:      &ebitenui.UI{},
	}

	main.ui.Container = CreateMenu(&main)
	return &main

}

func (m *Main) StartSinglePlayer(name string, input InputType) {
	players := make(map[string]InputType)
	players[name] = input
	players["Computer"] = COMPUTER

	m.game = NewGame(players)
	m.state = GAME
}

func (m *Main) StartTwoPlayers(name1, name2 string, input1, input2 InputType) {
	players := make(map[string]InputType)
	players[name1] = input1
	players[name2] = input2

	m.game = NewGame(players)
	m.state = GAME
}

func (m *Main) Update() error {
	if m.state == MAIN_MENU {
		m.ui.Update()
	} else if m.state == GAME {
		m.game.Update()
	}

	return nil
}

func (m *Main) Draw(screen *ebiten.Image) {
	if m.state == MAIN_MENU {
		m.ui.Draw(screen)
	} else if m.state == GAME {
		m.game.Draw(screen)
	}
}

func (m *Main) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	x, y := ebiten.Monitor().Size()
	ebiten.SetWindowSize(x, y)
	ebiten.SetWindowTitle("GoSnake")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	if err := ebiten.RunGame(NewMain()); err != nil {
		log.Fatal(err)
	}
}
