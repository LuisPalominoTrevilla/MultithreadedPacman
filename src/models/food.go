package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Food represents a food
type Food struct {
	Sprite *ebiten.Image
	super  bool
}

// InitFood of the maze
func InitFood(isSuper bool) (*Food, error) {
	food := Food{}
	var assetFile string
	if isSuper {
		assetFile = "src/assets/superFood.png"
	} else {
		assetFile = "src/assets/food.png"
	}

	img, _, err := ebitenutil.NewImageFromFile(assetFile)
	food.Sprite = img
	food.super = isSuper
	return &food, err
}
