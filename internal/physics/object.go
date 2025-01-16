package physics

import (
	"ebiten_fun/internal/geo"
	"math"
)

type Object struct {
	Position                geo.Vector
	Destination             geo.Vector
	Velocity                geo.Vector
	Width                   float64
	Height                  float64
	MaxSpeed                float64
	AccelerationCoefficient float64
	FrictionCoefficient     float64
}

type HasObject interface {
	GetObject() *Object
}

func (o *Object) Move() {
	o.ApplyAcceleration()
	o.ApplyFriction()
	o.ApplyMaxSpeedLimit()

	o.Position.Add(o.Velocity)
}

func (o *Object) ApplyGravity() {
	gravityVector := GetGravityVector()

	o.Velocity.Add(gravityVector)
}

func (o *Object) ApplyAcceleration() {
	acceleration := o.Destination.Clone()
	acceleration.Subtract(o.Position)
	acceleration.Normalize(acceleration.GetMagnitude())
	acceleration.Scale(o.AccelerationCoefficient)

	o.Velocity.Add(acceleration)
}

func (o *Object) ApplyFriction() {
	velocityMagnitude := o.Velocity.GetMagnitude()

	if velocityMagnitude < o.FrictionCoefficient {
		o.Velocity.MakeZero()
	}

	//if o.Destination.Equals(o.Position) {
	friction := o.Velocity.Clone()
	friction.Normalize(velocityMagnitude)
	friction.MakeNegative()
	friction.Scale(o.FrictionCoefficient)

	o.Velocity.Add(friction)
	//}
}

func (o *Object) ApplyMaxSpeedLimit() {
	if o.MaxSpeed <= 0 {
		return
	}

	magnitude := o.Velocity.GetMagnitude()

	if magnitude > o.MaxSpeed {
		o.Velocity.Normalize(magnitude)
		o.Velocity.Scale(o.MaxSpeed)
	}
}

func (o *Object) ApplyMaxSpeedLimitOld() {
	magnitude := o.Velocity.GetMagnitude()
	scalar := magnitude - o.MaxSpeed
	scalar = (scalar + math.Abs(scalar)) / 2
	scalar = 1 - scalar/o.MaxSpeed
	o.Velocity.Scale(scalar)
}

func (o *Object) GetRadius() float64 {
	return math.Sqrt(o.Width*o.Width+o.Height*o.Height) / 2
}
