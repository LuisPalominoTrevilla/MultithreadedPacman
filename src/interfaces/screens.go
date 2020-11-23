package interfaces

import "github.com/hajimehoshi/ebiten/v2"

// Screen represents any game screen
type Screen interface {
	Run()
	Draw(screen *ebiten.Image)
}
