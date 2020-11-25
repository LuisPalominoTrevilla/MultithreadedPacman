package constants

// TileSize represents the size of the side of a square tile
const (
	TileSize            = 32
	MaxGhostsAllowed    = 8
	DefaultPacmanFPS    = 6
	DefaultGhostFPS     = 6
	TimeBetweenSpawns   = 3
	ScatterModeDuration = 7
	ChaseModeDuration   = 20
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

// EatFood - Indicates a food was eaten
// Scatter - Whenever a ghost starts scattering
// ChasePacman - Whenever a ghost starts chasing pacman
const (
	EatFood EventType = iota
	Scatter
	ChasePacman
)

// SoundEffect represents a type of sound effect
type SoundEffect int

// MunchEffect - Pacman munch sound FX
// GameStart - Pacman's main game start theme
// GhostSiren - Ghost siren sound
const (
	MunchEffect SoundEffect = iota
	GameStart
	GhostSiren
)

// AudioFiles for each sound effect
var AudioFiles = map[SoundEffect][]string{
	MunchEffect: {"assets/audio/munch_1.wav", "assets/audio/munch_2.wav"},
	GameStart:   {"assets/audio/game_start.wav"},
	GhostSiren:  {"assets/audio/siren_1.wav"},
}

// GhostType represents a type of ghost
type GhostType string

// Blinky - The classic pacman red ghost
// Inky - The classic pacman cyan ghost
// Pinky - The classic pacman pink ghost
// Clyde - The classic pacman orange ghost
const (
	Blinky GhostType = "red"
	Inky   GhostType = "cyan"
	Pinky  GhostType = "pink"
	Clyde  GhostType = "orange"
)

// GhostState represents a ghost state
type GhostState int

// IdleState - Initial ghost value
// ScatterState - Normal ghost behavior
// ChaseState - Behavior to chase pacman
const (
	IdleState GhostState = iota
	ScatterState
	ChaseState
)

// Direction expresses a direction
type Direction struct {
	X int
	Y int
}

// IsOpposite to a given direction
func (d Direction) IsOpposite(other Direction) bool {
	return d.X*-1 == other.X && d.Y*-1 == other.Y
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

// PossibleDirections to move
var PossibleDirections = []Direction{
	DirUp,
	DirDown,
	DirLeft,
	DirRight,
}
