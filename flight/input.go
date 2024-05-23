package flight

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct{}

func (i *Input) Update(ship *Ship, config *Config) {
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
