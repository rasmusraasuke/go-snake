package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Direction int

const (
	Up    Direction = iota // EnumIndex = 0
	Right                  // EnumIndex = 1
	Down                   // EnumIndex = 2
	Left                   // EnumIndex = 3
)

type Snake struct {
	XPos        int
	YPos        int
	Orientation Direction
}

func (s *Snake) UpdateOrientation() {
	if ebiten.IsKeyPressed(ebiten.KeyW) && s.Orientation != Down {
		s.Orientation = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyS) && s.Orientation != Up {
		s.Orientation = Down
	} else if ebiten.IsKeyPressed(ebiten.KeyA) && s.Orientation != Right {
		s.Orientation = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyD) && s.Orientation != Left {
		s.Orientation = Right
	}
}

func (s *Snake) Move() {
	switch s.Orientation {
	case Up:
		s.YPos -= 1
	case Right:
		s.XPos += 1
	case Down:
		s.YPos += 1
	case Left:
		s.XPos -= 1
	}
}
