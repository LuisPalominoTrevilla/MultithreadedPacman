package models

import (
	"time"

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
	case constants.PowerState:
		return InitPower(pacman, ctx)
	default:
		return nil
	}
}

//----------------------------------------------------------------------------//
//----------------------------------WALKING-----------------------------------//
//----------------------------------------------------------------------------//

// Walking state of the player
type Walking struct {
	pacman        *Pacman
	ctx           *contexts.GameContext
	transitions   map[constants.StateEvent]constants.PacmanState
	prevDirection constants.Direction
}

// ApplyTransition given an event
func (w *Walking) ApplyTransition(event constants.StateEvent) interfaces.PacmanState {
	state, found := w.transitions[event]
	if !found {
		return w
	}

	return getPacmanStateInstance(state, w.pacman, w.ctx)
}

func (w *Walking) handleCollisions() {
	target := w.pacman.collisionDetector.DetectCollision()
	switch obj := target.(type) {
	case *Wall:
		if w.pacman.direction != w.prevDirection {
			w.pacman.direction = w.prevDirection
			w.handleCollisions()
		}
	case *Pellet:
		w.ctx.SoundPlayer.PlayOnce(constants.MunchEffect)
		w.ctx.Maze.MoveElement(w.pacman)
		w.ctx.Maze.RemoveElement(obj)
		w.pacman.sprites.Advance()
		w.ctx.Msg.EatPellet <- obj.isPowerful
		if obj.isPowerful {
			w.pacman.ChangeState(constants.PowerPelletEaten)
		}
	case *Ghost:
		obj.AttemptEatPacman(w.pacman)
		w.ctx.Maze.MoveElement(w.pacman)
		w.pacman.sprites.Advance()
	default:
		w.ctx.Maze.MoveElement(w.pacman)
		w.pacman.sprites.Advance()
	}
}

// Run main logic of state
func (w *Walking) Run() {
	if w.pacman.keyDirection != constants.DirStatic {
		w.pacman.direction = w.pacman.keyDirection
	}
	w.handleCollisions()
	w.prevDirection = w.pacman.direction
}

// GetSprite corresponding to state
func (w *Walking) GetSprite() *ebiten.Image {
	return w.pacman.sprites.GetCurrentFrame()
}

// InitWalking state instance
func InitWalking(pacman *Pacman, ctx *contexts.GameContext) *Walking {
	pacman.speed = constants.DefaultPacmanFPS
	walking := Walking{
		pacman:        pacman,
		ctx:           ctx,
		transitions:   make(map[constants.StateEvent]constants.PacmanState),
		prevDirection: pacman.direction,
	}
	walking.transitions[constants.PowerPelletEaten] = constants.PowerState
	return &walking
}

//----------------------------------------------------------------------------//
//-----------------------------------POWER------------------------------------//
//----------------------------------------------------------------------------//

// Power state of the player
type Power struct {
	pacman        *Pacman
	ctx           *contexts.GameContext
	transitions   map[constants.StateEvent]constants.PacmanState
	createdAt     time.Time
	prevDirection constants.Direction
}

// ApplyTransition given an event
func (p *Power) ApplyTransition(event constants.StateEvent) interfaces.PacmanState {
	state, found := p.transitions[event]
	if !found {
		return p
	}

	return getPacmanStateInstance(state, p.pacman, p.ctx)
}

func (p *Power) handleCollisions() {
	target := p.pacman.collisionDetector.DetectCollision()
	switch obj := target.(type) {
	case *Wall:
		if p.pacman.direction != p.prevDirection {
			p.pacman.direction = p.prevDirection
			p.handleCollisions()
		}
	case *Pellet:
		// TODO: increment score, set appropriate state if pellet was power pellet
		p.ctx.SoundPlayer.PlayOnce(constants.MunchEffect)
		p.ctx.Maze.MoveElement(p.pacman)
		p.ctx.Maze.RemoveElement(obj)
		p.pacman.sprites.Advance()
		p.ctx.Msg.EatPellet <- obj.isPowerful
		if obj.isPowerful {
			p.pacman.ChangeState(constants.PowerPelletEaten)
		}
	case *Ghost:
		obj.AttemptEatPacman(p.pacman)
		p.ctx.Maze.MoveElement(p.pacman)
		p.pacman.sprites.Advance()
	default:
		p.ctx.Maze.MoveElement(p.pacman)
		p.pacman.sprites.Advance()
	}
}

// Run main logic of state
func (p *Power) Run() {
	if p.pacman.keyDirection != constants.DirStatic {
		p.pacman.direction = p.pacman.keyDirection
	}
	p.handleCollisions()
	p.prevDirection = p.pacman.direction
	timer := time.Now().Sub(p.createdAt).Seconds()
	if timer > constants.PowerPelletDuration {
		p.ctx.Msg.PowerPelletWoreOff <- struct{}{}
		p.pacman.ChangeState(constants.PowerPelletWearOff)
	}
}

// GetSprite corresponding to state
func (p *Power) GetSprite() *ebiten.Image {
	return p.pacman.sprites.GetCurrentFrame()
}

// InitPower state instance
func InitPower(pacman *Pacman, ctx *contexts.GameContext) *Power {
	pacman.speed = constants.PowerPacmanFPS
	power := Power{
		pacman:        pacman,
		ctx:           ctx,
		transitions:   make(map[constants.StateEvent]constants.PacmanState),
		createdAt:     time.Now(),
		prevDirection: pacman.direction,
	}
	power.transitions[constants.PowerPelletEaten] = constants.PowerState
	power.transitions[constants.PowerPelletWearOff] = constants.WalkingState
	return &power
}
