package models

import "github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"

// GameObjectGroup represents a group of objects that are coexisting in the same tile
type GameObjectGroup struct {
	elements []interfaces.GameObject
}

func (g *GameObjectGroup) findIndex(target interfaces.GameObject) int {
	for i, elem := range g.elements {
		if elem == target {
			return i
		}
	}

	return -1
}

func (g *GameObjectGroup) removeByIndex(index int) interfaces.GameObject {
	if index < 0 || index >= len(g.elements) {
		return nil
	}
	removed := g.elements[index]
	g.elements = append(g.elements[:index], g.elements[index+1:]...)
	return removed
}

// RemoveElement from the group
func (g *GameObjectGroup) RemoveElement(target interfaces.GameObject) bool {
	index := g.findIndex(target)
	return g.removeByIndex(index) != nil
}

// RemoveTopElement from the group
func (g *GameObjectGroup) RemoveTopElement() interfaces.GameObject {
	return g.removeByIndex(len(g.elements) - 1)
}

// ElementOnTop of the group object. Nil if empty
func (g *GameObjectGroup) ElementOnTop() interfaces.GameObject {
	if len(g.elements) == 0 {
		return nil
	}
	return g.elements[len(g.elements)-1]
}

// AddElement on top of the group object
func (g *GameObjectGroup) AddElement(object interfaces.GameObject) {
	g.elements = append(g.elements, object)
}

// InitGameObjectGroup from a tile
func InitGameObjectGroup() *GameObjectGroup {
	objectGroup := GameObjectGroup{
		elements: []interfaces.GameObject{},
	}
	return &objectGroup
}
