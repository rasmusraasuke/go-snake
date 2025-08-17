package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type SnakeGrid struct {
	asset *ebiten.Image
}

func NewSnakeGrid() *SnakeGrid {
	image, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	snakegrid := SnakeGrid{asset: image}
	return &snakegrid
}

func (g *SnakeGrid) Draw(screen *ebiten.Image) {
	for j := range GRID_SIZE {
		for i := range GRID_SIZE {
			op := &ebiten.DrawImageOptions{}
			x := i * TILE_SIZE
			y := j * TILE_SIZE
			op.GeoM.Scale(tileXScale, tileYScale)
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.asset, op)
		}
	}
}

