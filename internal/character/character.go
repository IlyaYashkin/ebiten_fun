package character

import (
	"ebiten_fun/internal/geo"
	"ebiten_fun/internal/object"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Character struct {
	Object object.Object
}

func GetCharacter(initialPosition geo.Point) Character {
	return Character{
		Object: object.Object{
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
	geoM.Translate(c.Object.Position.X, c.Object.Position.Y)

	return image, &ebiten.DrawImageOptions{GeoM: geoM}
}

func (c *Character) HandleInputs() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		point := geo.Point{X: float64(x), Y: float64(y)}
		c.Object.Destination = point

		return
	}

	dir := geo.Vector{X: 0, Y: 0}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dir.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		dir.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		dir.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		dir.X += 1
	}

	c.Object.Destination.X = c.Object.Position.X + dir.X
	c.Object.Destination.Y = c.Object.Position.Y + dir.Y
}

func (c *Character) Move() {
	c.Object.Move()
}
