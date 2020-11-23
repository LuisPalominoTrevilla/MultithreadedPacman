package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Wall represents a wall
type Wall struct {
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

// InitWall of the maze
func InitWall() (*Wall, error) {
	wall := Wall{}
	img, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
	wall.sprite = img
	wall.animator = modules.InitAnimator(&wall)
	return &wall, err
}
