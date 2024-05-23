package flight

import (
	_ "embed"
	"encoding/json"
	"image/color"
	"log"
)

//go:embed config.json
var configJson []byte

type Config struct {
	ScreenWidth     int         `json:"screenWidth"`
	ScreenHeight    int         `json:"screenHeight"`
	Title           string      `json:"title"`
	BgColor         color.NRGBA `json:"bgColor"`
	ShipSpeedFactor float64     `json:"shipSpeedFactor"`
}

func loadConfig() *Config {
	var cfg Config
	err := json.Unmarshal(configJson, &cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}
	return &cfg
}
