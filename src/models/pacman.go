package models

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Pacman represents the player
type Pacman struct {
	score             uint
	scoreMutex        sync.Mutex
	keepRunning       bool
	state             interfaces.PacmanState
	position          interfaces.Location
	speed             int
	keyDirection      constants.Direction
	direction         constants.Direction
	sprites           map[string]*structures.SpriteSequence
	animator          *modules.Animator
	collisionDetector *modules.CollisionDetector
}

func (p *Pacman) keyListener() {
	lastPressed := time.Now()
	for {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.keyDirection = constants.DirUp
			lastPressed = time.Now()
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.keyDirection = constants.DirDown
			lastPressed = time.Now()
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.keyDirection = constants.DirRight
			lastPressed = time.Now()
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.keyDirection = constants.DirLeft
			lastPressed = time.Now()
		}

		// Reset keyDirection if no key was pressed in the last 150 milliseconds
		if time.Now().Sub(lastPressed).Milliseconds() > 150 {
			p.keyDirection = constants.DirStatic
		}
		time.Sleep(time.Duration(30) * time.Millisecond)
	}
}

// ChangeState given an event
func (p *Pacman) ChangeState(event constants.StateEvent) {
	newState := p.state.ApplyTransition(event)
	if newState != nil {
		p.state = newState
	}
}

// EatPellet and perform logic
func (p *Pacman) EatPellet(pellet *Pellet, ctx *contexts.GameContext) {
	ctx.SoundPlayer.PlayOnce(constants.MunchEffect)
	ctx.Maze.RemoveElement(pellet)
	ctx.Msg.EatPellet <- pellet.isPowerful
	p.scoreMutex.Lock()
	if pellet.isPowerful {
		p.ChangeState(constants.PowerPelletEaten)
		p.score += 50
	} else {
		p.score += 10
	}
	p.scoreMutex.Unlock()
}

// EatGhost and send it back to hell
func (p *Pacman) EatGhost(g *Ghost, ctx *contexts.GameContext) {
	p.scoreMutex.Lock()
	p.score += 200
	p.scoreMutex.Unlock()
	ctx.SoundPlayer.PlayOnce(constants.EatGhostEffect)
	g.ChangeState(constants.GhostEaten)
}

// Run the behavior of the player
func (p *Pacman) Run(ctx *contexts.GameContext) {
	if p.collisionDetector == nil {
		log.Fatal("Collision detector is not attached")
	}

	p.state = InitWalking(p, ctx)
	go p.keyListener()
	for p.keepRunning {
		ctx.MazeMutex.Lock()
		p.state.Run()
		ctx.MazeMutex.Unlock()
		time.Sleep(time.Duration(1000/p.speed) * time.Millisecond)
	}
}

// Draw the element to the screen in given position
func (p *Pacman) Draw(screen *ebiten.Image, x, y int) {
	p.animator.DrawFrame(screen, x, y)
}

// GetSprite of the element
func (p *Pacman) GetSprite() *ebiten.Image {
	if p.state == nil {
		return p.sprites["alive"].GetCurrentFrame()
	}
	return p.state.GetSprite()
}

// GetDirection of the element
func (p *Pacman) GetDirection() constants.Direction {
	return p.direction
}

// IsMatrixEditable based on the object direction
func (p *Pacman) IsMatrixEditable() bool {
	return true
}

// CanGhostsGoThrough by any force
func (p *Pacman) CanGhostsGoThrough() bool {
	return true
}

// GetLayerIndex of the element
func (p *Pacman) GetLayerIndex() int {
	return constants.PacmanLayerIdx
}

// GetPosition of the element
func (p *Pacman) GetPosition() interfaces.Location {
	return p.position
}

// SetPosition of the element
func (p *Pacman) SetPosition(x, y int) {
	p.position.SetX(x)
	p.position.SetY(y)
}

// AttachCollisionDetector to the element
func (p *Pacman) AttachCollisionDetector(collisionDetector *modules.CollisionDetector) {
	p.collisionDetector = collisionDetector
}

// InitPacman player for the level
func InitPacman(x, y int) (*Pacman, error) {
	pacman := Pacman{
		score:        0,
		keepRunning:  true,
		position:     structures.InitPosition(x, y),
		speed:        constants.DefaultPacmanFPS,
		direction:    constants.DirLeft,
		keyDirection: constants.DirLeft,
		sprites:      make(map[string]*structures.SpriteSequence),
	}
	aliveSprites := []string{
		"assets/pacman/pacman-1.png",
		"assets/pacman/pacman-2.png",
		"assets/pacman/pacman-3.png",
		"assets/pacman/pacman-2.png",
	}
	deathSprites := make([]string, 11)
	for i := 0; i < 11; i++ {
		deathSprites[i] = fmt.Sprintf("assets/pacman/death-%d.png", i+1)
	}

	seq, err := structures.InitSpriteSequence(aliveSprites)
	if err != nil {
		return nil, err
	}
	pacman.sprites["alive"] = seq
	seq, err = structures.InitSpriteSequence(deathSprites)
	if err != nil {
		return nil, err
	}
	pacman.sprites["dead"] = seq

	pacman.animator = modules.InitAnimator(&pacman)
	return &pacman, nil
}
