package main

import (
	"fmt"
	_ "image/png"
	"log"

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
	time   float64
	wind   env.Wind
	gun    obj.Gun
	bullet obj.Bullet
}

func (g *Game) Update() error {
	g.bullet.UpdateSpeed(g.wind, g.time)
	g.bullet.Movement(g.time)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.bullet.Position.X, g.bullet.Position.Y)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}

func main() {
	g := &Game{}
	g.time = float64(1) / 60

	g.wind = env.Wind{X: 0.01, Y: 0.01}
	g.bullet = obj.Bullet{
		Position:       obj.XY_Vector{X: 50, Y: 500},
		Speed:          obj.XY_Vector{X: 50, Y: -100},
		Mass:           float64(1),
		Air_resistance: float64(0.01),
		Time:           float64(100),
	}
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("bullet")
	fmt.Print("start game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
