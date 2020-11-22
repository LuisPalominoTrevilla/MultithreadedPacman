package interfaces

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// GameObject interface exposes basic methods for each object inside the maze
type GameObject interface {
	Draw(screen *ebiten.Image, x, y int)
	GetSprite() *ebiten.Image
	GetDirection() utils.Direction
}
