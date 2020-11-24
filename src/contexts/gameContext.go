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
	SoundPlayer *modules.SoundPlayer
	Msg         chan constants.EventType
}
