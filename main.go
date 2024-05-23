package main

import (
	"game/flight"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game := flight.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
