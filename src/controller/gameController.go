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
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
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
		g.activeScreen = screens.NewLoading(g.screenWidth, g.screenHeight, g.ctx)
		go func(controller *GameController) {
			level, err := screens.NewLevel("assets/level1.txt", g.nEnemies, g.ctx)
			if err != nil {
				log.Fatal(err)
			}
			controller.activeScreen = level
			controller.activeScreen.Run()
		}(g)
	case constants.GameOverState:
		g.activeScreen = screens.NewGameOver(g.screenWidth, g.screenHeight, g.ctx)
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

	soundPlayer, err := modules.InitSoundPlayer()
	if err != nil {
		return nil, err
	}

	tt, err := truetype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		return nil, err
	}

	fontFace := truetype.NewFace(tt, &truetype.Options{
		Size: 30, DPI: 72, Hinting: font.HintingFull,
	})

	w := constants.HorizontalTiles * constants.TileSize
	h := constants.VerticalTiles * constants.TileSize + 100
	gameController := GameController{
		nEnemies:     nEnemies,
		screenWidth:  w,
		screenHeight: h,
		ctx: &contexts.AnchorContext{
			ChangeState:  make(chan constants.GameState),
			AssetManager: assetManager,
			SoundPlayer:  soundPlayer,
			FontFace:     fontFace,
		},
		isActive: false,
	}

	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Pacman")
	return &gameController, nil
}
