package structures

import (
	"errors"
	"log"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/interfaces"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// Maze represents the level map/maze
type Maze struct {
	rows     int
	cols     int
	logicMap [][]*GameObjectGroup
}

// Dimensions of the logic map
func (m *Maze) Dimensions() (width, height int) {
	return m.cols, m.rows
}

// MoveElement within the maze without checking whether the move is appropriate
func (m *Maze) MoveElement(elem interfaces.MovableGameObject, delDestElement bool) {
	from := elem.GetPosition()
	sourceGroup := m.logicMap[from.Y()][from.X()]

	valid := sourceGroup.RemoveElement(elem)
	if !valid {
		log.Println("Could not find element to move")
		return
	}

	direction := elem.GetDirection()
	toX := utils.Mod(from.X()+direction.X, m.cols)
	toY := utils.Mod(from.Y()+direction.Y, m.rows)
	destinationGroup := m.logicMap[toY][toX]
	if delDestElement {
		// TODO: Dispose of the sprite of the removed element
		destinationGroup.RemoveTopElement()
	}
	destinationGroup.AddElement(elem)
	elem.SetPosition(toX, toY)
}

// ElementsAt the specified position. Nil if position is out of bounds
func (m *Maze) ElementsAt(x, y int) *GameObjectGroup {
	if x < 0 || x >= m.cols || y < 0 || y >= m.rows {
		return nil
	}

	return m.logicMap[y][x]
}

// Draw the complete maze to the screen
func (m *Maze) Draw(screen *ebiten.Image) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			groupObject := m.logicMap[i][j]
			for _, object := range groupObject.Elements() {
				object.Draw(screen, j, i)
			}
		}
	}
}

// AddElement to the maze
func (m *Maze) AddElement(i, j int, elem interfaces.GameObject) error {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return errors.New("Invalid position to add element to maze")
	}

	groupObject := m.logicMap[i][j]
	groupObject.AddElement(elem)
	return nil
}

// AddRow to the maze
func (m *Maze) AddRow(cols int) error {
	if m.cols > 0 && m.cols != cols {
		return errors.New("Number of columns cannot be different for each row")
	}

	m.logicMap = append(m.logicMap, make([]*GameObjectGroup, cols))
	for j := 0; j < cols; j++ {
		m.logicMap[m.rows][j] = InitGameObjectGroup()
	}
	m.rows++
	m.cols = cols
	return nil
}

// InitMaze of the level with generic data
func InitMaze() *Maze {
	maze := Maze{
		rows:     0,
		cols:     0,
		logicMap: make([][]*GameObjectGroup, 0),
	}

	return &maze
}
