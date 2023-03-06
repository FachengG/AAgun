package obj

import (
	"math"

	env "github.com/FachengG/AAgun/pkgs/envirment"
)

type Bullet struct {
	Position       XY_Vector
	Speed          XY_Vector
	Mass           float64
	Air_resistance float64
	Lifetime       float64
}

func (b *Bullet) UpdateSpeed(wind env.Wind, grid float64) {
	b.Speed.X += (-b.Air_resistance * math.Abs(wind.Y-b.Speed.X)) * grid
	b.Speed.Y -= (env.Gravity - b.Air_resistance*math.Abs(wind.Y-b.Speed.Y)) * grid
}

func (b *Bullet) Movement(grid float64) {
	b.Position.X += b.Speed.X * grid
	b.Position.Y += b.Speed.Y * grid
	b.Lifetime = b.Lifetime - grid
}

func (b Bullet) Boundary(upper_x float64, lower_x float64, upper_y float64, lower_y float64) bool {
	return !(b.Position.X > upper_x || b.Position.X < lower_x || b.Position.Y > upper_y || b.Position.Y < lower_y)
}

func (b Bullet) Explosition() bool {
	return b.Lifetime <= float64(0)
}

func (b Bullet) New() Bullet {
	bullet := Bullet{
		Position:       XY_Vector{X: 50, Y: 500},
		Speed:          XY_Vector{X: 50, Y: -100},
		Mass:           float64(1),
		Air_resistance: float64(0.01),
		Lifetime:       float64(100),
	}
	return bullet
}
