package structures

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

// InitPosition of a game object
func InitPosition(x, y int) *Position {
	return &Position{x: x, y: y}
}

// TODO: Method to calculate distance between positions
