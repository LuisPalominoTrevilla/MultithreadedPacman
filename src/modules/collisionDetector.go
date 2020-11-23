package modules

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/utils"
)

// CollisionDetector represents an implementation to detect collissions for movable game objects
type CollisionDetector struct {
	source interfaces.MovableGameObject
	maze   *structures.Maze
}

// DetectCollision given the direction of the object
func (c *CollisionDetector) DetectCollision() interfaces.GameObject {
	fromX, fromY := c.source.GetPosition()
	direction := c.source.GetDirection()
	cols, rows := c.maze.Dimensions()
	toX := utils.Mod(fromX+direction.X, cols)
	toY := utils.Mod(fromY+direction.Y, rows)
	elementsAtDestination := c.maze.ElementsAt(toX, toY)
	if elementsAtDestination == nil {
		return nil
	}

	target := elementsAtDestination.ElementOnTop()
	return target
}

// InitCollisionDetector instantiates a collision detector for a source
func InitCollisionDetector(source interfaces.MovableGameObject, maze *structures.Maze) *CollisionDetector {
	detector := CollisionDetector{
		source: source,
		maze:   maze,
	}

	return &detector
}
