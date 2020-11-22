package main

import (
	"bufio"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Direction expresses a direction
type Direction int

// DirUp direction upwards
// DirDown direction downwards
// DirLeft direction left
// DirRight direction right
const (
	DirUp    Direction = iota
	DirDown  Direction = iota
	DirLeft  Direction = iota
	DirRight Direction = iota
)

// TileWidth represents the width of each tile
const TileWidth = 32

// TileHeight represents the height of each tile
const TileHeight = 32

// Level represents a level with all of its contents
type Level struct {
	maze   *Maze
	player *Pacman
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
			switch elem {
			case '#':
				wall, err := InitWall()
				if err != nil {
					return err
				}
				l.maze.AddElement(row, col, wall)
			case 'P':
				player, err := InitPacman(col, row)
				if err != nil {
					return err
				}
				l.player = player
				l.maze.AddElement(row, col, player)
			}
		}
	}

	if err := input.Err(); err != nil {
		return err
	}

	return nil
}

// Draw the entire level
func (l *Level) Draw(screen *ebiten.Image) {
	for i := 0; i < l.maze.rows; i++ {
		for j := 0; j < l.maze.cols; j++ {
			switch obj := l.maze.maze[i][j].(type) {
			case *Wall:
				op := &ebiten.DrawImageOptions{}
				w, h := obj.sprite.Size()
				op.GeoM.Scale(TileWidth/float64(w), TileHeight/float64(h))
				op.GeoM.Translate(TileWidth*float64(j), TileHeight*float64(i))
				screen.DrawImage(obj.sprite, op)
			case *Pacman:
				op := &ebiten.DrawImageOptions{}
				w, h := obj.sprite.Size()
				op.GeoM.Scale(TileWidth/float64(w), TileHeight/float64(h))
				op.GeoM.Translate(TileWidth*float64(j), TileHeight*float64(i))
				screen.DrawImage(obj.sprite, op)
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
