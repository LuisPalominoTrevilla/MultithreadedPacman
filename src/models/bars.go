package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Bars represents bars that only ghosts can go through
type Bars struct {
	position interfaces.Location
	sprite   *ebiten.Image
	animator *modules.Animator
}

// Draw the element to the screen in given position
func (w *Bars) Draw(screen *ebiten.Image, x, y int) {
	w.animator.DrawFrame(screen, x, y)
}

// GetSprite of the element
func (w *Bars) GetSprite() *ebiten.Image {
	return w.sprite
}

// GetDirection of the element
func (w *Bars) GetDirection() constants.Direction {
	return constants.DirStatic
}

// IsMatrixEditable based on the object direction
func (w *Bars) IsMatrixEditable() bool {
	return false
}

// CanGhostsGoThrough this element
func (w *Bars) CanGhostsGoThrough() bool {
	return true
}

// GetLayerIndex of the element
func (w *Bars) GetLayerIndex() int {
	return constants.BarsLayerIdx
}

// GetPosition of the element
func (w *Bars) GetPosition() interfaces.Location {
	return w.position
}

// InitBars of the maze
func InitBars(x, y int) (*Bars, error) {
	bars := Bars{
		position: structures.InitPosition(x, y),
	}
	img, _, err := ebitenutil.NewImageFromFile("assets/bars.png")
	bars.sprite = img
	bars.animator = modules.InitAnimator(&bars)
	return &bars, err
}
