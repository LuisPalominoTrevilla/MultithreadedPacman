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
	isAlive           bool
	state             interfaces.GhostState
	kind              constants.GhostType
	layerIndex        int
	phase             int
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

func (g *Ghost) turnTowards(target interfaces.Location, runAway, blockReverse bool) {
	viableTiles := g.collisionDetector.ViableTiles(blockReverse)
	options := len(viableTiles)
	if options == 0 {
		return
	}
	var selected constants.Direction
	var bestDistance float64
	if runAway {
		bestDistance = math.MinInt64
	} else {
		bestDistance = math.MaxInt64
	}

	directions := make([]constants.Direction, 0, len(viableTiles))
	for direction, position := range viableTiles {
		if options == 1 && direction == g.direction {
			return
		}
		if target == nil {
			directions = append(directions, direction)
			continue
		}
		currentDistance := position.DistanceTo(target)
		var meetsDistanceCriteria bool
		if runAway {
			meetsDistanceCriteria = currentDistance > bestDistance
		} else {
			meetsDistanceCriteria = currentDistance < bestDistance
		}

		if meetsDistanceCriteria {
			bestDistance = currentDistance
			selected = direction
		}
	}

	if target == nil {
		rand.Seed(time.Now().UnixNano())
		selected = directions[rand.Intn(len(directions))]
	}
	g.direction = selected
}

// ChangeState given an event
func (g *Ghost) ChangeState(event constants.StateEvent) {
	newState := g.state.ApplyTransition(event)
	if newState != nil {
		g.state = newState
	}
}

// AttemptEatPacman given the current ghost state
func (g *Ghost) AttemptEatPacman(obj interfaces.MovableGameObject) bool {
	return g.state.AttemptEatPacman(obj)
}

// Run the behavior of the ghost
func (g *Ghost) Run(ctx *contexts.GameContext) {
	if g.collisionDetector == nil {
		log.Fatal("Collision detector is not attached")
	}

	g.state = InitIdle(g, ctx)
	for g.isAlive {
		ctx.MazeMutex.Lock()
		g.state.Run()
		ctx.MazeMutex.Unlock()
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

// CanGhostsGoThrough by any force
func (g *Ghost) CanGhostsGoThrough() bool {
	return true
}

// GetLayerIndex of the element
func (g *Ghost) GetLayerIndex() int {
	return g.layerIndex
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
		isAlive:       true,
		layerIndex:    constants.GhostLayerIdx,
		kind:          ghostType,
		phase:         0,
		position:      structures.InitPosition(x, y),
		idleStateTime: idleStateTime,
		speed:         constants.DefaultGhostFPS,
		sprites:       make(map[string]*structures.SpriteSequence),
	}

	categories := []string{"left", "right", "down", "up", "panic", "flicker"}
	for _, category := range categories {
		gType := string(ghostType)
		if category == "panic" || category == "flicker" {
			gType = ""
		}
		sprites := []string{
			"assets/ghost/" + gType + "/ghost-" + category + "-1.png",
			"assets/ghost/" + gType + "/ghost-" + category + "-2.png",
		}
		seq, err := structures.InitSpriteSequence(sprites)
		if err != nil {
			return nil, err
		}
		ghost.sprites[category] = seq
	}

	categories = []string{"eaten-left", "eaten-right", "eaten-down", "eaten-up"}
	for _, category := range categories {
		sprites := []string{
			"assets/ghost/ghost-" + category + ".png",
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
