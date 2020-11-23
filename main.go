package main

import (
	_ "image/png"
	"log"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/controller"
	"github.com/hajimehoshi/ebiten/v2"
)

var gameController *controller.GameController

func init() {
	var err error
	gameController, err = controller.InitGameController()
	if err != nil {
		log.Fatal(err)
	}
}

// Game represents an Ebite game instance
type Game struct{}

// Update game logic
func (g *Game) Update() error {
	if gameController.State() == constants.InactiveState {
		go gameController.InitGame()
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
	// ebiten.SetScreenClearedEveryFrame(false)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
