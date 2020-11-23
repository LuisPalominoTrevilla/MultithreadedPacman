package models

import (
	"errors"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

// Maze represents the level map/maze
type Maze struct {
	rows     int
	cols     int
	logicMap [][]interfaces.GameObject
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

// MoveElement within the maze assuming that the move is appropriate
func (m *Maze) MoveElement(fromX, fromY int, direction constants.Direction) {
	switch obj := m.logicMap[fromY][fromX].(type) {
	case *Pacman:
		m.logicMap[fromY][fromX] = nil
		toX := mod(fromX+direction.X, m.cols)
		toY := mod(fromY+direction.Y, m.rows)
		m.logicMap[toY][toX] = obj
		obj.x = toX
		obj.y = toY
	}
}

// Draw the complete maze to the screen
func (m *Maze) Draw(screen *ebiten.Image) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if object := m.logicMap[i][j]; object != nil {
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

	m.logicMap[i][j] = elem
	return nil
}

// AddRow to the maze
func (m *Maze) AddRow(cols int) error {
	if m.cols > 0 && m.cols != cols {
		return errors.New("Number of columns cannot be different for each row")
	}

	m.logicMap = append(m.logicMap, make([]interfaces.GameObject, cols))
	m.rows++
	m.cols = cols
	return nil
}

// InitMaze of the level with generic data
func InitMaze() *Maze {
	maze := Maze{
		rows:     0,
		cols:     0,
		logicMap: make([][]interfaces.GameObject, 0),
	}

	return &maze
}
