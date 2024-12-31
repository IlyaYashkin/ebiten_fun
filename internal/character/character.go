package character

import (
	"ebiten_fun/internal/geo"
	"ebiten_fun/internal/object"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Character struct {
	object object.Object
}

func GetCharacter(initialPosition geo.Vector) Character {
	return Character{
		object: object.Object{
			Position:    initialPosition,
			Destination: initialPosition,
			MaxSpeed:    5.,
			Velocity:    .01,
		},
	}
}

func (c *Character) GetImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	image := ebiten.NewImage(20, 20)
	image.Fill(colornames.Lime)

	geoM := ebiten.GeoM{}
	geoM.Translate(c.object.Position.X, c.object.Position.Y)

	return image, &ebiten.DrawImageOptions{GeoM: geoM}
}

func (c *Character) HandleInputs() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		destination := geo.Vector{X: float64(x), Y: float64(y)}
		c.object.Destination = destination

		return
	}

	destination := geo.Vector{X: 0, Y: 0}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		destination.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		destination.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		destination.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		destination.X += 1
	}

	destination.Add(c.object.Position)

	c.object.Destination = destination
}

func (c *Character) Move() {
	c.object.Move()
}
