package flight

import (
	_ "embed"
	"encoding/json"
	"image/color"
	"log"
)

//go:embed config.json
var configJson []byte

//go:embed images/780.jpeg
var ShipImp []byte

//go:embed images/over.jpeg
var OverImgData []byte

//go:embed music/CMJ.ogg
var MusicOgg []byte

type Config struct {
	ScreenWidth     int         `json:"screenWidth"`  //页面的宽度
	ScreenHeight    int         `json:"screenHeight"` //页面的高度
	Title           string      `json:"title"`
	BgColor         color.NRGBA `json:"bgColor"`
	ShipSpeedFactor float64     `json:"shipSpeedFactor"` //一次性移动的距离
	GridSize        float64     `json:"gridSize"`        //炮弹的大小
	ShipAutoBullet  bool        `json:"shipAutoBullet"`  // 是否自动发射子弹
}

func loadConfig() Config {
	var cfg Config
	err := json.Unmarshal(configJson, &cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}
	return cfg
}
