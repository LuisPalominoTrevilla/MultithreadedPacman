package modules

import (
	"math"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// Animator represents an implementation to animate a GameObject
type Animator struct {
	object interfaces.GameObject
}

// DrawFrame of the game object to the specified position in the screen
func (a *Animator) DrawFrame(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	frame := a.object.GetSprite()
	width, height := frame.Size()
	op.GeoM.Scale(constants.TileSize/float64(width), constants.TileSize/float64(height))
	switch a.object.GetDirection() {
	case utils.DirUp:
		op.GeoM.Translate(-constants.TileSize/2, -constants.TileSize/2)
		op.GeoM.Rotate(3 * math.Pi / 2)
		op.GeoM.Translate(constants.TileSize/2, constants.TileSize/2)
	case utils.DirDown:
		op.GeoM.Translate(-constants.TileSize/2, -constants.TileSize/2)
		op.GeoM.Rotate(math.Pi / 2)
		op.GeoM.Translate(constants.TileSize/2, constants.TileSize/2)
	case utils.DirLeft:
		op.GeoM.Translate(-constants.TileSize/2, -constants.TileSize/2)
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(constants.TileSize/2, constants.TileSize/2)
	}
	op.GeoM.Translate(constants.TileSize*float64(x), constants.TileSize*float64(y))
	screen.DrawImage(frame, op)
}

// InitAnimator instantiates the animator linked to a game object
func InitAnimator(object interfaces.GameObject) *Animator {
	animator := Animator{
		object: object,
	}

	return &animator
}
