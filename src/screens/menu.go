package screens

import (
	"image/color"
	"log"
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Menu represents the menu screen
type Menu struct {
	w           int
	h           int
	font        font.Face
	keepRunning bool
}

// Run menu key listener
func (m *Menu) Run(nextScreen chan constants.GameState) {
	for m.keepRunning {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			m.keepRunning = false
			nextScreen <- constants.PlayState
		}
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
}

// Draw the menu screen
func (m *Menu) Draw(screen *ebiten.Image) {
	// TODO: All draw logic goes here
	str := "PRESS ENTER TO START"
	x := (m.w - len(str)*30) / 2
	y := (m.h + 30) / 2
	text.Draw(screen, str, m.font, x, y, color.White)
}

// NewMenu screen
func NewMenu(w, h int) *Menu {
	tt, err := truetype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	fontFace := truetype.NewFace(tt, &truetype.Options{
		Size: 30, DPI: 72, Hinting: font.HintingFull,
	})

	return &Menu{
		w:           w,
		h:           h,
		font:        fontFace,
		keepRunning: true,
	}
}
