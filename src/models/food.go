package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
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
func (f *Food) GetDirection() constants.Direction {
	return constants.DirStatic
}

// IsMatrixEditable based on the object direction
func (f *Food) IsMatrixEditable() bool {
	return false
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
	food.animator = modules.InitAnimator(&food)
	return &food, err
}
