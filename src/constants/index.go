package constants

// TileSize represents the size of the side of a square tile
const (
	TileSize         = 32
	DefaultPacmanFPS = 6
)

// GameState represents the game state
type GameState int

// InactiveState - inactive game state
// PlayState - playing game state
const (
	InactiveState GameState = iota
	PlayState
)

// EventType represents a type of event
type EventType int

// FoodEaten - event that indicates a food was eaten
const (
	FoodEaten EventType = iota
)

// Direction expresses a direction
type Direction struct {
	X int
	Y int
}

// DirUp - direction upwards
// DirDown - direction downwards
// DirLeft - direction left
// DirRight - direction right
var (
	DirUp     = Direction{X: 0, Y: -1}
	DirDown   = Direction{X: 0, Y: 1}
	DirLeft   = Direction{X: -1, Y: 0}
	DirRight  = Direction{X: 1, Y: 0}
	DirStatic = Direction{X: 0, Y: 0}
)
