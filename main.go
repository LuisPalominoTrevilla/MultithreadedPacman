package main

import (
	"flag"
	_ "image/png"
	"log"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/controller"
	"github.com/hajimehoshi/ebiten/v2"
)

var gameController *controller.GameController

func init() {
	nEnemies := flag.Int("n", 1, "Number of enemies to go against")
	flag.Parse()
	var err error
	gameController, err = controller.InitGameController(*nEnemies)
	if err != nil {
		log.Fatal(err)
	}
}

// Game represents an Ebite game instance
type Game struct{}

// Update game logic
func (g *Game) Update() error {
	if !gameController.IsActive() {
		gameController.InitGame()
	}
	return nil
}

// Draw frame by frame the scene
func (g *Game) Draw(screen *ebiten.Image) {
	gameController.Draw(screen)
}

// Layout of the game
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
