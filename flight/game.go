package flight

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path"
	"path/filepath"
)

type GameMatrix struct {
	X float64
	Y float64
}
type Game struct {
	input  *Input
	cfg    Config
	ship   *Ship
	player *Player
}

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	return &Game{
		input:  &Input{},
		cfg:    cfg,
		ship:   NewShip(cfg),
		player: NewPlayer(),
	}
}

// 每帧（frame）调用。帧是渲染使用的一个时间单位，依赖显示器的刷新率。如果显示器的刷新率为60Hz，Draw将会每秒被调用60次
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen, g.cfg)
}

// 该方法接收游戏窗口的尺寸作为参数，返回游戏的逻辑屏幕大小
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}

// 每个tick都会被调用。tick是引擎更新的一个时间单位，默认为1/60s。tick的倒数我们一般称为帧，即游戏的更新频率。
func (g *Game) Update() error {
	g.input.Update(g.ship, g.cfg)
	return g.player.Update()
}

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	x      float64
	y      float64
	minx   float64
	maxx   float64
	bullet []*GameMatrix
	config Config
}

func NewShip(config Config) *Ship {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	staticDir := filepath.Join(cwd, "flight", "images")
	img, _, err := ebitenutil.NewImageFromFile(path.Join(staticDir, "780.jpeg"))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	screenWidth, screenHeight := config.ScreenWidth, config.ScreenHeight
	ship := &Ship{
		config: config,
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
		minx:   0,
		maxx:   float64(screenWidth - width),
	}
	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg Config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, op)
	for i, v := range ship.bullet {
		v.Y -= ship.config.GridSize
		if v.Y <= 0 {
			ship.bullet = ship.bullet[:i]
			continue
		}
		vector.DrawFilledRect(screen, float32(v.X), float32(v.Y), float32(ship.config.GridSize), float32(ship.config.GridSize), color.RGBA{0x80, 0xa0, 0xc0, 0xff}, false)
	}
}
