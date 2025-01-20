package entity

import (
	"ebiten_fun/internal/control"
	"ebiten_fun/internal/physics"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	GetObject() *physics.Object
	GetImage() (*ebiten.Image, *ebiten.DrawImageOptions)
	Update(control control.Control, searchStructure SearchStructure)
}

type SearchStructure interface {
	Update([]Entity)
	GetNeighbours(Entity) []Entity
}
