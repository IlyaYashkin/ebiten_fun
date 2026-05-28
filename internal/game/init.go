package game

import (
	"ebiten_fun/config"
	"ebiten_fun/internal/entity"
	"ebiten_fun/internal/geo"
	"ebiten_fun/internal/physics"
	"ebiten_fun/internal/structures"
	"math/rand/v2"
)

func (g *Game) initGame() {
	physics.SetGravity(config.Gravity)

	g.pixels = make([]byte, config.ScreenWidth*config.ScreenHeight*4)

	g.searchStructure = &structures.KDTree{}

	image := entity.NewImage()

	for i := 0; i < config.NumberOfCharacters; i++ {
		x := rand.IntN(config.ScreenWidth)
		y := rand.IntN(config.ScreenHeight)

		initialPosition := geo.Vector{X: float64(x), Y: float64(y)}
		char := entity.NewCharacter(
			initialPosition,
			image,
			config.CharacterMaxSpeed,
			config.CharacterAccelerationCoefficient,
			config.CharacterFrictionCoefficient,
		)

		g.entities = append(g.entities, &char)
	}

	g.startWorkers()

	g.NeedInit = false
}
