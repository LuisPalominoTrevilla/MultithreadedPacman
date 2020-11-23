package modules

import (
	"bytes"
	"io/ioutil"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

// SoundPlayer represents the global sound player of the app
type SoundPlayer struct {
	audioContext *audio.Context
	sounds       map[constants.SoundEffect]*structures.SoundSequence
}

func (s *SoundPlayer) playSound(effect constants.SoundEffect) *audio.Player {
	seq := s.sounds[effect]
	sound := seq.GetCurrentAudio()
	audioPlayer := audio.NewPlayerFromBytes(s.audioContext, sound)
	seq.Advance()
	audioPlayer.Play()
	return audioPlayer
}

// PlayOnceAndNotify when the sound has stopped
func (s *SoundPlayer) PlayOnceAndNotify(effect constants.SoundEffect, ready chan<- struct{}) {
	audioPlayer := s.playSound(effect)
	go func(player *audio.Player) {
		for player.IsPlaying() {
		}
		ready <- struct{}{}
	}(audioPlayer)
}

// PlayOnce the specified sound effect once
func (s *SoundPlayer) PlayOnce(effect constants.SoundEffect) {
	s.playSound(effect)
}

// InitSoundPlayer with preconfigured sounds from constants
func InitSoundPlayer() (*SoundPlayer, error) {
	soundPlayer := SoundPlayer{
		audioContext: audio.NewContext(44100),
		sounds:       make(map[constants.SoundEffect]*structures.SoundSequence),
	}
	for sound, src := range constants.AudioFiles {
		files := make([][]byte, len(src))
		for i, soundFile := range src {
			dat, err := ioutil.ReadFile(soundFile)
			if err != nil {
				return nil, err
			}
			s, err := wav.Decode(soundPlayer.audioContext, bytes.NewReader(dat))
			if err != nil {
				return nil, err
			}
			b, err := ioutil.ReadAll(s)
			if err != nil {
				return nil, err
			}
			files[i] = b
		}
		soundPlayer.sounds[sound] = structures.InitSoundSequence(files)
	}

	return &soundPlayer, nil
}