package structures

import (
	"sort"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
)

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

// RemoveElement from the group
func (g *GameObjectGroup) RemoveElement(target interfaces.GameObject) bool {
	index := g.findIndex(target)
	if index < 0 || index >= len(g.elements) {
		return false
	}
	g.elements = append(g.elements[:index], g.elements[index+1:]...)
	return true
}

// ElementOnTop of the group object. Nil if empty
func (g *GameObjectGroup) ElementOnTop() interfaces.GameObject {
	if len(g.elements) == 0 {
		return nil
	}
	return g.elements[len(g.elements)-1]
}

// GetObjects that belong to the group
func (g *GameObjectGroup) GetObjects() []interfaces.GameObject {
	objs := make([]interfaces.GameObject, len(g.elements))
	copy(objs, g.elements)
	sort.SliceStable(objs, func(i, j int) bool {
		return objs[i].GetLayerIndex() > objs[j].GetLayerIndex()
	})
	return objs
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
