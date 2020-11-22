package main

import (
	"errors"
	"fmt"
)

// Maze represents the level map/maze
type Maze struct {
	rows int
	cols int
	maze [][]interface{}
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
	switch obj := m.maze[fromY][fromX].(type) {
	case *Pacman:
		m.maze[fromY][fromX] = nil
		toX := mod(fromX+dx, m.cols)
		toY := mod(fromY+dy, m.rows)
		m.maze[toY][toX] = obj
		obj.x = toX
		obj.y = toY
	default:
		fmt.Println("Not pacman")
	}
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
