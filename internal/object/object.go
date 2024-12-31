package object

import "ebiten_fun/internal/geo"

var (
	epsilon = 0.001
)

type Object struct {
	Position    geo.Vector
	Destination geo.Vector
	Direction   geo.Vector
	MaxSpeed    float64
	Velocity    float64
}

func (o *Object) Move() {
	o.ApplyVelocity()

	o.Position.Add(o.Direction)
}

func (o *Object) ApplyVelocity() {
	directionLength := o.Direction.GetLength()
	if directionLength < epsilon {
		o.Direction = geo.Vector{}
	}

	direction := o.Destination.Clone()
	direction.Subtract(o.Position)

	direction.Normalize()

	direction.Scale(o.MaxSpeed)

	direction.Subtract(o.Direction)
	direction.Scale(o.Velocity)

	o.Direction.Add(direction)
}
