package models

import (
	"log"
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
	position          interfaces.Location
	speed             int
	keyDirection      constants.Direction
	direction         constants.Direction
	sprites           *structures.SpriteSequence
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

// Refactored method to handle collisions.
// Retries once if collision is a wall to stop users from switching to a colliding direction
func (p *Pacman) handleCollisions(
	prevDirection constants.Direction,
	gameContext *contexts.GameContext,
) {
	target := p.collisionDetector.DetectCollision()
	switch target.(type) {
	case *Wall:
		if p.direction != prevDirection {
			p.direction = prevDirection
			p.handleCollisions(prevDirection, gameContext)
		}
	case *Food:
		// TODO: increment score, set appropriate state if food is super food
		gameContext.Maze.MoveElement(p, true)
		p.sprites.Advance()
		gameContext.SoundPlayer.PlayOnce(constants.MunchEffect)
		gameContext.Msg.EatFood <- struct{}{}
	default:
		gameContext.Maze.MoveElement(p, false)
		p.sprites.Advance()
	}
}

// Run the behavior of the player
func (p *Pacman) Run(gameContext *contexts.GameContext) {
	if p.collisionDetector == nil {
		log.Fatal("Collision detector is not attached")
	}

	prevDirection := p.direction
	go p.keyListener()
	for {
		gameContext.MazeMutex.Lock()
		if p.keyDirection != constants.DirStatic {
			p.direction = p.keyDirection
		}
		p.handleCollisions(prevDirection, gameContext)
		prevDirection = p.direction
		gameContext.MazeMutex.Unlock()
		time.Sleep(time.Duration(1000/p.speed) * time.Millisecond)
	}
}

// Draw the element to the screen in given position
func (p *Pacman) Draw(screen *ebiten.Image, x, y int) {
	p.animator.DrawFrame(screen, x, y)
}

// GetSprite of the element
func (p *Pacman) GetSprite() *ebiten.Image {
	return p.sprites.GetCurrentFrame()
}

// GetDirection of the element
func (p *Pacman) GetDirection() constants.Direction {
	return p.direction
}

// IsMatrixEditable based on the object direction
func (p *Pacman) IsMatrixEditable() bool {
	return true
}

// IsUnmovable by any force
func (p *Pacman) IsUnmovable() bool {
	return false
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
		position:     structures.InitPosition(x, y),
		speed:        constants.DefaultPacmanFPS,
		direction:    constants.DirLeft,
		keyDirection: constants.DirLeft,
	}
	sprites := []string{
		"assets/pacman/pacman-1.png",
		"assets/pacman/pacman-2.png",
		"assets/pacman/pacman-3.png",
		"assets/pacman/pacman-2.png",
	}
	seq, err := structures.InitSpriteSequence(sprites)
	pacman.sprites = seq
	pacman.animator = modules.InitAnimator(&pacman)
	return &pacman, err
}
