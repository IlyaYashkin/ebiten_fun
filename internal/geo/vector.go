package geo

import (
	"math"
)

type Vector struct {
	X, Y float64
}

func (v1 *Vector) GetValues() (float64, float64) {
	return v1.X, v1.Y
}

func (v1 *Vector) GetMagnitude() float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
}

func (v1 *Vector) Normalize(magnitude float64) {
	if v1.IsZero() {
		return
	}

	//magnitude += 1e-16

	invLength := 1.0 / magnitude

	v1.Scale(invLength)
}

func (v1 *Vector) Scale(scalar float64) {
	v1.X *= scalar
	v1.Y *= scalar
}

func (v1 *Vector) Subtract(v2 Vector) {
	v1.X -= v2.X
	v1.Y -= v2.Y
}

func (v1 *Vector) Add(v2 Vector) {
	v1.X += v2.X
	v1.Y += v2.Y
}

func (v1 *Vector) MakeNegative() {
	v1.X = -v1.X
	v1.Y = -v1.Y
}

func (v1 *Vector) MakeZero() {
	v1.X = 0
	v1.Y = 0
}

func (v1 *Vector) IsZero() bool {
	return v1.X == 0 && v1.Y == 0
}

func (v1 *Vector) Equals(v2 Vector) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v1 *Vector) Clone() Vector {
	return Vector{v1.X, v1.Y}
}

func (v1 *Vector) Distance(v2 Vector) float64 {
	return math.Sqrt((v1.X-v2.X)*(v1.X-v2.X) + (v1.Y-v2.Y)*(v1.Y-v2.Y))
}
