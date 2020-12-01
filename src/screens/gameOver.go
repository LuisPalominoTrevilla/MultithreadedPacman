package screens

import (
	"fmt"
	"image/color"
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// GameOver represents the game over screen
type GameOver struct {
	w         int
	h         int
	anchorCtx *contexts.AnchorContext
	createdAt time.Time
}
var overScreen *ebiten.Image

// Run game over screen timer before transitioning to menu
func (g *GameOver) Run() {
	for time.Now().Sub(g.createdAt).Seconds() < 4 {
	}
	g.anchorCtx.ChangeState <- constants.MenuState
}

// Draw the GameOver screen
func (g *GameOver) Draw(screen *ebiten.Image) {
	var x, y int
	var str string

	if overScreen != nil {
		op := &ebiten.DrawImageOptions{}
		_, h := overScreen.Size()
		op.GeoM.Translate(100, float64(g.h)/3-float64(h)/2)
		screen.DrawImage(overScreen, op)
	}
	str = fmt.Sprintf("Your Score: %05d", g.anchorCtx.GameScore)
	x = (g.w - len(str)*30) / 2
	y = (g.h + 120) * 2 / 3
	text.Draw(screen, str, g.anchorCtx.FontFace, x, y, color.White)
}

// NewGameOver screen
func NewGameOver(w, h int, anchorCtx *contexts.AnchorContext) *GameOver {
	overScreen, _, _ = ebitenutil.NewImageFromFile("assets/over-screen.png")
	return &GameOver{
		w:         w,
		h:         h,
		anchorCtx: anchorCtx,
		createdAt: time.Now(),
	}
}
