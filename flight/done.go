package flight

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Done struct {
	overImg *ebiten.Image
	config  Config
	width   int
	height  int
}

func NewDone(config Config) *Done {
	//cwd, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	//staticDir := filepath.Join(cwd, "flight", "images")
	//overImg, _, err := ebitenutil.NewImageFromFile(path.Join(staticDir, "over.jpeg"))
	overImg, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(OverImgData))
	if err != nil {
		log.Fatal(err)
	}
	width, height := overImg.Bounds().Dx(), overImg.Bounds().Dy()
	done := &Done{
		config:  config,
		overImg: overImg,
		width:   width,
		height:  height,
	}
	return done
}

// DrawOver 缩放图片填充
func (done *Done) DrawOver(screen *ebiten.Image) {
	scaleX := float64(done.config.ScreenWidth) / float64(done.width)
	scaleY := float64(done.config.ScreenHeight) / float64(done.height)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(done.overImg, opts)
}
