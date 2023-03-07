package main

import (
	"fmt"
	_ "image/png"
	"log"
	"os"

	env "github.com/FachengG/AAgun/pkgs/envirment"
	"github.com/FachengG/AAgun/pkgs/obj"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	if repeatingKeyPressed(ebiten.KeySpace) {
		if g.gun.Shoot() {
			g.bullets = append(g.bullets, obj.NewBullet())
		}
	}
	if repeatingKeyPressed(ebiten.KeyLeft) {
		if g.gun.Angle > 0 {
			g.gun.Angle--
		}
	}
	if repeatingKeyPressed(ebiten.KeyRight) {
		if g.gun.Angle < 180 {
			g.gun.Angle++
		}
	}
	for i := range g.bullets {
		b := &g.bullets[i]
		b.UpdateSpeed(g.wind, g.frame)
		b.Movement(g.frame)
	}
	fmt.Print(g.gun.Angle)

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
	return 1000, 1000
}

func main() {

	g := &Game{frame: float64(1) / 60,
		wind: env.Wind{X: 0.01, Y: 0.01},
		gun:  obj.Gun{BulletsNum: 99999999999, Angle: 90}}

	ebiten.SetWindowTitle("bullet")
	fmt.Print("start game")
	if err := ebiten.RunGame(g); err != nil {
		f, _ := os.Create("err.log")
		f.Write([]byte(err.Error()))
		f.Close()
	}
}
