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
	Body        []BodyElement
	Orientation Direction
}

func New(xPos int, yPos int, orientation Direction) *Snake {
	body := []BodyElement{*NewBodyElement(xPos, yPos)}

	snake := Snake{body, orientation}

	return &snake
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

func (s *Snake) CalculateNextTile() (int, int) {
	headX := s.Body[0].XPos
	headY := s.Body[0].YPos

	switch s.Orientation {
	case Up:
		headY -= 1
	case Right:
		headX += 1
	case Down:
		headY += 1
	case Left:
		headX -= 1
	}

	return headX, headY
}

func (s *Snake) Move(newX, newY int) {
	s.Body = moveLastToFirst(s.Body)
	s.Body[0].XPos = newX
	s.Body[0].YPos = newY
}

func moveLastToFirst(body []BodyElement) []BodyElement {
	if len(body) <= 1 {
		return body
	}

	lastBodyElement := body[len(body)-1]

	copy(body[1:], body[:len(body)-1])

	body[0] = lastBodyElement

	return body
}
