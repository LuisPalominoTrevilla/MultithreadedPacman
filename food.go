package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Food represents a food
type Food struct {
	sprite *ebiten.Image
	super  bool
}

// InitFood of the maze
func InitFood(isSuper bool) (*Food, error) {
	food := Food{}
	var assetFile string
	if isSuper {
		assetFile = "assets/superFood.png"
	} else {
		assetFile = "assets/food.png"
	}

	img, _, err := ebitenutil.NewImageFromFile(assetFile)
	food.sprite = img
	food.super = isSuper
	return &food, err
}
