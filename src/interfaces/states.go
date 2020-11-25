package interfaces

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

// GhostState represents a ghost state
type GhostState interface {
	ApplyTransition(event constants.StateEvent) GhostState
	Run()
	GetSprite() *ebiten.Image
}
