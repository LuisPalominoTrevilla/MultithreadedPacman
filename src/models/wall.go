package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Wall represents a wall
type Wall struct {
	position interfaces.Location
	sprite   *ebiten.Image
	animator *modules.Animator
}

// Draw the element to the screen in given position
func (w *Wall) Draw(screen *ebiten.Image, x, y int) {
	w.animator.DrawFrame(screen, x, y)
}

// GetSprite of the element
func (w *Wall) GetSprite() *ebiten.Image {
	return w.sprite
}

// GetDirection of the element
func (w *Wall) GetDirection() constants.Direction {
	return constants.DirStatic
}

// IsMatrixEditable based on the object direction
func (w *Wall) IsMatrixEditable() bool {
	return false
}

// CanGhostsGoThrough by any force
func (w *Wall) CanGhostsGoThrough() bool {
	return false
}

// GetLayerIndex of the element
func (w *Wall) GetLayerIndex() int {
	return constants.WallLayerIdx
}

// GetPosition of the element
func (w *Wall) GetPosition() interfaces.Location {
	return w.position
}

// InitWall of the maze
func InitWall(x, y int, assetManager *modules.AssetManager) *Wall {
	wall := Wall{
		position: structures.InitPosition(x, y),
	}
	wall.sprite = assetManager.WallSprite
	wall.animator = modules.InitAnimator(&wall)
	return &wall
}
