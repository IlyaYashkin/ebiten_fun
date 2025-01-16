package physics

import "ebiten_fun/internal/geo"

var (
	gravityScalar float64 = 1
)

func SetGravity(g float64) {
	gravityScalar = g
}

func GetGravityVector() geo.Vector {
	vector := geo.Vector{X: 0, Y: 1}
	vector.Scale(gravityScalar)
	return vector
}
