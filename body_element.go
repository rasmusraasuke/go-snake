package main

type BodyElement struct {
	xPos, yPos int
}

func NewBodyElement(x, y int) *BodyElement {
	return &BodyElement{
		xPos: x,
		yPos: y,
	}
}
