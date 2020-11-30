package modules

import (
	"fmt"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// AssetManager for all shared assets in the game
type AssetManager struct {
	PacmanSprites     map[string]*structures.SpriteSequence
	WallSprite        *ebiten.Image
	BarsSprite        *ebiten.Image
	PelletSprite      *ebiten.Image
	PowerPelletSprite *ebiten.Image
}

// NewAssetManager for the game
func NewAssetManager() (*AssetManager, error) {
	am := &AssetManager{
		PacmanSprites: make(map[string]*structures.SpriteSequence),
	}

	wallSprite, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
	if err != nil {
		return nil, err
	}

	barsSprite, _, err := ebitenutil.NewImageFromFile("assets/bars.png")
	if err != nil {
		return nil, err
	}

	pelletSprite, _, err := ebitenutil.NewImageFromFile("assets/pellet.png")
	if err != nil {
		return nil, err
	}

	powerPelletSprite, _, err := ebitenutil.NewImageFromFile("assets/power-pellet.png")
	if err != nil {
		return nil, err
	}

	aliveSrc := []string{
		"assets/pacman/pacman-1.png",
		"assets/pacman/pacman-2.png",
		"assets/pacman/pacman-3.png",
	}
	deathSrc := make([]string, 11)
	for i := 0; i < 11; i++ {
		deathSrc[i] = fmt.Sprintf("assets/pacman/death-%d.png", i+1)
	}

	aliveSprites, err := structures.InitSpriteSequence(aliveSrc)
	if err != nil {
		return nil, err
	}
	deadSprites, err := structures.InitSpriteSequence(deathSrc)
	if err != nil {
		return nil, err
	}

	am.WallSprite = wallSprite
	am.PelletSprite = pelletSprite
	am.PowerPelletSprite = powerPelletSprite
	am.BarsSprite = barsSprite
	am.PacmanSprites["alive"] = aliveSprites
	am.PacmanSprites["dead"] = deadSprites
	return am, nil
}
