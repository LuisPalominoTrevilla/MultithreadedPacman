package screens

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/models"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

// Level represents a level with all of its contents
type Level struct {
	phase           int
	ctx             *contexts.GameContext
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

	ghostBases := map[rune]constants.GhostType{
		'B': constants.Blinky,
		'P': constants.Pinky,
		'I': constants.Inky,
		'C': constants.Clyde,
	}
	l.ctx.Maze = structures.InitMaze()
	input := bufio.NewScanner(f)
	for row := 0; input.Scan(); row++ {
		line := input.Text()
		l.ctx.Maze.AddRow((len(line)))
		for col, elem := range line {
			switch elem {
			case '#', 'B', 'P', 'I', 'C':
				if ghostType, ok := ghostBases[elem]; ok {
					l.ctx.GhostBases[ghostType] = structures.InitPosition(col, row)
				}

				wall, err := models.InitWall(col, row)
				if err != nil {
					return err
				}
				l.ctx.Maze.AddElement(row, col, wall)
			case '|':
				bars, err := models.InitBars(col, row)
				if err != nil {
					return err
				}
				l.ctx.Maze.AddElement(row, col, bars)
			case 'S':
				player, err := models.InitPacman(col, row)
				if err != nil {
					return err
				}
				player.AttachCollisionDetector(modules.InitCollisionDetector(player, l.ctx.Maze))
				l.ctx.MainPlayer = player
				l.player = player
				l.ctx.Maze.AddElement(row, col, player)
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
						float64(i)*constants.TimeBetweenSpawns,
						allGhosts[i%len(allGhosts)],
					)
					if err != nil {
						return err
					}
					ghost.AttachCollisionDetector(modules.InitCollisionDetector(ghost, l.ctx.Maze))
					l.enemies = append(l.enemies, ghost)
				}
				l.ctx.GhostHome = structures.InitPosition(col, row)
				// Add to maze in reverse order so that red ghost will always be painted first
				for i := len(l.enemies) - 1; i >= 0; i-- {
					l.ctx.Maze.AddElement(row, col, l.enemies[i])
				}
			case '.', '@':
				pellet, err := models.InitPellet(col, row, elem == '@')
				if err != nil {
					return err
				}
				l.ctx.Maze.AddElement(row, col, pellet)
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
	return l.ctx.Maze.Dimensions()
}

// Run logic of the level
func (l *Level) Run(nextScreen chan constants.GameState) {
	// TODO: Uncomment lines to play initial sound of level
	// wait := make(chan struct{})
	// l.ctx.SoundPlayer.PlayOnceAndNotify(constants.GameStart, wait)
	// <-wait
	sirenSounds := []constants.SoundEffect{
		constants.GhostSirenPhase1,
		constants.GhostSirenPhase2,
		constants.GhostSirenPhase3,
		constants.GhostSirenPhase4,
	}

	l.backgroundSound = l.ctx.SoundPlayer.PlayOnLoop(sirenSounds[l.phase])
	go l.player.Run(l.ctx)
	for _, enemy := range l.enemies {
		go enemy.Run(l.ctx)
	}
	for {
		select {
		case newPhase := <-l.ctx.Msg.PhaseChange:
			if newPhase > l.phase {
				l.backgroundSound.Replace(sirenSounds[newPhase%len(sirenSounds)], false)
				l.phase = newPhase
			}
		case isPowerful := <-l.ctx.Msg.EatPellet:
			// TODO: increment counter and check for end game
			if isPowerful {
				l.backgroundSound.Replace(constants.PowerPellet, true)
				for _, enemy := range l.enemies {
					enemy.ChangeState(constants.PowerPelletEaten)
				}
			}
		case <-l.ctx.Msg.PowerPelletWoreOff:
			l.backgroundSound.Replace(sirenSounds[l.phase%len(sirenSounds)], true)
		}
	}
}

// Draw the entire level
func (l *Level) Draw(screen *ebiten.Image) {
	l.ctx.Maze.Draw(screen)
	// TODO: Draw scoreboard and stuff on the bottom of the screen (Add more space first)
}

// NewLevel given a valid level file
func NewLevel(levelFile string, numEnemies int) (*Level, error) {
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
		ctx: &contexts.GameContext{
			GhostBases: make(map[constants.GhostType]interfaces.Location),
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
	l.ctx.SoundPlayer = soundPlayer
	err = l.parseLevel(levelFile, numEnemies)
	return &l, err
}
