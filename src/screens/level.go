package screens

import (
	"bufio"
	"math"
	"os"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/models"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/hajimehoshi/ebiten/v2"
)

// TileWidth represents the width of each tile
const TileWidth = 32

// TileHeight represents the height of each tile
const TileHeight = 32

// Level represents a level with all of its contents
type Level struct {
	Maze   *models.Maze
	Player *models.Pacman
}

func (l *Level) parseLevel(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	l.Maze = models.InitMaze()
	input := bufio.NewScanner(f)
	for row := 0; input.Scan(); row++ {
		line := input.Text()
		l.Maze.AddRow((len(line)))
		for col, elem := range line {
			switch elem {
			case '#':
				wall, err := models.InitWall()
				if err != nil {
					return err
				}
				l.Maze.AddElement(row, col, wall)
			case 'P':
				player, err := models.InitPacman(col, row)
				if err != nil {
					return err
				}
				l.Player = player
				l.Maze.AddElement(row, col, player)
			case '.', '@':
				food, err := models.InitFood(elem == '@')
				if err != nil {
					return err
				}
				l.Maze.AddElement(row, col, food)
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
	for i := 0; i < l.Maze.Rows; i++ {
		for j := 0; j < l.Maze.Cols; j++ {
			switch obj := l.Maze.Maze[i][j].(type) {
			case *models.Wall:
				op := &ebiten.DrawImageOptions{}
				w, h := obj.Sprite.Size()
				op.GeoM.Scale(TileWidth/float64(w), TileHeight/float64(h))
				op.GeoM.Translate(TileWidth*float64(j), TileHeight*float64(i))
				screen.DrawImage(obj.Sprite, op)
			case *models.Pacman:
				op := &ebiten.DrawImageOptions{}
				frame := obj.Sprites.GetCurrentFrame()
				w, h := frame.Size()
				op.GeoM.Scale(TileWidth/float64(w), TileHeight/float64(h))
				switch obj.Direction {
				case modules.DirUp:
					op.GeoM.Translate(-TileWidth/2, -TileHeight/2)
					op.GeoM.Rotate(3 * math.Pi / 2)
					op.GeoM.Translate(TileWidth/2, TileHeight/2)
				case modules.DirDown:
					op.GeoM.Translate(-TileWidth/2, -TileHeight/2)
					op.GeoM.Rotate(math.Pi / 2)
					op.GeoM.Translate(TileWidth/2, TileHeight/2)
				case modules.DirLeft:
					op.GeoM.Translate(-TileWidth/2, -TileHeight/2)
					op.GeoM.Scale(-1, 1)
					op.GeoM.Translate(TileWidth/2, TileHeight/2)
				}
				op.GeoM.Translate(TileWidth*float64(j), TileHeight*float64(i))
				screen.DrawImage(frame, op)
			case *models.Food:
				op := &ebiten.DrawImageOptions{}
				w, h := obj.Sprite.Size()
				op.GeoM.Scale(TileWidth/float64(w), TileHeight/float64(h))
				op.GeoM.Translate(TileWidth*float64(j), TileHeight*float64(i))
				screen.DrawImage(obj.Sprite, op)
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
