package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Pacman represents the player
type Pacman struct {
	sprite *ebiten.Image
}

// InitPacman player for the level
func InitPacman() (*Pacman, error) {
	pacman := Pacman{}
	img, _, err := ebitenutil.NewImageFromFile("assets/pacman-1.png")
	pacman.sprite = img
	return &pacman, err
}
