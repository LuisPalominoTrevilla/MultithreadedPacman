package main

import (
	"bufio"
	"fmt"
	"os"
)

// TileWidth represents the width of each tile
const TileWidth = 32

// TileHeight represents the height of each tile
const TileHeight = 32

// Level represents a level with all of its contents
type Level struct {
	maze *Maze
}

func (l *Level) parseLevel(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	l.maze = InitMaze()
	input := bufio.NewScanner(f)
	for row := 0; input.Scan(); row++ {
		line := input.Text()
		l.maze.AddRow((len(line)))
		for col, elem := range line {
			if elem == '#' {
				wall, err := InitWall()
				if err != nil {
					return err
				}
				l.maze.AddElement(row, col, wall)
			}
		}
	}

	fmt.Println(l.maze.rows, l.maze.cols)

	if err := input.Err(); err != nil {
		return err
	}

	return nil
}

func (l *Level) Draw() {
	for i := 0; i < l.maze.rows; i++ {
		for j := 0; j < l.maze.cols; j++ {
			switch l.maze.maze[i][j].(type) {
			case *Wall:
				fmt.Println("This is a wall pointer")
			default:
				fmt.Println("Could not recognize object")
			}
		}
	}
}

// InitLevel given a valid level file
func InitLevel(levelFile string) (*Level, error) {
	l := Level{}
	err := l.parseLevel(levelFile)
	return &l, err
}
