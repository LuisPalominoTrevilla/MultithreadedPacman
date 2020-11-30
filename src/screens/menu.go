package screens

import (
	"image/color"
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/modules"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Menu represents the menu screen
type Menu struct {
	w           int
	h           int
	anchorCtx   *contexts.AnchorContext
	keepRunning bool
	mainTheme   *modules.InfiniteAudio
}

var menuScreen *ebiten.Image

// Run menu key listener
func (m *Menu) Run() {
	for m.keepRunning {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			m.keepRunning = false
			m.anchorCtx.ChangeState <- constants.PlayState
		}
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
	m.mainTheme.Stop()
}

// Draw the menu screen
func (m *Menu) Draw(screen *ebiten.Image) {
	if menuScreen != nil {
		op := &ebiten.DrawImageOptions{}
		_, h := menuScreen.Size()
		op.GeoM.Translate(0, float64(m.h)/3-float64(h)/2)
		screen.DrawImage(menuScreen, op)
	}
	str := "PRESS ENTER TO START"
	x := (m.w - len(str)*30) / 2
	y := (m.h+30)/2 + 100
	text.Draw(screen, str, m.anchorCtx.FontFace, x, y, color.White)
}

// NewMenu screen
func NewMenu(w, h int, anchorCtx *contexts.AnchorContext) *Menu {
	menuScreen, _, _ = ebitenutil.NewImageFromFile("assets/menu-screen.jpg")
	return &Menu{
		w:           w,
		h:           h,
		anchorCtx:   anchorCtx,
		keepRunning: true,
		mainTheme:   anchorCtx.SoundPlayer.PlayOnLoop(constants.MainTheme),
	}
}
