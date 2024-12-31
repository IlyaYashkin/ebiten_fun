package geo

import (
	"math"
)

type Vector struct {
	X, Y float64
}

func (v *Vector) GetLength() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() {
	if v.IsZero() {
		return
	}

	invLength := 1.0 / v.GetLength()

	v.Scale(invLength)
}

func (v *Vector) Scale(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

func (v *Vector) Subtract(v2 Vector) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vector) Multiply(v2 Vector) {
	v.X *= v2.X
	v.Y *= v2.Y
}

func (v *Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vector) Clone() Vector {
	return Vector{v.X, v.Y}
}

func (v *Vector) ToPoint() Point {
	return Point{v.X, v.Y}
}
