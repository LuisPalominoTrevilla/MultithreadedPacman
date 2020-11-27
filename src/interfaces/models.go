package interfaces

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

// GameObject interface exposes basic methods for each object inside the maze
type GameObject interface {
	Draw(screen *ebiten.Image, x, y int)
	GetSprite() *ebiten.Image
	GetDirection() constants.Direction
	IsMatrixEditable() bool
	CanGhostsGoThrough() bool
	GetLayerIndex() int
	GetPosition() Location
}

// MovableGameObject interface special tipe of GameObject
type MovableGameObject interface {
	Draw(screen *ebiten.Image, x, y int)
	GetSprite() *ebiten.Image
	GetDirection() constants.Direction
	IsMatrixEditable() bool
	CanGhostsGoThrough() bool
	GetLayerIndex() int
	GetPosition() Location
	SetPosition(x, y int)
}
