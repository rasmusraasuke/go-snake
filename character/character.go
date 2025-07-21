package character

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MovementSpeed = 4
)

type Character struct {
	XPos float64
	YPos float64
}

func (c *Character) MoveCharacter() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.XPos -= MovementSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.XPos += MovementSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.YPos -= MovementSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.YPos += MovementSpeed
	}
}
