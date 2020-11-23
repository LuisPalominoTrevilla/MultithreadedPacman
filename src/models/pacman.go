package models

import (
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/hajimehoshi/ebiten/v2"
)

// Pacman represents the player
type Pacman struct {
	x         int
	y         int
	speed     int
	direction constants.Direction
	sprites   *modules.SpriteSequence
	animator  *modules.Animator
}

func (p *Pacman) keyListener() {
	for {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.direction = constants.DirUp
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.direction = constants.DirDown
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.direction = constants.DirRight
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.direction = constants.DirLeft
		}

		time.Sleep(time.Duration(30) * time.Millisecond)
	}
}

// Run the behavior of the player
func (p *Pacman) Run(maze *Maze) {
	go p.keyListener()
	for {
		// TODO: set mutex here
		maze.MoveElement(p.x, p.y, p.direction)
		p.sprites.Advance()
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

// InitPacman player for the level
func InitPacman(x, y int) (*Pacman, error) {
	pacman := Pacman{}
	sprites := []string{
		"src/assets/pacman/pacman-1.png",
		"src/assets/pacman/pacman-2.png",
		"src/assets/pacman/pacman-3.png",
		"src/assets/pacman/pacman-2.png",
	}
	seq, err := modules.InitSpriteSequence(sprites)
	pacman.x = x
	pacman.y = y
	pacman.speed = 5
	pacman.direction = constants.DirLeft
	pacman.sprites = seq
	pacman.animator = modules.InitAnimator(&pacman)
	return &pacman, err
}
