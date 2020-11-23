package controller

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/screens"
	"github.com/hajimehoshi/ebiten/v2"
)

// GameController represents the main controller of the pacman game
type GameController struct {
	state   constants.GameState
	screens map[constants.GameState]interfaces.Screen
}

func (g *GameController) run() {
}

func (g *GameController) currentScreen() interfaces.Screen {
	screen, ok := g.screens[g.state]
	if !ok {
		return nil
	}
	return screen
}

func (g *GameController) switchScreen(newState constants.GameState) {
	// TODO: If current screen is not null, create a method to stop all execution
	g.state = newState
	screen := g.currentScreen()
	if screen == nil {
		return
	}
	screen.Run()
}

// Draw TEMPORAL
func (g *GameController) Draw(mainScreen *ebiten.Image) {
	screen := g.currentScreen()
	screen.Draw(mainScreen)
}

// InitGame to start the logic
func (g *GameController) InitGame() {
	go g.run()
	g.switchScreen(constants.PlayState)
}

// State of the game
func (g *GameController) State() constants.GameState {
	return g.state
}

// InitGameController instantiaes the main game controller
func InitGameController() (*GameController, error) {
	gameController := GameController{
		state:   constants.InactiveState,
		screens: make(map[constants.GameState]interfaces.Screen),
	}

	level, err := screens.InitLevel("src/assets/level1.txt")
	if err != nil {
		return nil, err
	}

	gameController.screens[constants.PlayState] = level
	w, h := level.Size()
	ebiten.SetWindowSize(constants.TileSize*w, constants.TileSize*h)
	ebiten.SetWindowTitle("Pacman")
	return &gameController, nil
}
