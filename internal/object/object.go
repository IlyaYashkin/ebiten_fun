package object

import (
	"ebiten_fun/internal/geo"
)

type Object struct {
	Position               geo.Vector
	Destination            geo.Vector
	Velocity               geo.Vector
	MaxSpeed               float64
	AccelerationMultiplier float64
}

func (o *Object) Move() {
	o.ApplyAcceleration()

	o.Position.Add(o.Velocity)
}

func (o *Object) ApplyAcceleration() {
	acceleration := o.Destination.Clone()
	acceleration.Subtract(o.Position)
	acceleration.Normalize()
	acceleration.Scale(o.MaxSpeed)
	acceleration.Subtract(o.Velocity)
	acceleration.Scale(o.AccelerationMultiplier)

	o.Velocity.Add(acceleration)
}
