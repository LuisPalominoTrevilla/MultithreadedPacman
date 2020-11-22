package utils

// Direction expresses a direction
type Direction int

// DirUp direction upwards
// DirDown direction downwards
// DirLeft direction left
// DirRight direction right
const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
	DirStatic
)
