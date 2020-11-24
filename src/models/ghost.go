package models

import (
	"log"
	"math/rand"
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Ghost represents the main enemy
type Ghost struct {
	x                 int
	y                 int
	speed             int
	direction         constants.Direction
	sprites           map[string]*structures.SpriteSequence
	animator          *modules.Animator
	collisionDetector *modules.CollisionDetector
}

func pickRandomDirection() constants.Direction {
	allDirections := constants.PossibleDirections
	rand.Seed(time.Now().UnixNano())
	return allDirections[rand.Intn(len(allDirections))]
}

func (g *Ghost) advanceSprites() {
	for _, seq := range g.sprites {
		seq.Advance()
	}
}

func (g *Ghost) attemptChangeDirection() {
	viableTiles := g.collisionDetector.ViableTiles()
	options := len(viableTiles)
	if options == 0 {
		return
	}
	tmp := make([]constants.Direction, 0)
	for direction := range viableTiles {
		if options == 1 && direction == g.direction {
			return
		}
		tmp = append(tmp, direction)
	}
	rand.Seed(time.Now().UnixNano())
	selected := tmp[rand.Intn(len(tmp))]
	g.direction = selected
}

// Run the behavior of the player
func (g *Ghost) Run(gameContext *contexts.GameContext) {
	if g.collisionDetector == nil {
		log.Fatal("Collision detector is not attached")
	}

	prevDirection := g.direction
	recentlyChangedDirection := false
	for {
		gameContext.MazeMutex.Lock()
		if !recentlyChangedDirection {
			g.attemptChangeDirection()
		}
		recentlyChangedDirection = g.direction != prevDirection
		if !recentlyChangedDirection {
			target := g.collisionDetector.DetectCollision()
			switch target.(type) {
			case *Wall:
				g.direction = pickRandomDirection()
			default:
				gameContext.Maze.MoveElement(g, false)
				g.advanceSprites()
			}
		}
		prevDirection = g.direction
		gameContext.MazeMutex.Unlock()
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

// IsUnmovable by any force
func (g *Ghost) IsUnmovable() bool {
	return false
}

// GetPosition of the element
func (g *Ghost) GetPosition() (x, y int) {
	return g.x, g.y
}

// SetPosition of the element
func (g *Ghost) SetPosition(x, y int) {
	g.x = x
	g.y = y
}

// AttachCollisionDetector to the element
func (g *Ghost) AttachCollisionDetector(collisionDetector *modules.CollisionDetector) {
	g.collisionDetector = collisionDetector
}

// InitGhost enemy for the level
func InitGhost(x, y int, ghostType constants.GhostType) (*Ghost, error) {
	ghost := Ghost{
		x:       x,
		y:       y,
		speed:   constants.InitialGhostFPS,
		sprites: make(map[string]*structures.SpriteSequence),
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

	ghost.direction = pickRandomDirection()
	ghost.animator = modules.InitAnimator(&ghost)
	return &ghost, nil
}
