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
	players map[string]*Player
	game    *Game
	ui      *ebitenui.UI
}

func NewMain() *Main {
	players := make(map[string]*Player)
	players["Computer"] = NewComputer()

	main := Main{
		state:   MAIN_MENU,
		players: players,
		game:    &Game{},
		ui:      &ebitenui.UI{},
	}

	main.ui.Container = CreateMenu(&main)
	return &main
}

func (m *Main) FindOrCreatePlayer(name string, input InputType) *Player {
	if m.players[name] != nil {
		return m.players[name]
	}

	player := NewPlayer(name, input)
	m.players[name] = player
	return player

}

func (m *Main) StartSinglePlayer(name string, input InputType) {
	players := *new([2]*Player)

	players[0] = m.FindOrCreatePlayer(name, input)
	players[1] = m.players["Computer"]

	m.game = NewGame(players)
	m.state = GAME
}

func (m *Main) StartTwoPlayers(name1, name2 string, input1, input2 InputType) {
	players := *new([2]*Player)

	players[0] = m.FindOrCreatePlayer(name1, input1)
	players[1] = m.FindOrCreatePlayer(name2, input2)

	m.game = NewGame(players)
	m.state = GAME
}

func (m *Main) Update() error {
	switch m.state {
	case MAIN_MENU, GAME_OVER:
		m.ui.Update()
	case GAME:
		error := m.game.Update()
		if error != nil {
			m.state = GAME_OVER
			m.ui.Container = CreateGameOver(m, error.Error(), m.game.scores)
		}
	}
	return nil
}

func (m *Main) Draw(screen *ebiten.Image) {
	switch m.state {
	case MAIN_MENU, GAME_OVER:
		m.ui.Draw(screen)
	case GAME:
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
