package models

import (
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

func getGhostStateInstance(
	state constants.GhostState,
	ghost *Ghost,
	ctx *contexts.GameContext,
) interfaces.GhostState {
	switch state {
	case constants.IdleState:
		return InitIdle(ghost, ctx)
	case constants.ScatterState:
		return InitScatter(ghost, ctx)
	case constants.ChaseState:
		return InitChase(ghost, ctx)
	default:
		return nil
	}
}

//----------------------------------------------------------------------------//
//------------------------------------IDLE------------------------------------//
//----------------------------------------------------------------------------//

// Idle state of a ghost
type Idle struct {
	ghost       *Ghost
	gameContext *contexts.GameContext
	transitions map[constants.StateEvent]constants.GhostState
	createdAt   time.Time
}

// ApplyTransition given an event
func (i *Idle) ApplyTransition(event constants.StateEvent) interfaces.GhostState {
	state, found := i.transitions[event]
	if !found {
		return i
	}

	return getGhostStateInstance(state, i.ghost, i.gameContext)
}

// Run main logic of state
func (i *Idle) Run() {
	i.ghost.advanceSprites()
	if time.Now().Sub(i.createdAt).Seconds() > i.ghost.idleStateTime {
		i.ghost.ChangeState(constants.Scatter)
	}
}

// GetSprite corresponding to state
func (i *Idle) GetSprite() *ebiten.Image {
	return i.ghost.orientedSprite()
}

// InitIdle state instance
func InitIdle(ghost *Ghost, ctx *contexts.GameContext) *Idle {
	idle := Idle{
		ghost:       ghost,
		gameContext: ctx,
		transitions: make(map[constants.StateEvent]constants.GhostState),
		createdAt:   time.Now(),
	}
	idle.transitions[constants.Scatter] = constants.ScatterState
	return &idle
}

//----------------------------------------------------------------------------//
//---------------------------------SCATTER------------------------------------//
//----------------------------------------------------------------------------//

// Scatter state of a ghost
type Scatter struct {
	state                    constants.GhostState
	ghost                    *Ghost
	gameContext              *contexts.GameContext
	transitions              map[constants.StateEvent]constants.GhostState
	createdAt                time.Time
	prevDirection            constants.Direction
	recentlyChangedDirection bool
}

// ApplyTransition given an event
func (s *Scatter) ApplyTransition(event constants.StateEvent) interfaces.GhostState {
	state, found := s.transitions[event]
	if !found {
		return s
	}

	return getGhostStateInstance(state, s.ghost, s.gameContext)
}

// Run main logic of state
func (s *Scatter) Run() {
	if !s.recentlyChangedDirection {
		s.ghost.attemptChangeDirection(nil)
	}
	s.recentlyChangedDirection = s.ghost.direction != s.prevDirection
	if !s.recentlyChangedDirection {
		target := s.ghost.collisionDetector.DetectCollision()
		switch target.(type) {
		case *Wall:
			s.ghost.direction = pickRandomDirection()
		default:
			s.gameContext.Maze.MoveElement(s.ghost, false)
			s.ghost.advanceSprites()
		}
	}
	s.prevDirection = s.ghost.direction
	if time.Now().Sub(s.createdAt).Seconds() > constants.ScatterModeDuration {
		s.ghost.ChangeState(constants.ChasePacman)
	}
}

// GetSprite corresponding to state
func (s *Scatter) GetSprite() *ebiten.Image {
	return s.ghost.orientedSprite()
}

// InitScatter state instance
func InitScatter(ghost *Ghost, ctx *contexts.GameContext) *Scatter {
	scatter := Scatter{
		ghost:                    ghost,
		gameContext:              ctx,
		transitions:              make(map[constants.StateEvent]constants.GhostState),
		createdAt:                time.Now(),
		prevDirection:            ghost.direction,
		recentlyChangedDirection: false,
	}
	scatter.transitions[constants.ChasePacman] = constants.ChaseState
	return &scatter
}

//----------------------------------------------------------------------------//
//----------------------------------CHASE-------------------------------------//
//----------------------------------------------------------------------------//

// Chase state of a ghost
type Chase struct {
	state                    constants.GhostState
	ghost                    *Ghost
	gameContext              *contexts.GameContext
	transitions              map[constants.StateEvent]constants.GhostState
	createdAt                time.Time
	prevDirection            constants.Direction
	recentlyChangedDirection bool
}

// ApplyTransition given an event
func (c *Chase) ApplyTransition(event constants.StateEvent) interfaces.GhostState {
	state, found := c.transitions[event]
	if !found {
		return c
	}

	return getGhostStateInstance(state, c.ghost, c.gameContext)
}

// Run main logic of state
func (c *Chase) Run() {
	if !c.recentlyChangedDirection {
		c.ghost.attemptChangeDirection(c.gameContext.MainPlayer.GetPosition())
	}
	c.recentlyChangedDirection = c.ghost.direction != c.prevDirection
	if !c.recentlyChangedDirection {
		target := c.ghost.collisionDetector.DetectCollision()
		switch target.(type) {
		case *Wall:
			c.ghost.direction = pickRandomDirection()
		default:
			c.gameContext.Maze.MoveElement(c.ghost, false)
			c.ghost.advanceSprites()
		}
	}
	c.prevDirection = c.ghost.direction
	timer := time.Now().Sub(c.createdAt).Seconds()
	if c.ghost.phase < constants.InfiniteChasePhase && timer > constants.ChaseModeDuration {
		c.ghost.phase++
		c.gameContext.Msg.PhaseChange <- c.ghost.phase
		c.ghost.ChangeState(constants.Scatter)
	}
}

// GetSprite corresponding to state
func (c *Chase) GetSprite() *ebiten.Image {
	return c.ghost.orientedSprite()
}

// InitChase state instance
func InitChase(ghost *Ghost, ctx *contexts.GameContext) *Chase {
	chase := Chase{
		ghost:                    ghost,
		gameContext:              ctx,
		transitions:              make(map[constants.StateEvent]constants.GhostState),
		createdAt:                time.Now(),
		prevDirection:            ghost.direction,
		recentlyChangedDirection: false,
	}
	chase.transitions[constants.Scatter] = constants.ScatterState
	return &chase
}
