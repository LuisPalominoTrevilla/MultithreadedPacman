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
	case constants.DeadState:
		return InitDead(pacman, ctx)
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
	targets := w.pacman.collisionDetector.DetectCollision()
targetsLoop:
	for _, target := range targets {
		switch obj := target.(type) {
		case *Wall, *Bars:
			if w.pacman.direction != w.prevDirection {
				w.pacman.direction = w.prevDirection
				w.handleCollisions()
			}
			return
		case *Pellet:
			w.pacman.EatPellet(obj, w.ctx)
		case *Ghost:
			if obj.AttemptEatPacman(w.pacman) {
				// Stop processing more targets if PacMan died
				break targetsLoop
			}
		}
	}
	w.pacman.sprites["alive"].Advance()
	w.ctx.Maze.MoveElement(w.pacman)
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
	return w.pacman.sprites["alive"].GetCurrentFrame()
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
	walking.transitions[constants.PacManEaten] = constants.DeadState
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
	targets := p.pacman.collisionDetector.DetectCollision()
targetsLoop:
	for _, target := range targets {
		switch obj := target.(type) {
		case *Wall, *Bars:
			if p.pacman.direction != p.prevDirection {
				p.pacman.direction = p.prevDirection
				p.handleCollisions()
			}
			return
		case *Pellet:
			p.pacman.EatPellet(obj, p.ctx)
		case *Ghost:
			if obj.AttemptEatPacman(p.pacman) {
				// Stop processing more targets if PacMan died
				break targetsLoop
			}
		}
	}
	p.pacman.sprites["alive"].Advance()
	p.ctx.Maze.MoveElement(p.pacman)
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
	return p.pacman.sprites["alive"].GetCurrentFrame()
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
	power.transitions[constants.PacManEaten] = constants.DeadState
	return &power
}

//----------------------------------------------------------------------------//
//-----------------------------------DEAD-------------------------------------//
//----------------------------------------------------------------------------//

// Dead state of the player
type Dead struct {
	pacman            *Pacman
	finishedAnimation bool
	ctx               *contexts.GameContext
}

// ApplyTransition given an event
func (w *Dead) ApplyTransition(event constants.StateEvent) interfaces.PacmanState {
	return w
}

// Run main logic of state
func (w *Dead) Run() {
	if !w.finishedAnimation {
		w.finishedAnimation = w.pacman.sprites["dead"].Advance()
	} else {
		w.ctx.Maze.RemoveElement(w.pacman)
		w.pacman.keepRunning = false
	}
}

// GetSprite corresponding to state
func (w *Dead) GetSprite() *ebiten.Image {
	if w.finishedAnimation {
		return nil
	}
	return w.pacman.sprites["dead"].GetCurrentFrame()
}

// InitDead state instance
func InitDead(pacman *Pacman, ctx *contexts.GameContext) *Dead {
	ctx.SoundPlayer.PlayOnce(constants.DyingEffect)
	ctx.Msg.GameOver <- struct{}{}
	return &Dead{
		finishedAnimation: false,
		pacman:            pacman,
		ctx:               ctx,
	}
}
