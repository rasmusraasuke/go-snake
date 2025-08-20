package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

func GetGrid() *ebiten.Image {
	img := ebiten.NewImage(GRID_SIZE*TILE_SIZE, GRID_SIZE*TILE_SIZE)
	colors := []color.Color{
		colornames.Forestgreen,
		colornames.Green,
	}

	for j := range GRID_SIZE {
		for i := range GRID_SIZE {
			x := float32(i * TILE_SIZE)
			y := float32(j * TILE_SIZE)
			vector.DrawFilledRect(img, x, y, TILE_SIZE, TILE_SIZE, colors[(i+j)%len(colors)], true)
		}
	}
	return img
}
