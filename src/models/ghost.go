package models

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Ghost represents the main enemy
type Ghost struct {
	state             interfaces.GhostState
	position          interfaces.Location
	speed             int
	idleStateTime     float64
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

func (g *Ghost) orientedSprite() *ebiten.Image {
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

// ChangeState given an event
func (g *Ghost) ChangeState(event constants.EventType) {
	newState := g.state.ApplyTransition(event)
	if newState != nil {
		g.state = newState
	}
}

func (g *Ghost) attemptChangeDirection(target interfaces.MovableGameObject) {
	viableTiles := g.collisionDetector.ViableTiles()
	options := len(viableTiles)
	if options == 0 {
		return
	}
	var selected constants.Direction
	shortestDistance := math.MaxFloat64
	directions := make([]constants.Direction, 0, len(viableTiles))
	for direction, position := range viableTiles {
		if options == 1 && direction == g.direction {
			return
		}
		if target == nil {
			directions = append(directions, direction)
			continue
		}
		currentDistance := position.DistanceTo(target.GetPosition())
		if currentDistance < shortestDistance {
			shortestDistance = currentDistance
			selected = direction
		}
	}

	if target == nil {
		rand.Seed(time.Now().UnixNano())
		selected = directions[rand.Intn(len(directions))]
	}
	g.direction = selected
}

// Run the behavior of the ghost
func (g *Ghost) Run(gameContext *contexts.GameContext) {
	if g.collisionDetector == nil {
		log.Fatal("Collision detector is not attached")
	}

	g.state = InitIdle(g, gameContext)
	for {
		g.state.Run()
		time.Sleep(time.Duration(1000/g.speed) * time.Millisecond)
	}
}

// Draw the element to the screen in given position
func (g *Ghost) Draw(screen *ebiten.Image, x, y int) {
	g.animator.DrawFrame(screen, x, y)
}

// GetSprite of the element
func (g *Ghost) GetSprite() *ebiten.Image {
	if g.state == nil {
		return g.orientedSprite()
	}
	return g.state.GetSprite()
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
func (g *Ghost) GetPosition() interfaces.Location {
	return g.position
}

// SetPosition of the element
func (g *Ghost) SetPosition(x, y int) {
	g.position.SetX(x)
	g.position.SetY(y)
}

// AttachCollisionDetector to the element
func (g *Ghost) AttachCollisionDetector(collisionDetector *modules.CollisionDetector) {
	g.collisionDetector = collisionDetector
}

// InitGhost enemy for the level
func InitGhost(x, y int, idleStateTime float64, ghostType constants.GhostType) (*Ghost, error) {
	ghost := Ghost{
		position:      structures.InitPosition(x, y),
		idleStateTime: idleStateTime,
		speed:         constants.DefaultGhostFPS,
		sprites:       make(map[string]*structures.SpriteSequence),
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
