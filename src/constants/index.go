package constants

// Standard constants used in the codebase
const (
	TileSize           = 32
	MaxGhostsAllowed   = 8
	InfiniteChasePhase = 3
	TimeBetweenSpawns  = 3
)

// Speed constants
const (
	DefaultPacmanFPS = 6
	PowerPacmanFPS   = 8
	DefaultGhostFPS  = 6
	FleeingGhostFPS  = 4
	EatenGhostFPS    = 12
)

// Fixed duration of phases
const (
	ScatterModeDuration     = 7
	ChaseModeDuration       = 20
	FlickeringStateDuration = 2
	PowerPelletDuration     = 7
)

// Default layer indexes for objects
const (
	WallLayerIdx         = 6
	BarsLayerIdx         = 5
	GhostLayerIdx        = 4
	PacmanLayerIdx       = 3
	FleeingGhostLayerIdx = 2
	PelletLayerIdx       = 1
)

// GameState represents the game state
type GameState int

// InactiveState - inactive game state
// PlayState - playing game state
const (
	InactiveState GameState = iota
	PlayState
)

// SoundEffect represents a type of sound effect
type SoundEffect int

// MunchEffect - Pacman munch sound FX
// GameStart - Pacman's main game start theme
// GhostSirenPhaseX - Ghost siren sounds
// PowerPellet - Power pellet background music
// EatGhostEffect - Eating ghost sound effect
// Retreating - Ghost reatreating audio
const (
	MunchEffect SoundEffect = iota
	GameStart
	GhostSirenPhase1
	GhostSirenPhase2
	GhostSirenPhase3
	GhostSirenPhase4
	PowerPellet
	EatGhostEffect
	Retreating
)

// AudioFiles for each sound effect
var AudioFiles = map[SoundEffect][]string{
	MunchEffect:      {"assets/audio/munch_1.wav", "assets/audio/munch_2.wav"},
	GameStart:        {"assets/audio/game_start.wav"},
	GhostSirenPhase1: {"assets/audio/siren_1.wav"},
	GhostSirenPhase2: {"assets/audio/siren_2.wav"},
	GhostSirenPhase3: {"assets/audio/siren_3.wav"},
	GhostSirenPhase4: {"assets/audio/siren_4.wav"},
	PowerPellet:      {"assets/audio/power_pellet.wav"},
	EatGhostEffect:   {"assets/audio/eat_ghost.wav"},
	Retreating:       {"assets/audio/retreating.wav"},
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

// StateEvent represents a type of event used to transition between states
type StateEvent int

// Scatter - Whenever a ghost starts scattering
// ChasePacman - Whenever a ghost starts chasing PacMan
// PowerPelletEaten - Whenever PacMan eats a power pellet
// PowerPelletWearOff - Whenever PacMan's power pellet wears off
// StartFlickering - A ghost will get imune to the pellet soon
// GhostEaten - Whenever pacman eats a ghost
// ReachBase = Whenever a ghost reaches base after retreating
const (
	Scatter StateEvent = iota
	ChasePacman
	PowerPelletEaten
	PowerPelletWearOff
	StartFlickering
	EatGhost
	ReachBase
)

// GhostState represents a ghost state
type GhostState int

// IdleState - Initial ghost value
// ScatterState - Normal ghost behavior
// ChaseState - Behavior to chase PacMan
// FleeingState - Fleeing PacMan
// FlickeringState - Still fleeing PacMan but about to stop
// EatenState - When the Ghost was just eaten by PacMan
const (
	IdleState GhostState = iota
	ScatterState
	ChaseState
	FleeingState
	FlickeringState
	EatenState
)

// PacmanState represents a ghost state
type PacmanState int

// WalkingState - Normal pacman behavior
// PowerState - Pacman behavior with power pellet
const (
	WalkingState PacmanState = iota
	PowerState
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
