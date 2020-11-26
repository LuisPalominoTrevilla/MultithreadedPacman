package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

func getPacmanStateInstance(
	state constants.PacmanState,
	pacman *Pacman,
	ctx *contexts.GameContext,
) interfaces.PacmanState {
	switch state {
	case constants.WalkingState:
		return InitWalking(pacman, ctx)
	// case constants.PowerState:
	// return InitPower(pacman, ctx)
	default:
		return nil
	}
}

//----------------------------------------------------------------------------//
//--------------------------------- WALKING-----------------------------------//
//----------------------------------------------------------------------------//

// Walking state of the player
type Walking struct {
	pacman        *Pacman
	gameContext   *contexts.GameContext
	transitions   map[constants.StateEvent]constants.PacmanState
	prevDirection constants.Direction
}

// ApplyTransition given an event
func (i *Walking) ApplyTransition(event constants.StateEvent) interfaces.PacmanState {
	state, found := i.transitions[event]
	if !found {
		return i
	}

	return getPacmanStateInstance(state, i.pacman, i.gameContext)
}

func (i *Walking) handleCollisions() {
	target := i.pacman.collisionDetector.DetectCollision()
	switch target.(type) {
	case *Wall:
		if i.pacman.direction != i.prevDirection {
			i.pacman.direction = i.prevDirection
			i.handleCollisions()
		}
	case *Pellet:
		// TODO: increment score, set appropriate state if pellet was power pellet
		i.gameContext.SoundPlayer.PlayOnce(constants.MunchEffect)
		i.gameContext.Maze.MoveElement(i.pacman, true)
		i.pacman.sprites.Advance()
		i.gameContext.Msg.EatPellet <- struct{}{}
	default:
		i.gameContext.Maze.MoveElement(i.pacman, false)
		i.pacman.sprites.Advance()
	}
}

// Run main logic of state
func (i *Walking) Run() {
	if i.pacman.keyDirection != constants.DirStatic {
		i.pacman.direction = i.pacman.keyDirection
	}
	i.handleCollisions()
	i.prevDirection = i.pacman.direction
}

// GetSprite corresponding to state
func (i *Walking) GetSprite() *ebiten.Image {
	return i.pacman.sprites.GetCurrentFrame()
}

// InitWalking state instance
func InitWalking(pacman *Pacman, ctx *contexts.GameContext) *Walking {
	walking := Walking{
		pacman:        pacman,
		gameContext:   ctx,
		transitions:   make(map[constants.StateEvent]constants.PacmanState),
		prevDirection: pacman.direction,
	}
	// walking.transitions[constants.Scatter] = constants
	return &walking
}
