package screens

import (
	"bufio"
	"os"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/models"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Level represents a level with all of its contents
type Level struct {
	maze   *structures.Maze
	player *models.Pacman
}

func (l *Level) parseLevel(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	l.maze = structures.InitMaze()
	input := bufio.NewScanner(f)
	for row := 0; input.Scan(); row++ {
		line := input.Text()
		l.maze.AddRow((len(line)))
		for col, elem := range line {
			switch elem {
			case '#':
				wall, err := models.InitWall()
				if err != nil {
					return err
				}
				l.maze.AddElement(row, col, wall)
			case 'P':
				player, err := models.InitPacman(col, row)
				if err != nil {
					return err
				}
				player.AttachCollisionDetector(modules.InitCollisionDetector(player, l.maze))
				l.player = player
				l.maze.AddElement(row, col, player)
			case '.', '@':
				food, err := models.InitFood(elem == '@')
				if err != nil {
					return err
				}
				l.maze.AddElement(row, col, food)
			}
		}
	}

	if err := input.Err(); err != nil {
		return err
	}

	return nil
}

// Size of the level
func (l *Level) Size() (width, height int) {
	return l.maze.Dimensions()
}

// Run logic of the level, incluiding redraws
func (l *Level) Run() {
	// TODO: sleep while playing initial sound of level
	// time.Sleep(time.Duration(2) * time.Second)
	levelMsg := make(chan constants.EventType)
	go l.player.Run(l.maze, levelMsg)
	for {
		<-levelMsg
		// TODO: switch case to detect what message was sent
	}
}

// Draw the entire level
func (l *Level) Draw(screen *ebiten.Image) {
	l.maze.Draw(screen)
	// TODO: Draw scoreboard and stuff on the bottom of the screen (Add more space first)
}

// InitLevel given a valid level file
func InitLevel(levelFile string) (*Level, error) {
	l := Level{}
	err := l.parseLevel(levelFile)
	return &l, err
}
