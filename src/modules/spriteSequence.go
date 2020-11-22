package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// SpriteSequence represents a sequence of animation frames
type SpriteSequence struct {
	current int
	frames  []*ebiten.Image
}

// Advance current frame
func (s *SpriteSequence) Advance() {
	s.current = (s.current + 1) % len(s.frames)
}

// GetCurrentFrame to be used by an animator
func (s *SpriteSequence) GetCurrentFrame() *ebiten.Image {
	return s.frames[s.current]
}

// InitSpriteSequence instantiates a sprite sequence
func InitSpriteSequence(sprites []string) (*SpriteSequence, error) {
	seq := SpriteSequence{
		current: 0,
		frames:  make([]*ebiten.Image, 0, len(sprites)),
	}

	for _, sprite := range sprites {
		img, _, err := ebitenutil.NewImageFromFile(sprite)
		if err != nil {
			return nil, err
		}
		seq.frames = append(seq.frames, img)
	}

	return &seq, nil
}
