package character

import (
	"ebiten_fun/internal/geo"
	"ebiten_fun/internal/object"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Character struct {
	Object object.Object
	Image  *ebiten.Image
}

func GetCharacter(initialPosition geo.Vector, image *ebiten.Image) Character {
	return Character{
		Object: object.Object{
			Position:               initialPosition,
			Destination:            initialPosition,
			MaxSpeed:               1000.,
			AccelerationMultiplier: .01,
		},
		Image: image,
	}
}

func GetImage() *ebiten.Image {
	image := ebiten.NewImage(20, 20)
	image.Fill(colornames.Lime)

	return image
}

func (c *Character) GetPosition() *ebiten.DrawImageOptions {
	geoM := ebiten.GeoM{}
	geoM.Translate(c.Object.Position.X, c.Object.Position.Y)

	return &ebiten.DrawImageOptions{GeoM: geoM}
}

func (c *Character) GetRawPosition() (float64, float64) {
	return c.Object.Position.X, c.Object.Position.Y
}

func (c *Character) HandleInputs() {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		destination := geo.Vector{X: float64(x), Y: float64(y)}
		c.Object.Destination = destination

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

	destination.Add(c.Object.Position)

	c.Object.Destination = destination
}

func (c *Character) Move() {
	c.Object.Move()
}
