package structures

import (
	"math"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
)

// Position of a game object
type Position struct {
	x int
	y int
}

// X getter
func (p Position) X() int {
	return p.x
}

// Y getter
func (p Position) Y() int {
	return p.y
}

// SetX coordinate
func (p *Position) SetX(x int) {
	p.x = x
}

// SetY coordinate
func (p *Position) SetY(y int) {
	p.y = y
}

// DistanceTo a given position
func (p Position) DistanceTo(o interfaces.Location) float64 {
	return math.Sqrt(math.Pow(float64(o.X()-p.x), 2) + math.Pow(float64(o.Y()-p.y), 2))
}

// InitPosition of a game object
func InitPosition(x, y int) *Position {
	return &Position{x: x, y: y}
}
