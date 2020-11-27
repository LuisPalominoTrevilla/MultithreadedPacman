package models

import (
	"fmt"
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
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
	case constants.FleeingState:
		return InitFleeing(ghost, ctx)
	case constants.FlickeringState:
		return InitFlickering(ghost, ctx)
	case constants.EatenState:
		return InitEaten(ghost, ctx)
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

// AttemptEatPacman given the current state
func (i *Idle) AttemptEatPacman(obj interfaces.MovableGameObject) bool {
	return false
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

// AttemptEatPacman given the current state
func (s *Scatter) AttemptEatPacman(obj interfaces.MovableGameObject) bool {
	_, ok := obj.(*Pacman)
	if !ok {
		return false
	}

	fmt.Println("Pacman dead")
	// TODO: Change pacman state to eaten
	return true
}

// Run main logic of state
func (s *Scatter) Run() {
	if !s.recentlyChangedDirection {
		s.ghost.turnTowards(nil, false, true)
	}
	s.recentlyChangedDirection = s.ghost.direction != s.prevDirection
	if !s.recentlyChangedDirection {
		target := s.ghost.collisionDetector.DetectCollision()
		switch obj := target.(type) {
		case *Wall:
			s.ghost.direction = pickRandomDirection()
		case *Pacman:
			s.AttemptEatPacman(obj)
			s.gameContext.Maze.MoveElement(s.ghost, false)
			s.ghost.advanceSprites()
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
	ghost.speed = constants.DefaultGhostFPS
	scatter := Scatter{
		ghost:                    ghost,
		gameContext:              ctx,
		transitions:              make(map[constants.StateEvent]constants.GhostState),
		createdAt:                time.Now(),
		prevDirection:            ghost.direction,
		recentlyChangedDirection: false,
	}
	scatter.transitions[constants.ChasePacman] = constants.ChaseState
	scatter.transitions[constants.PowerPelletEaten] = constants.FleeingState
	return &scatter
}

//----------------------------------------------------------------------------//
//----------------------------------CHASE-------------------------------------//
//----------------------------------------------------------------------------//

// Chase state of a ghost
type Chase struct {
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

// AttemptEatPacman given the current state
func (c *Chase) AttemptEatPacman(obj interfaces.MovableGameObject) bool {
	_, ok := obj.(*Pacman)
	if !ok {
		return false
	}

	fmt.Println("Pacman dead")
	// TODO: Change pacman state to eaten
	return true
}

// Run main logic of state
func (c *Chase) Run() {
	if !c.recentlyChangedDirection {
		c.ghost.turnTowards(c.gameContext.MainPlayer.GetPosition(), false, true)
	}
	c.recentlyChangedDirection = c.ghost.direction != c.prevDirection
	if !c.recentlyChangedDirection {
		target := c.ghost.collisionDetector.DetectCollision()
		switch obj := target.(type) {
		case *Wall:
			c.ghost.direction = pickRandomDirection()
		case *Pacman:
			c.AttemptEatPacman(obj)
			c.gameContext.Maze.MoveElement(c.ghost, false)
			c.ghost.advanceSprites()
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
	ghost.speed = constants.DefaultGhostFPS
	chase := Chase{
		ghost:                    ghost,
		gameContext:              ctx,
		transitions:              make(map[constants.StateEvent]constants.GhostState),
		createdAt:                time.Now(),
		prevDirection:            ghost.direction,
		recentlyChangedDirection: false,
	}
	chase.transitions[constants.Scatter] = constants.ScatterState
	chase.transitions[constants.PowerPelletEaten] = constants.FleeingState
	return &chase
}

//----------------------------------------------------------------------------//
//---------------------------------FLEEING------------------------------------//
//----------------------------------------------------------------------------//

// Fleeing state of a ghost
type Fleeing struct {
	ghost                    *Ghost
	gameContext              *contexts.GameContext
	transitions              map[constants.StateEvent]constants.GhostState
	createdAt                time.Time
	prevDirection            constants.Direction
	blockReverse             bool
	recentlyChangedDirection bool
}

// ApplyTransition given an event
func (f *Fleeing) ApplyTransition(event constants.StateEvent) interfaces.GhostState {
	state, found := f.transitions[event]
	if !found {
		return f
	}

	return getGhostStateInstance(state, f.ghost, f.gameContext)
}

// AttemptEatPacman given the current state
func (f *Fleeing) AttemptEatPacman(obj interfaces.MovableGameObject) bool {
	pacman, ok := obj.(*Pacman)
	if !ok {
		return false
	}
	pacman.EatGhost(f.ghost, f.gameContext)
	return false
}

// Run main logic of state
func (f *Fleeing) Run() {
	if !f.recentlyChangedDirection {
		f.ghost.turnTowards(f.gameContext.MainPlayer.GetPosition(), true, f.blockReverse)
		f.blockReverse = true
	}
	f.recentlyChangedDirection = f.ghost.direction != f.prevDirection
	if !f.recentlyChangedDirection {
		target := f.ghost.collisionDetector.DetectCollision()
		switch obj := target.(type) {
		case *Wall:
			f.ghost.direction = pickRandomDirection()
		case *Pacman:
			f.AttemptEatPacman(obj)
			f.gameContext.Maze.MoveElement(f.ghost, false)
			f.ghost.advanceSprites()
		default:
			f.gameContext.Maze.MoveElement(f.ghost, false)
			f.ghost.advanceSprites()
		}
	}
	f.prevDirection = f.ghost.direction
	timer := time.Now().Sub(f.createdAt).Seconds()
	if timer > constants.PowerPelletDuration-constants.FlickeringStateDuration {
		f.ghost.ChangeState(constants.StartFlickering)
	}
}

// GetSprite corresponding to state
func (f *Fleeing) GetSprite() *ebiten.Image {
	return f.ghost.sprites["panic"].GetCurrentFrame()
}

// InitFleeing state instance
func InitFleeing(ghost *Ghost, ctx *contexts.GameContext) *Fleeing {
	ghost.speed = constants.FleeingGhostFPS
	fleeing := Fleeing{
		ghost:                    ghost,
		gameContext:              ctx,
		transitions:              make(map[constants.StateEvent]constants.GhostState),
		createdAt:                time.Now(),
		prevDirection:            ghost.direction,
		blockReverse:             false,
		recentlyChangedDirection: false,
	}
	fleeing.transitions[constants.StartFlickering] = constants.FlickeringState
	fleeing.transitions[constants.EatGhost] = constants.EatenState
	fleeing.transitions[constants.PowerPelletEaten] = constants.FleeingState
	return &fleeing
}

//----------------------------------------------------------------------------//
//--------------------------------Flickering----------------------------------//
//----------------------------------------------------------------------------//

// Flickering state of a ghost
type Flickering struct {
	ghost                    *Ghost
	gameContext              *contexts.GameContext
	transitions              map[constants.StateEvent]constants.GhostState
	createdAt                time.Time
	prevDirection            constants.Direction
	recentlyChangedDirection bool
}

// ApplyTransition given an event
func (f *Flickering) ApplyTransition(event constants.StateEvent) interfaces.GhostState {
	state, found := f.transitions[event]
	if !found {
		return f
	}

	return getGhostStateInstance(state, f.ghost, f.gameContext)
}

// AttemptEatPacman given the current state
func (f *Flickering) AttemptEatPacman(obj interfaces.MovableGameObject) bool {
	pacman, ok := obj.(*Pacman)
	if !ok {
		return false
	}
	pacman.EatGhost(f.ghost, f.gameContext)
	return false
}

// Run main logic of state
func (f *Flickering) Run() {
	if !f.recentlyChangedDirection {
		f.ghost.turnTowards(f.gameContext.MainPlayer.GetPosition(), true, true)
	}
	f.recentlyChangedDirection = f.ghost.direction != f.prevDirection
	if !f.recentlyChangedDirection {
		target := f.ghost.collisionDetector.DetectCollision()
		switch obj := target.(type) {
		case *Wall:
			f.ghost.direction = pickRandomDirection()
		case *Pacman:
			f.AttemptEatPacman(obj)
			f.gameContext.Maze.MoveElement(f.ghost, false)
			f.ghost.advanceSprites()
		default:
			f.gameContext.Maze.MoveElement(f.ghost, false)
			f.ghost.advanceSprites()
		}
	}
	f.prevDirection = f.ghost.direction
	timer := time.Now().Sub(f.createdAt).Seconds()
	if timer > constants.FlickeringStateDuration {
		f.ghost.ChangeState(constants.PowerPelletWearOff)
	}
}

// GetSprite corresponding to state
func (f *Flickering) GetSprite() *ebiten.Image {
	return f.ghost.sprites["flicker"].GetCurrentFrame()
}

// InitFlickering state instance
func InitFlickering(ghost *Ghost, ctx *contexts.GameContext) *Flickering {
	ghost.speed = constants.FleeingGhostFPS
	flickering := Flickering{
		ghost:                    ghost,
		gameContext:              ctx,
		transitions:              make(map[constants.StateEvent]constants.GhostState),
		createdAt:                time.Now(),
		prevDirection:            ghost.direction,
		recentlyChangedDirection: false,
	}
	flickering.transitions[constants.PowerPelletWearOff] = constants.ScatterState
	flickering.transitions[constants.EatGhost] = constants.EatenState
	flickering.transitions[constants.PowerPelletEaten] = constants.FleeingState
	return &flickering
}

//----------------------------------------------------------------------------//
//----------------------------------EATEN-------------------------------------//
//----------------------------------------------------------------------------//

// Eaten state of a ghost
type Eaten struct {
	ghost         *Ghost
	gameContext   *contexts.GameContext
	transitions   map[constants.StateEvent]constants.GhostState
	prevDirection constants.Direction
	audioEffect   *modules.InfiniteAudio
}

// ApplyTransition given an event
func (e *Eaten) ApplyTransition(event constants.StateEvent) interfaces.GhostState {
	state, found := e.transitions[event]
	if !found {
		return e
	}

	e.audioEffect.Stop()
	return getGhostStateInstance(state, e.ghost, e.gameContext)
}

// AttemptEatPacman given the current state
func (e *Eaten) AttemptEatPacman(obj interfaces.MovableGameObject) bool {
	return false
}

// Run main logic of state
func (e *Eaten) Run() {
	e.ghost.turnTowards(e.gameContext.GhostBase, false, true)
	target := e.ghost.collisionDetector.DetectCollision()
	switch obj := target.(type) {
	case *Wall:
		e.ghost.direction = pickRandomDirection()
	case *Pacman:
		e.AttemptEatPacman(obj)
		e.gameContext.Maze.MoveElement(e.ghost, false)
		e.ghost.advanceSprites()
	default:
		e.gameContext.Maze.MoveElement(e.ghost, false)
		e.ghost.advanceSprites()
	}
	e.prevDirection = e.ghost.direction
	if e.ghost.position.DistanceTo(e.gameContext.GhostBase) < 1 {
		e.ghost.ChangeState(constants.ReachBase)
	}
}

// GetSprite corresponding to state
func (e *Eaten) GetSprite() *ebiten.Image {
	switch e.ghost.direction {
	case constants.DirUp:
		return e.ghost.sprites["eaten-up"].GetCurrentFrame()
	case constants.DirDown:
		return e.ghost.sprites["eaten-down"].GetCurrentFrame()
	case constants.DirLeft:
		return e.ghost.sprites["eaten-left"].GetCurrentFrame()
	case constants.DirRight:
		return e.ghost.sprites["eaten-right"].GetCurrentFrame()
	default:
		return e.ghost.sprites["eaten-left"].GetCurrentFrame()
	}
}

// InitEaten state instance
func InitEaten(ghost *Ghost, ctx *contexts.GameContext) *Eaten {
	ghost.speed = constants.EatenGhostFPS
	eaten := Eaten{
		ghost:         ghost,
		gameContext:   ctx,
		transitions:   make(map[constants.StateEvent]constants.GhostState),
		prevDirection: ghost.direction,
		audioEffect:   ctx.SoundPlayer.PlayOnLoop(constants.Retreating),
	}
	eaten.transitions[constants.ReachBase] = constants.ScatterState
	return &eaten
}
