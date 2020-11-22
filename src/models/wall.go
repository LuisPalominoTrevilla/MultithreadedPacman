package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Wall represents a wall
type Wall struct {
	Sprite *ebiten.Image
}

// InitWall of the maze
func InitWall() (*Wall, error) {
	wall := Wall{}
	img, _, err := ebitenutil.NewImageFromFile("src/assets/wall.png")
	wall.Sprite = img
	return &wall, err
}
