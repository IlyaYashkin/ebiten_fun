package entity

import (
	"ebiten_fun/internal/control"
	"ebiten_fun/internal/geo"
	"ebiten_fun/internal/physics"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"slices"
)

type Character struct {
	Object physics.Object
	Image  *ebiten.Image
}

type CharacterParams struct {
	Width, Height int
}

func NewCharacter(initialPosition geo.Vector, image *ebiten.Image) Character {
	return Character{
		Object: physics.Object{
			Position:                initialPosition,
			Destination:             initialPosition,
			Width:                   75,
			Height:                  75,
			MaxSpeed:                40,
			AccelerationCoefficient: 5,
			FrictionCoefficient:     0.01,
		},
		Image: image,
	}
}

func NewImage() *ebiten.Image {
	image := ebiten.NewImage(75, 75)
	image.Fill(colornames.Red)

	return image
}

func (c *Character) GetObject() *physics.Object {
	return &c.Object
}

func (c *Character) GetImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	drawImageOptions := c.GetImagePosition()

	return c.Image, drawImageOptions
}

func (c *Character) Update(control control.Control) {
	c.ProcessInputs(control)
	c.Object.Move()
}

func (c *Character) GetImagePosition() *ebiten.DrawImageOptions {
	geoM := ebiten.GeoM{}
	geoM.Translate(c.Object.Position.X-c.Object.Width/2, c.Object.Position.Y-c.Object.Height/2)

	return &ebiten.DrawImageOptions{GeoM: geoM}
}

func (c *Character) ProcessInputs(controls control.Control) {
	destination := geo.Vector{X: 0, Y: 0}

	if slices.Contains(controls.PressedMouseButtons, ebiten.MouseButtonLeft) {
		destination = geo.Vector{X: float64(controls.CursorPosition.X), Y: float64(controls.CursorPosition.Y)}
		c.Object.Destination = destination

		return
	}

	if slices.Contains(controls.PressedKeys, ebiten.KeyW) {
		destination.Y -= 1
	}
	if slices.Contains(controls.PressedKeys, ebiten.KeyS) {
		destination.Y += 1
	}
	if slices.Contains(controls.PressedKeys, ebiten.KeyA) {
		destination.X -= 1
	}
	if slices.Contains(controls.PressedKeys, ebiten.KeyD) {
		destination.X += 1
	}

	destination.Add(c.Object.Position)

	c.Object.Destination = destination
}
