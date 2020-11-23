package constants

// TileSize represents the size of the side of a square tile
const (
	TileSize = 32
)

// Direction expresses a direction
type Direction struct {
	X int
	Y int
}

// DirUp direction upwards
// DirDown direction downwards
// DirLeft direction left
// DirRight direction right
var (
	DirUp     Direction = Direction{X: 0, Y: -1}
	DirDown   Direction = Direction{X: 0, Y: 1}
	DirLeft   Direction = Direction{X: -1, Y: 0}
	DirRight  Direction = Direction{X: 1, Y: 0}
	DirStatic Direction = Direction{X: 0, Y: 0}
)
