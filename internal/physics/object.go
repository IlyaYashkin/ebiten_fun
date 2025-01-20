package physics

import (
	"ebiten_fun/config"
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

func (o *Object) Update(neighbourObjects []*Object) {
	o.ApplyAcceleration()
	o.ApplyGravity()
	o.ApplyFriction()
	o.ApplyMaxSpeedLimit()

	//o.ApplyCollision(neighbourObjects)
	o.ApplyScreenCollision()

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

	friction := o.Velocity.Clone()
	friction.Normalize(velocityMagnitude)
	friction.MakeNegative()
	friction.Scale(o.FrictionCoefficient)

	o.Velocity.Add(friction)
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

func (o *Object) ApplyCollision(neighbourObjects []*Object) {
	for _, neighbourObject := range neighbourObjects {
		if o.checkCollision(neighbourObject) {
			resolveCollision(o, neighbourObject)
		}
	}
}

func (o *Object) GetRadius() float64 {
	return math.Sqrt(o.Width*o.Width+o.Height*o.Height) / 2
}

func (o *Object) checkCollision(otherObject *Object) bool {
	halfWidth1 := o.Width / 2
	halfHeight1 := o.Height / 2
	halfWidth2 := otherObject.Width / 2
	halfHeight2 := otherObject.Height / 2

	left1 := o.Position.X - halfWidth1
	right1 := o.Position.X + halfWidth1
	top1 := o.Position.Y - halfHeight1
	bottom1 := o.Position.Y + halfHeight1

	left2 := otherObject.Position.X - halfWidth2
	right2 := otherObject.Position.X + halfWidth2
	top2 := otherObject.Position.Y - halfHeight2
	bottom2 := otherObject.Position.Y + halfHeight2

	return left1 < right2 && right1 > left2 && top1 < bottom2 && bottom1 > top2
}

func resolveCollision(o1, o2 *Object) {
	halfWidth1, halfHeight1 := o1.Width/2, o1.Height/2
	halfWidth2, halfHeight2 := o2.Width/2, o2.Height/2

	deltaX := o1.Position.X - o2.Position.X
	deltaY := o1.Position.Y - o2.Position.Y

	overlapX := halfWidth1 + halfWidth2 - math.Abs(deltaX)
	overlapY := halfHeight1 + halfHeight2 - math.Abs(deltaY)

	if overlapX > 0 && overlapY > 0 {
		if overlapX < overlapY {
			if deltaX > 0 {
				o1.Position.X += overlapX / 2
				o2.Position.X -= overlapX / 2
			} else {
				o1.Position.X -= overlapX / 2
				o2.Position.X += overlapX / 2
			}
			o1.Velocity.X, o2.Velocity.X = o2.Velocity.X, o1.Velocity.X
		} else {
			if deltaY > 0 {
				o1.Position.Y += overlapY / 2
				o2.Position.Y -= overlapY / 2
			} else {
				o1.Position.Y -= overlapY / 2
				o2.Position.Y += overlapY / 2
			}
			o1.Velocity.Y, o2.Velocity.Y = o2.Velocity.Y, o1.Velocity.Y
		}

		o1.Destination = o1.Position
		o2.Destination = o2.Position
	}
}

func (o *Object) ApplyScreenCollision() {
	halfWidth := o.Width / 2
	halfHeight := o.Height / 2

	screenWidth := float64(config.ScreenWidth)
	screenHeight := float64(config.ScreenHeight)

	if o.Position.X-halfWidth < 0 {
		o.Position.X = halfWidth
		if o.Velocity.X < 0 {
			o.Velocity.X = -o.Velocity.X
		}
	}
	if o.Position.X+halfWidth > screenWidth {
		o.Position.X = screenWidth - halfWidth
		if o.Velocity.X > 0 {
			o.Velocity.X = -o.Velocity.X
		}
	}

	if o.Position.Y-halfHeight < 0 {
		o.Position.Y = halfHeight
		if o.Velocity.Y < 0 {
			o.Velocity.Y = -o.Velocity.Y
		}
	}
	if o.Position.Y+halfHeight > screenHeight {
		o.Position.Y = screenHeight - halfHeight
		if o.Velocity.Y > 0 {
			o.Velocity.Y = -o.Velocity.Y
		}
	}
}
