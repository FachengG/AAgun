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
	Time           float64
}

func (b *Bullet) UpdateSpeed(wind env.Wind, time float64) {
	b.Speed.X += (-b.Air_resistance * math.Abs(wind.Y-b.Speed.X)) * time
	b.Speed.Y -= (env.Gravity - b.Air_resistance*math.Abs(wind.Y-b.Speed.Y)) * time
}

func (b *Bullet) Movement(time float64) {
	b.Position.X += b.Speed.X * time
	b.Position.Y += b.Speed.Y * time
	b.Time = b.Time - time
}

func (b Bullet) Boundary(upper_x float64, lower_x float64, upper_y float64, lower_y float64) bool {
	return !(b.Position.X > upper_x || b.Position.X < lower_x || b.Position.Y > upper_y || b.Position.Y < lower_y)
}

func (b Bullet) Explosition() bool {
	return b.Time <= float64(0)
}
