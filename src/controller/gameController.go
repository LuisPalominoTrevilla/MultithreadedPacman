package controller

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/screens"
	"github.com/hajimehoshi/ebiten/v2"
)

// GameController represents the main controller of the pacman game
type GameController struct {
	changeState chan constants.GameState
	state       constants.GameState
	screens     map[constants.GameState]interfaces.Screen
}

func (g *GameController) run() {
	for {
		select {
		case newState := <-g.changeState:
			g.switchScreen(newState)
		}
	}
}

func (g *GameController) currentScreen() interfaces.Screen {
	screen, ok := g.screens[g.state]
	if !ok {
		return nil
	}
	return screen
}

func (g *GameController) switchScreen(newState constants.GameState) {
	g.state = newState
	screen := g.currentScreen()
	if screen == nil {
		return
	}
	go screen.Run(g.changeState)
}

// Draw TEMPORAL
func (g *GameController) Draw(mainScreen *ebiten.Image) {
	screen := g.currentScreen()
	screen.Draw(mainScreen)
}

// InitGame to start the logic
func (g *GameController) InitGame() {
	go g.run()
	g.switchScreen(constants.MenuState)
}

// State of the game
func (g *GameController) State() constants.GameState {
	return g.state
}

// InitGameController instantiaes the main game controller
func InitGameController(nEnemies int) (*GameController, error) {
	gameController := GameController{
		changeState: make(chan constants.GameState),
		state:       constants.InactiveState,
		screens:     make(map[constants.GameState]interfaces.Screen),
	}

	level, err := screens.NewLevel("assets/level1.txt", nEnemies)
	if err != nil {
		return nil, err
	}

	w, h := level.Size()
	menu := screens.NewMenu(constants.TileSize*w, constants.TileSize*h)
	gameController.screens[constants.PlayState] = level
	gameController.screens[constants.MenuState] = menu
	ebiten.SetWindowSize(constants.TileSize*w, constants.TileSize*h)
	ebiten.SetWindowTitle("Pacman")
	return &gameController, nil
}
