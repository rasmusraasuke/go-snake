package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Food struct {
	Type  FoodType
	xPos  int
	yPos  int
	asset *ebiten.Image
}

func NewFood(foodType FoodType, xPos, yPos int) *Food {
	asset := determineAsset(foodType)
	image, _, err := ebitenutil.NewImageFromFile(asset)
	if err != nil {
		log.Fatal(err)
	}

	food := Food{Type: foodType, xPos: xPos, yPos: yPos, asset: image}
	return &food
}

func determineAsset(foodType FoodType) string {
	var asset string

	switch foodType {
	case Cherry:
		asset = "assets/cherry.png"
	}
	return asset
}

func (f *Food) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	x := TILE_SIZE * f.xPos
	y := TILE_SIZE * f.yPos
	op.GeoM.Scale(cherryXScale, cherryYScale)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(f.asset, op)
}
