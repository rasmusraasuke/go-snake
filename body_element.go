package main

type BodyElement struct {
	XPos, YPos int
}

func NewBodyElement(x, y int) *BodyElement {
	return &BodyElement{
		XPos: x,
		YPos: y,
	}
}
