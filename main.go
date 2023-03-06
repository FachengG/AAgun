package main

import (
	"fmt"
	_ "image/png"
	"log"

	env "github.com/FachengG/AAgun/pkgs/envirment"
	"github.com/FachengG/AAgun/pkgs/obj"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("imgs/bullet.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	frame   float64
	wind    env.Wind
	gun     obj.Gun
	bullets []obj.Bullet
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		fmt.Print("shot")
		if g.gun.Shoot() {
			g.bullets = append(g.bullets, obj.Bullet{}.New())
		}
	}
	for _, b := range g.bullets {
		b.UpdateSpeed(g.wind, g.frame)
		b.Movement(g.frame)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range g.bullets {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.Position.X, b.Position.Y)
		screen.DrawImage(img, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}

func main() {
	g := &Game{frame: float64(1) / 60,
		wind: env.Wind{X: 0.01, Y: 0.01},
		gun:  obj.Gun{BulletsNum: 99}}

	ebiten.SetWindowTitle("bullet")
	fmt.Print("start game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
