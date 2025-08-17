package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Snake struct {
	Body               []BodyElement
	orientation        Direction
	pendingOrientation Direction
	keys               map[Direction]ebiten.Key

	asset *ebiten.Image
}

func NewSnake(xPos, yPos int, orient Direction, input InputType) *Snake {
	image, _, err := ebitenutil.NewImageFromFile("assets/square.png")
	if err != nil {
		log.Fatal(err)
	}

	body := []BodyElement{*NewBodyElement(xPos, yPos)}

	keys := setKeys(input)

	snake := Snake{body, orient, orient, keys, image}

	return &snake
}

func setKeys(input InputType) map[Direction]ebiten.Key {
	keys := make(map[Direction]ebiten.Key)

	switch input {
	case WASD:
		keys[Up] = ebiten.KeyW
		keys[Right] = ebiten.KeyD
		keys[Down] = ebiten.KeyS
		keys[Left] = ebiten.KeyA
	case ARROWS:
		keys[Up] = ebiten.KeyUp
		keys[Right] = ebiten.KeyRight
		keys[Down] = ebiten.KeyDown
		keys[Left] = ebiten.KeyLeft
	case VIM:
		keys[Up] = ebiten.KeyK
		keys[Right] = ebiten.KeyL
		keys[Down] = ebiten.KeyJ
		keys[Left] = ebiten.KeyH
	}

	return keys
}

func (s *Snake) UpdatePendingOrientation() {
	if ebiten.IsKeyPressed(s.keys[Up]) && s.orientation != Down {
		s.pendingOrientation = Up
	} else if ebiten.IsKeyPressed(s.keys[Down]) && s.orientation != Up {
		s.pendingOrientation = Down
	} else if ebiten.IsKeyPressed(s.keys[Left]) && s.orientation != Right {
		s.pendingOrientation = Left
	} else if ebiten.IsKeyPressed(s.keys[Right]) && s.orientation != Left {
		s.pendingOrientation = Right
	}
}

func (s *Snake) CalculateNextPos() (int, int) {
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
