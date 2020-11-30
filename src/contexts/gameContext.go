package contexts

import (
	"sync"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
)

// GameContext represents the game context
type GameContext struct {
	MainPlayer  interfaces.MovableGameObject
	MazeMutex   sync.Mutex
	Maze        *structures.Maze
	GhostHome   interfaces.Location
	GhostBases  map[constants.GhostType]interfaces.Location
	SoundPlayer *modules.SoundPlayer
	Msg         *structures.MessageBroker
}

// AnchorContext represents the game context shared among screens
type AnchorContext struct {
	ChangeState  chan constants.GameState
	AssetManager *modules.AssetManager
}
