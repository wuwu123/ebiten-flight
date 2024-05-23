package flight

import (
	"crypto/rand"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/big"
)

type Bullet struct {
	List         []*GameMatrix
	lastNum      int
	downNum      int
	bottomHeight int
}

func NewBullet(bottomHeight int) *Bullet {
	return &Bullet{bottomHeight: bottomHeight}
}

func (l *Bullet) Update(config Config) {
	l.lastNum += 1
	if l.lastNum%15 == 0 {
		l.lastNum = 0
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(config.ScreenWidth-int(config.GridSize))))
		var newY = float64(n.Int64()) + config.GridSize/2
		l.List = append(l.List, &GameMatrix{
			X: newY,
			Y: 0,
		})
	}

}

func (l *Bullet) Draw(screen *ebiten.Image, config Config) {
	var listLen = len(l.List)
	for i, v := range l.List {
		v.Y += 1
		if v.Y >= float64(config.ScreenHeight-l.bottomHeight-int(config.GridSize)) {
			if i+1 >= listLen {
				l.List = append(l.List[:i])
			} else {
				l.List = append(l.List[:i], l.List[i+1:]...)
			}
			continue
		}
		vector.DrawFilledRect(screen, float32(v.X), float32(v.Y), float32(config.GridSize), float32(config.GridSize), color.RGBA{0x80, 0xa0, 0xc0, 0xff}, false)
	}
}
