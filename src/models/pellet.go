package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Pellet represents any pellet in the game
type Pellet struct {
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

// IsUnmovable by any force
func (p *Pellet) IsUnmovable() bool {
	return false
}

// InitPellet of the maze
func InitPellet(isPowerful bool) (*Pellet, error) {
	pellet := Pellet{}
	var assetFile string
	if isPowerful {
		assetFile = "assets/super-pellet.png"
	} else {
		assetFile = "assets/pellet.png"
	}

	img, _, err := ebitenutil.NewImageFromFile(assetFile)
	pellet.sprite = img
	pellet.isPowerful = isPowerful
	pellet.animator = modules.InitAnimator(&pellet)
	return &pellet, err
}
