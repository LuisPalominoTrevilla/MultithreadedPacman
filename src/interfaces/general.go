package interfaces

import (
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

// ChaseBehavior of a ghost
type ChaseBehavior interface {
	SwitchDirection()
}

// Screen represents any game screen
type Screen interface {
	Run()
	Draw(screen *ebiten.Image)
}
