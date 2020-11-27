package modules

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/utils"
)

// CollisionDetector represents an implementation to detect collissions for movable game objects
type CollisionDetector struct {
	source interfaces.MovableGameObject
	maze   *structures.Maze
}

// ViableTiles by direction to take from the current movable object's position
func (c *CollisionDetector) ViableTiles(blockReverse bool) map[constants.Direction]*structures.Position {
	from := c.source.GetPosition()
	cols, rows := c.maze.Dimensions()
	viableTiles := make(map[constants.Direction]*structures.Position)
	for _, direction := range constants.PossibleDirections {
		if blockReverse && direction.IsOpposite(c.source.GetDirection()) {
			continue
		}
		toX := utils.Mod(from.X()+direction.X, cols)
		toY := utils.Mod(from.Y()+direction.Y, rows)
		targets := c.maze.ElementsAt(toX, toY)
		if targets == nil {
			continue
		}

		isViable := true
		for _, target := range targets {
			if target.IsUnmovable() {
				isViable = false
				break
			}
		}

		if isViable {
			viableTiles[direction] = structures.InitPosition(toX, toY)
		}
	}

	return viableTiles
}

// DetectCollision given the direction of the object
func (c *CollisionDetector) DetectCollision() []interfaces.GameObject {
	from := c.source.GetPosition()
	direction := c.source.GetDirection()
	cols, rows := c.maze.Dimensions()
	toX := utils.Mod(from.X()+direction.X, cols)
	toY := utils.Mod(from.Y()+direction.Y, rows)
	return c.maze.ElementsAt(toX, toY)
}

// InitCollisionDetector instantiates a collision detector for a source
func InitCollisionDetector(source interfaces.MovableGameObject, maze *structures.Maze) *CollisionDetector {
	detector := CollisionDetector{
		source: source,
		maze:   maze,
	}

	return &detector
}
