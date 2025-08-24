package main

type BodyElement struct {
	xPos, yPos       int
	lastX, lastY     int
	visualX, visualY float64
}

func NewBodyElement(x, y int) *BodyElement {
	return &BodyElement{
		xPos:    x,
		yPos:    y,
		lastX:   x,
		lastY:   y,
		visualX: float64(x),
		visualY: float64(y),
	}
}
