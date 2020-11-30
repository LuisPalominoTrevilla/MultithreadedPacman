package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/screens"
	"github.com/hajimehoshi/ebiten/v2"
)

// GameController represents the main controller of the pacman game
type GameController struct {
	nEnemies     int
	screenWidth  int
	screenHeight int
	ctx          *contexts.AnchorContext
	activeScreen interfaces.Screen
	isActive     bool
}

func (g *GameController) run() {
	for {
		select {
		case newState := <-g.ctx.ChangeState:
			g.mountScreen(newState)
		}
	}
}

func (g *GameController) mountScreen(newState constants.GameState) {
	switch newState {
	case constants.MenuState:
		g.activeScreen = screens.NewMenu(g.screenWidth, g.screenHeight, g.ctx)
	case constants.PlayState:
		// Set active screen to the loading screen while the level screen is prepared
		g.activeScreen = screens.NewLoading(g.screenWidth, g.screenHeight)
		go func(controller *GameController) {
			level, err := screens.NewLevel("assets/level1.txt", g.nEnemies, g.ctx)
			if err != nil {
				log.Fatal(err)
			}
			controller.activeScreen = level
			controller.activeScreen.Run()
		}(g)
	}
	go g.activeScreen.Run()
}

// Draw the active screen
func (g *GameController) Draw(mainScreen *ebiten.Image) {
	if g.activeScreen != nil {
		g.activeScreen.Draw(mainScreen)
	}
}

// InitGame to start the logic
func (g *GameController) InitGame() {
	go g.run()
	g.isActive = true
	g.mountScreen(constants.MenuState)
}

// IsActive game
func (g *GameController) IsActive() bool {
	return g.isActive
}

// InitGameController instantiaes the main game controller
func InitGameController(nEnemies int) (*GameController, error) {
	if nEnemies <= 0 {
		return nil, errors.New("At least one enemy must be spawned")
	}
	if nEnemies > constants.MaxGhostsAllowed {
		errMsg := fmt.Sprintf("Cannot instantiate more than %d enemies", constants.MaxGhostsAllowed)
		return nil, errors.New(errMsg)
	}

	assetManager, err := modules.NewAssetManager()
	if err != nil {
		return nil, err
	}

	w := constants.HorizontalTiles * constants.TileSize
	h := constants.VerticalTiles * constants.TileSize
	gameController := GameController{
		nEnemies:     nEnemies,
		screenWidth:  w,
		screenHeight: h,
		ctx: &contexts.AnchorContext{
			ChangeState:  make(chan constants.GameState),
			AssetManager: assetManager,
		},
		isActive: false,
	}

	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Pacman")
	return &gameController, nil
}
