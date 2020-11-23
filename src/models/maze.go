package models

import (
	"errors"
	"log"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

// Maze represents the level map/maze
type Maze struct {
	rows     int
	cols     int
	logicMap [][]*GameObjectGroup
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

// Dimensions of the logic map
func (m *Maze) Dimensions() (width, height int) {
	return m.cols, m.rows
}

// MoveElement within the maze without checking whether the move is appropriate
func (m *Maze) MoveElement(elem interfaces.MovableGameObject, delDestElement bool) {
	fromX, fromY := elem.GetPosition()
	sourceGroup := m.logicMap[fromY][fromX]

	switch elem.(type) {
	case *Pacman:
		valid := sourceGroup.RemoveElement(elem)
		if !valid {
			log.Println("Could not find element to move")
			return
		}

		direction := elem.GetDirection()
		toX := mod(fromX+direction.X, m.cols)
		toY := mod(fromY+direction.Y, m.rows)
		destinationGroup := m.logicMap[toY][toX]
		if delDestElement {
			// TODO: Dispose of the sprite of the removed element
			destinationGroup.RemoveTopElement()
		}
		destinationGroup.AddElement(elem)
		elem.SetPosition(toX, toY)
	}
}

// Draw the complete maze to the screen
func (m *Maze) Draw(screen *ebiten.Image) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			groupObject := m.logicMap[i][j]
			if object := groupObject.ElementOnTop(); object != nil {
				object.Draw(screen, j, i)
			}
		}
	}
}

// AddElement to the maze
func (m *Maze) AddElement(i, j int, elem interfaces.GameObject) error {
	if i >= m.rows || j >= m.cols {
		return errors.New("Invalid position to add element to maze")
	}

	groupObject := m.logicMap[i][j]
	groupObject.AddElement(elem)
	return nil
}

// AddRow to the maze
func (m *Maze) AddRow(cols int) error {
	if m.cols > 0 && m.cols != cols {
		return errors.New("Number of columns cannot be different for each row")
	}

	m.logicMap = append(m.logicMap, make([]*GameObjectGroup, cols))
	for j := 0; j < cols; j++ {
		m.logicMap[m.rows][j] = InitGameObjectGroup()
	}
	m.rows++
	m.cols = cols
	return nil
}

// InitMaze of the level with generic data
func InitMaze() *Maze {
	maze := Maze{
		rows:     0,
		cols:     0,
		logicMap: make([][]*GameObjectGroup, 0),
	}

	return &maze
}
