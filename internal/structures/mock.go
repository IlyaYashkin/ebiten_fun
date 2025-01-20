package structures

import "ebiten_fun/internal/entity"

type Mock struct {
}

func (m *Mock) Update(entities []entity.Entity) {
}

func (m *Mock) GetNeighbours(ent entity.Entity) []entity.Entity {
	var neighbors []entity.Entity

	return neighbors
}
