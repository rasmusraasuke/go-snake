package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Snake struct {
	Body               []BodyElement
	orientation        Direction
	pendingOrientation Direction

	asset *ebiten.Image
}

func NewSnake(xPos int, yPos int, orient Direction) *Snake {
	image, _, err := ebitenutil.NewImageFromFile("assets/square.png")
	if err != nil {
		log.Fatal(err)
	}

	body := []BodyElement{*NewBodyElement(xPos, yPos)}

	snake := Snake{body, orient, orient, image}

	return &snake
}

func (s *Snake) UpdatePendingOrientation() {
	if ebiten.IsKeyPressed(ebiten.KeyW) && s.orientation != Down {
		s.pendingOrientation = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyS) && s.orientation != Up {
		s.pendingOrientation = Down
	} else if ebiten.IsKeyPressed(ebiten.KeyA) && s.orientation != Right {
		s.pendingOrientation = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyD) && s.orientation != Left {
		s.pendingOrientation = Right
	}
}

func (s *Snake) CalculateNextTile() (int, int) {
	headX := s.Body[0].XPos
	headY := s.Body[0].YPos

	switch s.pendingOrientation {
	case Up:
		headY -= 1
	case Right:
		headX += 1
	case Down:
		headY += 1
	case Left:
		headX -= 1
	}
	s.orientation = s.pendingOrientation

	return headX, headY
}

func (s *Snake) EatCherry(newX, newY int) {
	newBodyElement := NewBodyElement(newX, newY)
	s.Body = append([]BodyElement{*newBodyElement}, s.Body...)
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

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, bodyElement := range s.Body {
		op := &ebiten.DrawImageOptions{}
		x := TILE_SIZE * bodyElement.XPos
		y := TILE_SIZE * bodyElement.YPos
		op.GeoM.Scale(tileXScale, tileYScale)
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(s.asset, op)
	}
}
