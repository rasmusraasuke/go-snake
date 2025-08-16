package snake

type BodyElement struct {
	XPos int
	YPos int
}

func NewBodyElement(x, y int) *BodyElement {
	return &BodyElement{
		XPos: x,
		YPos: y,
	}
}
