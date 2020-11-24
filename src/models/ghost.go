package models

import (
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Ghost represents the main enemy
type Ghost struct {
	x         int
	y         int
	speed     int
	direction constants.Direction
	sprites   map[string]*structures.SpriteSequence
	animator  *modules.Animator
}

func (g *Ghost) advanceSprites() {
	for _, seq := range g.sprites {
		seq.Advance()
	}
}

// Run the behavior of the player
func (g *Ghost) Run() {
	for {
		g.advanceSprites()
		time.Sleep(time.Duration(1000/g.speed) * time.Millisecond)
	}
}

// Draw the element to the screen in given position
func (g *Ghost) Draw(screen *ebiten.Image, x, y int) {
	g.animator.DrawFrame(screen, x, y)
}

// GetSprite of the element
func (g *Ghost) GetSprite() *ebiten.Image {
	switch g.direction {
	case constants.DirUp:
		return g.sprites["up"].GetCurrentFrame()
	case constants.DirDown:
		return g.sprites["down"].GetCurrentFrame()
	case constants.DirLeft:
		return g.sprites["left"].GetCurrentFrame()
	case constants.DirRight:
		return g.sprites["right"].GetCurrentFrame()
	default:
		return g.sprites["left"].GetCurrentFrame()
	}
}

// GetDirection of the element
func (g *Ghost) GetDirection() constants.Direction {
	return g.direction
}

// IsMatrixEditable based on the object direction
func (g *Ghost) IsMatrixEditable() bool {
	return false
}

// InitGhost enemy for the level
func InitGhost(x, y int, ghostType constants.GhostType) (*Ghost, error) {
	ghost := Ghost{
		x:         x,
		y:         y,
		speed:     constants.InitialGhostFPS,
		direction: constants.DirStatic,
		sprites:   make(map[string]*structures.SpriteSequence),
	}
	categories := []string{"left", "right", "down", "up"}
	for _, category := range categories {
		sprites := []string{
			"assets/ghost/" + string(ghostType) + "/ghost-" + category + "-1.png",
			"assets/ghost/" + string(ghostType) + "/ghost-" + category + "-2.png",
		}
		seq, err := structures.InitSpriteSequence(sprites)
		if err != nil {
			return nil, err
		}
		ghost.sprites[category] = seq
	}

	ghost.animator = modules.InitAnimator(&ghost)
	return &ghost, nil
}
