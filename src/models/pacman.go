package models

import (
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// Pacman represents the player
type Pacman struct {
	x         int
	y         int
	speed     int
	direction utils.Direction
	sprites   *modules.SpriteSequence
	animator  *modules.Animator
}

func (p *Pacman) keyListener() {
	for {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.direction = utils.DirUp
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.direction = utils.DirDown
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.direction = utils.DirRight
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.direction = utils.DirLeft
		}

		time.Sleep(time.Duration(30) * time.Millisecond)
	}
}

// Run the behavior of the player
func (p *Pacman) Run(maze *Maze) {
	go p.keyListener()
	for {
		var dx, dy int
		switch p.direction {
		case utils.DirUp:
			dx = 0
			dy = -1
		case utils.DirDown:
			dx = 0
			dy = 1
		case utils.DirLeft:
			dx = -1
			dy = 0
		case utils.DirRight:
			dx = 1
			dy = 0
		}

		// TODO: set mutex here
		maze.MoveElement(p.x, p.y, dx, dy)
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
func (p *Pacman) GetDirection() utils.Direction {
	return p.direction
}

// InitPacman player for the level
func InitPacman(x, y int) (*Pacman, error) {
	pacman := Pacman{}
	sprites := []string{"src/assets/pacman-1.png", "src/assets/pacman-2.png", "src/assets/pacman-3.png", "src/assets/pacman-2.png"}
	seq, err := modules.InitSpriteSequence(sprites)
	pacman.x = x
	pacman.y = y
	pacman.speed = 5
	pacman.direction = utils.DirLeft
	pacman.sprites = seq
	pacman.animator = modules.InitAnimator(&pacman)
	return &pacman, err
}
