package screens

import (
	"bufio"
	"errors"
	"fmt"
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
	phase           int
	context         *contexts.GameContext
	player          *models.Pacman
	enemies         []*models.Ghost
	backgroundSound *modules.InfiniteAudio
}

func (l *Level) parseLevel(file string, numEnemies int) error {
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
				allGhosts := []constants.GhostType{
					constants.Blinky,
					constants.Pinky,
					constants.Inky,
					constants.Clyde,
				}
				for i := 0; i < numEnemies; i++ {
					ghost, err := models.InitGhost(
						col,
						row,
						float64(i)*constants.TimeBetweenSpawns+float64(i),
						allGhosts[i%len(allGhosts)],
					)
					if err != nil {
						return err
					}
					ghost.AttachCollisionDetector(modules.InitCollisionDetector(ghost, l.context.Maze))
					l.enemies = append(l.enemies, ghost)
				}
				l.context.GhostBase = structures.InitPosition(col, row)
				// Add to maze in reverse order so that red ghost will always be painted first
				for i := len(l.enemies) - 1; i >= 0; i-- {
					l.context.Maze.AddElement(row, col, l.enemies[i])
				}
			case '.', '@':
				pellet, err := models.InitPellet(elem == '@')
				if err != nil {
					return err
				}
				l.context.Maze.AddElement(row, col, pellet)
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
	sirenSounds := []constants.SoundEffect{
		constants.GhostSirenPhase1,
		constants.GhostSirenPhase2,
		constants.GhostSirenPhase3,
		constants.GhostSirenPhase4,
	}

	l.backgroundSound = l.context.SoundPlayer.PlayOnLoop(sirenSounds[l.phase])
	go l.player.Run(l.context)
	for _, enemy := range l.enemies {
		go enemy.Run(l.context)
	}
	for {
		select {
		case newPhase := <-l.context.Msg.PhaseChange:
			if newPhase > l.phase {
				l.backgroundSound.Replace(sirenSounds[newPhase%len(sirenSounds)], false)
				l.phase = newPhase
			}
		case isPowerful := <-l.context.Msg.EatPellet:
			// TODO: increment counter and check for end game
			if isPowerful {
				l.backgroundSound.Replace(constants.PowerPellet, true)
				for _, enemy := range l.enemies {
					enemy.ChangeState(constants.PowerPelletEaten)
				}
			}
		case <-l.context.Msg.PowerPelletWoreOff:
			l.backgroundSound.Replace(sirenSounds[l.phase%len(sirenSounds)], true)
		}
	}
}

// Draw the entire level
func (l *Level) Draw(screen *ebiten.Image) {
	l.context.Maze.Draw(screen)
	// TODO: Draw scoreboard and stuff on the bottom of the screen (Add more space first)
}

// InitLevel given a valid level file
func InitLevel(levelFile string, numEnemies int) (*Level, error) {
	if numEnemies <= 0 {
		return nil, errors.New("At least one enemy must be spawned")
	}
	if numEnemies > constants.MaxGhostsAllowed {
		errMsg := fmt.Sprintf("Cannot instantiate more than %d enemies", constants.MaxGhostsAllowed)
		return nil, errors.New(errMsg)
	}
	l := Level{
		phase:   0,
		enemies: make([]*models.Ghost, 0, numEnemies),
		context: &contexts.GameContext{
			Msg: &structures.MessageBroker{
				EatPellet:          make(chan bool),
				PhaseChange:        make(chan int),
				PowerPelletWoreOff: make(chan struct{}),
			},
		},
	}
	soundPlayer, err := modules.InitSoundPlayer()
	if err != nil {
		return nil, err
	}
	l.context.SoundPlayer = soundPlayer
	err = l.parseLevel(levelFile, numEnemies)
	return &l, err
}
