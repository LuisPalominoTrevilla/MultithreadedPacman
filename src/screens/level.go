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
	maze        *structures.Maze
	soundPlayer *modules.SoundPlayer
	player      *models.Pacman
	enemies     []*models.Ghost
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
				player.AttachSoundPlayer(l.soundPlayer)
				l.player = player
				l.maze.AddElement(row, col, player)
			case 'G':
				// TODO: Create specified number of enemies (randomize ghost types)
				ghost, err := models.InitGhost(col, row, constants.RedGhost)
				if err != nil {
					return err
				}
				l.enemies = append(l.enemies, ghost)
				l.maze.AddElement(row, col, ghost)
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
	// TODO: Uncomment lines to play initial sound of level
	// wait := make(chan struct{})
	// l.soundPlayer.PlayOnceAndNotify(constants.GameStart, wait)
	// <-wait
	// close(wait)

	l.soundPlayer.PlayOnLoop(constants.GhostSiren)
	levelMsg := make(chan constants.EventType)
	go l.player.Run(l.maze, levelMsg)
	for _, enemy := range l.enemies {
		go enemy.Run()
	}
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
	// TODO: Init ghost array with capacity to hold all specified ghosts
	l := Level{
		enemies: make([]*models.Ghost, 0),
	}
	soundPlayer, err := modules.InitSoundPlayer()
	if err != nil {
		return nil, err
	}
	l.soundPlayer = soundPlayer
	err = l.parseLevel(levelFile)
	return &l, err
}
