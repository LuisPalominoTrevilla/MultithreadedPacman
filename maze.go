package main

import (
	"errors"
)

// Maze represents the level map/maze
type Maze struct {
	rows int
	cols int
	maze [][]interface{}
}

// AddElement to the maze
func (m *Maze) AddElement(i, j int, elem interface{}) error {
	if i >= m.rows || j >= m.cols {
		return errors.New("Invalid position to add element to maze")
	}

	m.maze[i][j] = elem
	return nil
}

// AddRow to the maze
func (m *Maze) AddRow(cols int) error {
	if m.cols > 0 && m.cols != cols {
		return errors.New("Number of columns cannot be different for each row")
	}

	m.maze = append(m.maze, make([]interface{}, cols))
	m.rows++
	m.cols = cols
	return nil
}

// InitMaze of the level with generic data
func InitMaze() *Maze {
	maze := Maze{
		rows: 0,
		cols: 0,
		maze: make([][]interface{}, 0),
	}

	return &maze
}
