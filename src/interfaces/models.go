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
}

// GameObject interface exposes basic methods for each object inside the maze
type GameObject interface {
	Draw(screen *ebiten.Image, x, y int)
	GetSprite() *ebiten.Image
	GetDirection() constants.Direction
	IsMatrixEditable() bool
	IsUnmovable() bool
}

// MovableGameObject interface special tipe of GameObject
type MovableGameObject interface {
	Draw(screen *ebiten.Image, x, y int)
	GetSprite() *ebiten.Image
	GetDirection() constants.Direction
	IsMatrixEditable() bool
	IsUnmovable() bool
	GetPosition() Location
	SetPosition(x, y int)
}
