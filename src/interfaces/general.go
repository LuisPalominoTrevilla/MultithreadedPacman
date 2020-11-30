package interfaces

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

// Location interface exposes basic methods exclusive to a location
type Location interface {
	X() int
	Y() int
	SetX(x int)
	SetY(y int)
	DistanceTo(Location) float64
}

// Screen represents any game screen
type Screen interface {
	Run(nextScreen chan constants.GameState)
	Draw(screen *ebiten.Image)
}
