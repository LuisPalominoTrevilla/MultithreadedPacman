package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Wall represents a wall
type Wall struct {
	sprite *ebiten.Image
}

// InitWall of the maze
func InitWall() (*Wall, error) {
	wall := Wall{}
	img, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
	wall.sprite = img
	return &wall, err
}
