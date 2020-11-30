package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Pellet represents any pellet in the game
type Pellet struct {
	position   interfaces.Location
	sprite     *ebiten.Image
	isPowerful bool
	animator   *modules.Animator
}

// Draw the element to the screen in given position
func (p *Pellet) Draw(screen *ebiten.Image, x, y int) {
	p.animator.DrawFrame(screen, x, y)
}

// GetSprite of the element
func (p *Pellet) GetSprite() *ebiten.Image {
	return p.sprite
}

// GetDirection of the element
func (p *Pellet) GetDirection() constants.Direction {
	return constants.DirStatic
}

// IsMatrixEditable based on the object direction
func (p *Pellet) IsMatrixEditable() bool {
	return false
}

// CanGhostsGoThrough by any force
func (p *Pellet) CanGhostsGoThrough() bool {
	return true
}

// GetLayerIndex of the element
func (p *Pellet) GetLayerIndex() int {
	return constants.PelletLayerIdx
}

// GetPosition of the element
func (p *Pellet) GetPosition() interfaces.Location {
	return p.position
}

// InitPellet of the maze
func InitPellet(x, y int, isPowerful bool, assetManager *modules.AssetManager) *Pellet {
	pellet := Pellet{
		isPowerful: isPowerful,
		position:   structures.InitPosition(x, y),
	}

	if isPowerful {
		pellet.sprite = assetManager.PowerPelletSprite
	} else {
		pellet.sprite = assetManager.PelletSprite
	}

	pellet.animator = modules.InitAnimator(&pellet)
	return &pellet
}
