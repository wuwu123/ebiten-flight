package flight

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

var (
	mplusFaceSource *text.GoTextFaceSource
	mplusNormalFace *text.GoTextFace
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
	mplusNormalFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}
}

type Text struct {
}

func (g *Text) Draw(screen *ebiten.Image, message string) {
	gray := color.RGBA{}
	const x, y = 20, 20
	w, h := text.Measure(message, mplusNormalFace, mplusNormalFace.Size*1.5)
	vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.LineSpacing = mplusNormalFace.Size * 1.5
	text.Draw(screen, message, mplusNormalFace, op)
}
