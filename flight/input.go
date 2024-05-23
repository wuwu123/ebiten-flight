package flight

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	lastIsUp bool
}

func (l *Input) Update(ship *Ship, config Config) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if !l.lastIsUp {
			ship.bullet = append(ship.bullet, &GameMatrix{X: ship.x + float64(ship.width/2) - config.GridSize/2, Y: ship.y})
		}
		l.lastIsUp = true
	} else {
		l.lastIsUp = false
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			ship.x -= config.ShipSpeedFactor
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			ship.x += config.ShipSpeedFactor
		}
		if ship.x <= ship.minx {
			ship.x = ship.minx
		}
		if ship.x >= ship.maxx {
			ship.x = ship.maxx
		}
	}

}
