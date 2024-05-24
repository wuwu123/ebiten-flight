package flight

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const (
	//音频频率
	//转换地址 https://www.audio2edit.com/zh/convert-to-ogg
	sampleRate     = 48000
	bytesPerSample = 4 // 2 channels * 2 bytes (16 bit)

	introLengthInSecond = 50
	loopLengthInSecond  = 50
)

type Player struct {
	player       *audio.Player
	audioContext *audio.Context
}

func NewPlayer() *Player {
	return &Player{}
}

func (g *Player) Update() error {
	if g.player != nil {
		return nil
	}

	if g.audioContext == nil {
		g.audioContext = audio.NewContext(sampleRate)
	}

	oggS, err := vorbis.DecodeWithoutResampling(bytes.NewReader(MusicOgg))
	if err != nil {
		return err
	}

	s := audio.NewInfiniteLoopWithIntro(oggS, introLengthInSecond*bytesPerSample*sampleRate, loopLengthInSecond*bytesPerSample*sampleRate)

	g.player, err = g.audioContext.NewPlayer(s)
	if err != nil {
		return err
	}

	g.player.Play()
	return nil
}
