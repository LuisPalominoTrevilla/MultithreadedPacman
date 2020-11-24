package constants

// TileSize represents the size of the side of a square tile
const (
	TileSize         = 32
	DefaultPacmanFPS = 6
	InitialGhostFPS  = 6
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

// SoundEffect represents a type of sound effect
type SoundEffect int

// MunchEffect - pacman munch sound FX
// GameStart - pacman's main game start theme
const (
	MunchEffect SoundEffect = iota
	GameStart
)

// AudioFiles for each sound effect
var AudioFiles = map[SoundEffect][]string{
	MunchEffect: {"assets/audio/munch_1.wav", "assets/audio/munch_2.wav"},
	GameStart:   {"assets/audio/game_start.wav"},
}

// GhostType represents a type of ghost
type GhostType string

// RedGhost - The classic pacman red ghost
// CyanGhost - The classic pacman cyan ghost
// PinkGhost - The classic pacman pink ghost
// OrangeGhost - The classic pacman orange ghost
const (
	RedGhost    GhostType = "red"
	CyanGhost   GhostType = "cyan"
	PinkGhost   GhostType = "pink"
	OrangeGhost GhostType = "orange"
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
