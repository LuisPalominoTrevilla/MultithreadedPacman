package screens

import (
	"image/color"
	"time"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Loading represents a loading screen
type Loading struct {
	w         int
	h         int
	anchorCtx *contexts.AnchorContext
	sprite    *structures.SpriteSequence
}

// Run loading screen logic
func (l *Loading) Run() {
	for {
		l.sprite.Advance()
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

// Draw the loading screen
func (l *Loading) Draw(screen *ebiten.Image) {
	pacmanSize := 64.0
	if l.sprite != nil {
		op := &ebiten.DrawImageOptions{}
		frame := l.sprite.GetCurrentFrame()
		width, height := frame.Size()
		op.GeoM.Scale(pacmanSize/float64(width), pacmanSize/float64(height))
		op.GeoM.Translate(float64(l.w)/2-pacmanSize/2, float64(l.h)/2.0-pacmanSize/2)
		screen.DrawImage(frame, op)
	}
	str := "LOADING..."
	x := (l.w - len(str)*30) / 2
	y := (l.h+30)/2 + int(pacmanSize)*2
	text.Draw(screen, str, l.anchorCtx.FontFace, x, y, color.White)
}

// NewLoading screen
func NewLoading(w, h int, anchorCtx *contexts.AnchorContext) *Loading {
	sprites := []string{
		"assets/pacman/pacman-1.png",
		"assets/pacman/pacman-2.png",
		"assets/pacman/pacman-3.png",
	}
	seq, _ := structures.InitSpriteSequence(sprites)
	return &Loading{
		w:         w,
		h:         h,
		anchorCtx: anchorCtx,
		sprite:    seq,
	}
}
