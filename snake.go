package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Snake struct {
	playerName    string
	body          []*BodyElement
	orientation   Direction
	movementQueue []Direction
	keys          map[Direction]ebiten.Key
}

func NewSnake(playerName string, pos Coordinate, orient Direction, input InputType) *Snake {
	head := NewBodyElement(pos.x, pos.y)
	body := []*BodyElement{head}
	queue := new([]Direction)
	keys := setKeys(input)

	snake := Snake{playerName, body, orient, *queue, keys}
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

func (s *Snake) UpdateMovementQueue() {
	if inpututil.IsKeyJustPressed(s.keys[Up]) {
		s.tryAddDirection(Up)
	} else if inpututil.IsKeyJustPressed(s.keys[Down]) {
		s.tryAddDirection(Down)
	} else if inpututil.IsKeyJustPressed(s.keys[Left]) {
		s.tryAddDirection(Left)
	} else if inpututil.IsKeyJustPressed(s.keys[Right]) {
		s.tryAddDirection(Right)
	}
}

func (s *Snake) tryAddDirection(newDirection Direction) {
	oppositeDirection := (newDirection + 2) % 4
	if s.orientation == oppositeDirection && len(s.movementQueue) == 0 {
		return
	}

	if len(s.movementQueue) > 0 {
		lastQueued := s.movementQueue[len(s.movementQueue)-1]
		if lastQueued == newDirection || lastQueued == oppositeDirection {
			return
		}
	}

	s.movementQueue = append(s.movementQueue, newDirection)
}

func (s *Snake) CalculateNextPos() (int, int) {
	headX := s.body[0].xPos
	headY := s.body[0].yPos

	var newOrientation Direction
	if len(s.movementQueue) != 0 {
		newOrientation = s.movementQueue[0]
		s.movementQueue = s.movementQueue[1:]
	} else {
		newOrientation = s.orientation
	}

	switch newOrientation {
	case Up:
		headY -= 1
	case Right:
		headX += 1
	case Down:
		headY += 1
	case Left:
		headX -= 1
	}
	s.orientation = newOrientation

	return headX, headY
}

func (s *Snake) EatCherry() {
	lastBodyElement := s.body[len(s.body)-1]
	newBodyElement := NewBodyElement(lastBodyElement.xPos, lastBodyElement.yPos)
	s.body = append(s.body, newBodyElement)
}

func (s *Snake) Move(newX, newY int) {
	s.body = moveLastToFirst(s.body)
	s.body[0].xPos = newX
	s.body[0].yPos = newY
}

func moveLastToFirst(body []*BodyElement) []*BodyElement {
	if len(body) <= 1 {
		return body
	}

	lastBodyElement := body[len(body)-1]

	copy(body[1:], body[:len(body)-1])

	body[0] = lastBodyElement

	return body
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, bodyElement := range s.body {
		x := float32(TILE_SIZE * bodyElement.xPos)
		y := float32(TILE_SIZE * bodyElement.yPos)
		vector.DrawFilledRect(screen, x, y, TILE_SIZE, TILE_SIZE, colornames.Peru, true)
	}
}
