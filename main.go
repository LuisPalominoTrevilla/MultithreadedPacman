package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image
var level *Level

func init() {
	var err error
	level, err = InitLevel("level1.txt")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err = ebitenutil.NewImageFromFile("assets/wall.png")
	if err != nil {
		log.Fatal(err)
	}
}

// Game represents an Ebite game instance
type Game struct{}

// Update game logic
func (g *Game) Update() error {
	return nil
}

// Draw frame by frame the scene
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := img.Size()
	op.GeoM.Scale(TileWidth/float64(w), TileHeight/float64(h))
	screen.DrawImage(img, op)
}

// Layout of the game
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	level.Draw()
	ebiten.SetWindowSize(TileWidth*level.maze.cols, TileHeight*level.maze.rows)
	ebiten.SetWindowTitle("Pacman")
	// ebiten.SetScreenClearedEveryFrame(false)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
