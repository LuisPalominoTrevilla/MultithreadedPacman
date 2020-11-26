package modules

import (
	"bytes"
	"io/ioutil"

	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/constants"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

// InfiniteAudio represents an infinite audio playing on loop
type InfiniteAudio struct {
	keepPlaying   bool
	queued        bool
	nextSound     constants.SoundEffect
	currentPlayer *audio.Player
}

// Stop infinite audio
func (p *InfiniteAudio) Stop() {
	if p.currentPlayer != nil {
		p.currentPlayer.Pause()
	}
	p.keepPlaying = false
}

// Replace current audio playing on loop with a different one
func (p *InfiniteAudio) Replace(effect constants.SoundEffect, instant bool) {
	if instant && p.currentPlayer != nil {
		p.currentPlayer.Pause()
	}
	p.queued = true
	p.nextSound = effect
	p.keepPlaying = false
}

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

// PlayOnLoop the specified sound effect
func (s *SoundPlayer) PlayOnLoop(sound constants.SoundEffect) *InfiniteAudio {
	player := &InfiniteAudio{
		keepPlaying: true,
	}
	go func(effect constants.SoundEffect) {
		wait := make(chan struct{})
		for {
			if !player.keepPlaying {
				if player.queued {
					player.queued = false
					player.keepPlaying = true
					effect = player.nextSound
				} else {
					break
				}
			}
			player.currentPlayer = s.PlayOnceAndNotify(effect, wait)
			<-wait
		}
	}(sound)
	return player
}

// PlayOnceAndNotify when the sound has stopped
func (s *SoundPlayer) PlayOnceAndNotify(effect constants.SoundEffect, ready chan<- struct{}) *audio.Player {
	audioPlayer := s.playSound(effect)
	go func(player *audio.Player) {
		for player.IsPlaying() {
		}
		ready <- struct{}{}
	}(audioPlayer)
	return audioPlayer
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
