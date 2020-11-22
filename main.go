package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image
var level *Level

const tileWidth = 48
const tileHeight = 48

func init() {
	var err error
	level, err = InitLevel("level1.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(level.layout)

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
	op.GeoM.Scale(tileWidth/float64(w), tileHeight/float64(h))
	screen.DrawImage(img, op)
}

// Layout of the game
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pacman")
	// ebiten.SetScreenClearedEveryFrame(false)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
