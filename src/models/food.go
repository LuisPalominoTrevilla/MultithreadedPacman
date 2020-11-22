package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Food represents a food
type Food struct {
	sprite   *ebiten.Image
	super    bool
	animator *modules.Animator
}

// Draw the element to the screen in given position
func (f *Food) Draw(screen *ebiten.Image, x, y int) {
	f.animator.DrawFrame(screen, x, y)
}

// GetSprite of the element
func (f *Food) GetSprite() *ebiten.Image {
	return f.sprite
}

// GetDirection of the element
func (f *Food) GetDirection() utils.Direction {
	return utils.DirStatic
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
	food.sprite = img
	food.super = isSuper
	food.animator = modules.InitAnimator(&food)
	return &food, err
}