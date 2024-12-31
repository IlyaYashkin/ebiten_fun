package object

import "ebiten_fun/internal/geo"

var (
	epsilon = 0.001
)

type Object struct {
	Position    geo.Point
	Destination geo.Point
	Direction   geo.Vector
	MaxSpeed    float64
	Velocity    float64
}

func (o *Object) Move() {
	o.ApplyVelocity()

	o.Position.Move(o.Direction)
}

func (o *Object) ApplyVelocity() {
	directionLength := o.Direction.GetLength()
	if directionLength < epsilon {
		o.Direction = geo.Vector{}
	}

	direction := o.Position.GetVector(o.Destination)
	direction.Normalize()
	direction.Scale(o.MaxSpeed)

	direction.Subtract(o.Direction)
	direction.Scale(o.Velocity)

	o.Direction.Add(direction)
}
