package models

import (
	"errors"
)

// Maze represents the level map/maze
type Maze struct {
	Rows int
	Cols int
	Maze [][]interface{}
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

// MoveElement within the maze. This method does not verify rules
func (m *Maze) MoveElement(fromX, fromY, dx, dy int) {
	switch obj := m.Maze[fromY][fromX].(type) {
	case *Pacman:
		m.Maze[fromY][fromX] = nil
		toX := mod(fromX+dx, m.Cols)
		toY := mod(fromY+dy, m.Rows)
		m.Maze[toY][toX] = obj
		obj.x = toX
		obj.y = toY
	}
}

// AddElement to the maze
func (m *Maze) AddElement(i, j int, elem interface{}) error {
	if i >= m.Rows || j >= m.Cols {
		return errors.New("Invalid position to add element to maze")
	}

	m.Maze[i][j] = elem
	return nil
}

// AddRow to the maze
func (m *Maze) AddRow(cols int) error {
	if m.Cols > 0 && m.Cols != cols {
		return errors.New("Number of columns cannot be different for each row")
	}

	m.Maze = append(m.Maze, make([]interface{}, cols))
	m.Rows++
	m.Cols = cols
	return nil
}

// InitMaze of the level with generic data
func InitMaze() *Maze {
	maze := Maze{
		Rows: 0,
		Cols: 0,
		Maze: make([][]interface{}, 0),
	}

	return &maze
}
