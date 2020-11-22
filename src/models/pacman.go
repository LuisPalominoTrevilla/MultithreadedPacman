package models

import (
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/hajimehoshi/ebiten/v2"
)

// Pacman represents the player
type Pacman struct {
	x         int
	y         int
	Direction modules.Direction
	Speed     int
	Sprites   *modules.SpriteSequence
}

func (p *Pacman) keyListener() {
	for {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.Direction = modules.DirUp
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.Direction = modules.DirDown
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.Direction = modules.DirRight
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.Direction = modules.DirLeft
		}

		time.Sleep(time.Duration(30) * time.Millisecond)
	}
}

// Run the behavior of the player
func (p *Pacman) Run(maze *Maze) {
	go p.keyListener()
	for {
		var dx, dy int
		switch p.Direction {
		case modules.DirUp:
			dx = 0
			dy = -1
		case modules.DirDown:
			dx = 0
			dy = 1
		case modules.DirLeft:
			dx = -1
			dy = 0
		case modules.DirRight:
			dx = 1
			dy = 0
		}

		// TODO: set mutex here
		maze.MoveElement(p.x, p.y, dx, dy)
		p.Sprites.Advance()
		time.Sleep(time.Duration(1000/p.Speed) * time.Millisecond)
	}
}

// InitPacman player for the level
func InitPacman(x, y int) (*Pacman, error) {
	pacman := Pacman{}
	sprites := []string{"src/assets/pacman-1.png", "src/assets/pacman-2.png", "src/assets/pacman-3.png", "src/assets/pacman-2.png"}
	seq, err := modules.InitSpriteSequence(sprites)
	pacman.Sprites = seq
	pacman.x = x
	pacman.y = y
	pacman.Direction = modules.DirLeft
	pacman.Speed = 5
	return &pacman, err
}
