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

func (m *Main) StartTwoPlayers() {
	player1 := NewPlayer(WASD)
	player2 := NewPlayer(ARROWS)
	game1 := NewSnakeGame(*player1)
	game2 := NewSnakeGame(*player2)

	m.players = append(m.players, player1, player2)
	m.games = append(m.games, game1, game2)
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

	size := screen.Bounds().Size()
	marginX := float64(size.X-GRID_SIZE*TILE_SIZE*len(m.games)) / float64(len(m.games)+1)
	marginY := float64(size.Y-GRID_SIZE*TILE_SIZE) / 2

	for i, game := range m.games {
		img := game.GetImage()
		x := float64(i)*float64(img.Bounds().Dx()) + float64(i+1)*marginX
		y := marginY

		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)

		screen.DrawImage(img, &op)
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
