package structures

import (
	"ebiten_fun/internal/entity"
	"math"
	"sort"
)

type KDTree struct {
	Node *Node
}

func (KDTree *KDTree) Update(entities []entity.Entity) {
	KDTree.Node = buildKDTree(entities, 0)
}

func (KDTree *KDTree) GetNeighbours(ent entity.Entity) []entity.Entity {
	var neighbors []entity.Entity

	findNeighborsInRadius(KDTree.Node, ent, 300, &neighbors)

	return neighbors
}

type Node struct {
	Entity     entity.Entity
	Left       *Node
	Right      *Node
	IsVertical bool
}

type NearestNeighborResult struct {
	Entity   entity.Entity
	Distance float64
}

func buildKDTree(entities []entity.Entity, depth int) *Node {
	if len(entities) == 0 {
		return nil
	}

	isVertical := depth%2 == 0

	sort.Slice(entities, func(i, j int) bool {
		if isVertical {
			return entities[i].GetObject().Position.X < entities[j].GetObject().Position.X
		} else {
			return entities[i].GetObject().Position.Y < entities[j].GetObject().Position.Y
		}
	})

	median := len(entities) / 2

	return &Node{
		Entity:     entities[median],
		Left:       buildKDTree(entities[:median], depth+1),
		Right:      buildKDTree(entities[median+1:], depth+1),
		IsVertical: isVertical,
	}
}

func findNeighborsInRadius(root *Node, target entity.Entity, radius float64, neighbors *[]entity.Entity) {
	if root == nil {
		return
	}

	if root.Entity == target {
		findNeighborsInRadius(root.Left, target, radius, neighbors)
		findNeighborsInRadius(root.Right, target, radius, neighbors)
		return
	}

	targetRadius := target.GetObject().GetRadius()
	currentRadius := root.Entity.GetObject().GetRadius()

	targetPos := target.GetObject().Position
	currentPos := root.Entity.GetObject().Position
	distance := targetPos.Distance(currentPos)

	if distance <= targetRadius+currentRadius {
		*neighbors = append(*neighbors, root.Entity)
	}

	var primary, secondary *Node
	if root.IsVertical {
		if targetPos.X < currentPos.X {
			primary, secondary = root.Left, root.Right
		} else {
			primary, secondary = root.Right, root.Left
		}
	} else {
		if targetPos.Y < currentPos.Y {
			primary, secondary = root.Left, root.Right
		} else {
			primary, secondary = root.Right, root.Left
		}
	}

	findNeighborsInRadius(primary, target, targetRadius, neighbors)

	if root.IsVertical {
		if math.Abs(targetPos.X-currentPos.X) <= targetRadius {
			findNeighborsInRadius(secondary, target, targetRadius, neighbors)
		}
	} else {
		if math.Abs(targetPos.Y-currentPos.Y) <= targetRadius {
			findNeighborsInRadius(secondary, target, targetRadius, neighbors)
		}
	}
}

func findNearestNeighbor(root *Node, target entity.Entity, best *NearestNeighborResult) *NearestNeighborResult {
	if root == nil {
		return best
	}

	if root.Entity == target {
		best = findNearestNeighbor(root.Left, target, best)
		best = findNearestNeighbor(root.Right, target, best)
		return best
	}

	targetPos := target.GetObject().Position
	currentPos := root.Entity.GetObject().Position
	distance := targetPos.Distance(currentPos)

	if best == nil || distance < best.Distance {
		best = &NearestNeighborResult{
			Entity:   root.Entity,
			Distance: distance,
		}
	}

	var primary, secondary *Node
	if root.IsVertical {
		if targetPos.X < currentPos.X {
			primary, secondary = root.Left, root.Right
		} else {
			primary, secondary = root.Right, root.Left
		}
	} else {
		if targetPos.Y < currentPos.Y {
			primary, secondary = root.Left, root.Right
		} else {
			primary, secondary = root.Right, root.Left
		}
	}

	best = findNearestNeighbor(primary, target, best)

	if root.IsVertical {
		if math.Abs(targetPos.X-currentPos.X) < best.Distance {
			best = findNearestNeighbor(secondary, target, best)
		}
	} else {
		if math.Abs(targetPos.Y-currentPos.Y) < best.Distance {
			best = findNearestNeighbor(secondary, target, best)
		}
	}

	return best
}
