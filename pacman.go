package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Pacman represents the player
type Pacman struct {
	x         int
	y         int
	direction Direction
	speed     int
	sprite    *ebiten.Image
}

func (p *Pacman) keyListener() {
	for {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.direction = DirUp
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.direction = DirDown
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.direction = DirRight
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.direction = DirLeft
		}

		time.Sleep(time.Duration(30) * time.Millisecond)
	}
}

// Run the behavior of the player
func (p *Pacman) Run(game *Level) {
	go p.keyListener()
	for {
		var dx, dy int
		switch p.direction {
		case DirUp:
			dx = 0
			dy = -1
		case DirDown:
			dx = 0
			dy = 1
		case DirLeft:
			dx = -1
			dy = 0
		case DirRight:
			dx = 1
			dy = 0
		}

		// TODO: set mutex here
		game.maze.MoveElement(p.x, p.y, dx, dy)
		time.Sleep(time.Duration(1000/p.speed) * time.Millisecond)
	}
}

// InitPacman player for the level
func InitPacman(x, y int) (*Pacman, error) {
	pacman := Pacman{}
	img, _, err := ebitenutil.NewImageFromFile("assets/pacman-1.png")
	pacman.sprite = img
	pacman.x = x
	pacman.y = y
	pacman.direction = DirLeft
	pacman.speed = 5
	return &pacman, err
}
