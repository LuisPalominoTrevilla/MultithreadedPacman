package screens

import (
	"bufio"
	"os"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/models"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Level represents a level with all of its contents
type Level struct {
	context *contexts.GameContext
	player  *models.Pacman
	enemies []*models.Ghost
}

func (l *Level) parseLevel(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	l.context.Maze = structures.InitMaze()
	input := bufio.NewScanner(f)
	for row := 0; input.Scan(); row++ {
		line := input.Text()
		l.context.Maze.AddRow((len(line)))
		for col, elem := range line {
			switch elem {
			case '#':
				wall, err := models.InitWall()
				if err != nil {
					return err
				}
				l.context.Maze.AddElement(row, col, wall)
			case 'P':
				player, err := models.InitPacman(col, row)
				if err != nil {
					return err
				}
				player.AttachCollisionDetector(modules.InitCollisionDetector(player, l.context.Maze))
				l.context.MainPlayer = player
				l.player = player
				l.context.Maze.AddElement(row, col, player)
			case 'G':
				// TODO: Create specified number of enemies (randomize ghost types)
				ghost, err := models.InitGhost(col, row, 0, constants.RedGhost)
				if err != nil {
					return err
				}
				ghost.AttachCollisionDetector(modules.InitCollisionDetector(ghost, l.context.Maze))
				l.enemies = append(l.enemies, ghost)
				l.context.Maze.AddElement(row, col, ghost)
			case '.', '@':
				food, err := models.InitFood(elem == '@')
				if err != nil {
					return err
				}
				l.context.Maze.AddElement(row, col, food)
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
	return l.context.Maze.Dimensions()
}

// Run logic of the level, incluiding redraws
func (l *Level) Run() {
	// TODO: Uncomment lines to play initial sound of level
	// wait := make(chan struct{})
	// l.context.SoundPlayer.PlayOnceAndNotify(constants.GameStart, wait)
	// <-wait

	l.context.SoundPlayer.PlayOnLoop(constants.GhostSiren)
	go l.player.Run(l.context)
	for _, enemy := range l.enemies {
		go enemy.Run(l.context)
	}
	for {
		<-l.context.Msg
		// TODO: switch case to detect what message was sent
	}
}

// Draw the entire level
func (l *Level) Draw(screen *ebiten.Image) {
	l.context.Maze.Draw(screen)
	// TODO: Draw scoreboard and stuff on the bottom of the screen (Add more space first)
}

// InitLevel given a valid level file
func InitLevel(levelFile string) (*Level, error) {
	// TODO: Init ghost array with capacity to hold all specified ghosts
	l := Level{
		enemies: make([]*models.Ghost, 0),
		context: &contexts.GameContext{
			Msg: make(chan constants.EventType),
		},
	}
	soundPlayer, err := modules.InitSoundPlayer()
	if err != nil {
		return nil, err
	}
	l.context.SoundPlayer = soundPlayer
	err = l.parseLevel(levelFile)
	return &l, err
}
