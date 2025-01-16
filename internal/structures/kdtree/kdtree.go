package kdtree

import (
	"ebiten_fun/internal/entity"
	"math"
	"sort"
)

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

func BuildKdTree(entities []entity.Entity, depth int) *Node {
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
		Left:       BuildKdTree(entities[:median], depth+1),
		Right:      BuildKdTree(entities[median+1:], depth+1),
		IsVertical: isVertical,
	}
}

func FindNeighborsInRadius(root *Node, target entity.Entity, radius float64, neighbors *[]entity.Entity) {
	if root == nil {
		return
	}

	if root.Entity == target {
		FindNeighborsInRadius(root.Left, target, radius, neighbors)
		FindNeighborsInRadius(root.Right, target, radius, neighbors)
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

	// Search the primary branch
	FindNeighborsInRadius(primary, target, radius, neighbors)

	// Check if we need to search the secondary branch
	if root.IsVertical {
		if math.Abs(targetPos.X-currentPos.X) <= radius {
			FindNeighborsInRadius(secondary, target, radius, neighbors)
		}
	} else {
		if math.Abs(targetPos.Y-currentPos.Y) <= radius {
			FindNeighborsInRadius(secondary, target, radius, neighbors)
		}
	}
}

func FindNearestNeighbor(root *Node, target entity.Entity, best *NearestNeighborResult) *NearestNeighborResult {
	if root == nil {
		return best
	}

	if root.Entity == target {
		best = FindNearestNeighbor(root.Left, target, best)
		best = FindNearestNeighbor(root.Right, target, best)
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

	// Search the primary branch
	best = FindNearestNeighbor(primary, target, best)

	// Check if we need to search the secondary branch
	if root.IsVertical {
		if math.Abs(targetPos.X-currentPos.X) < best.Distance {
			best = FindNearestNeighbor(secondary, target, best)
		}
	} else {
		if math.Abs(targetPos.Y-currentPos.Y) < best.Distance {
			best = FindNearestNeighbor(secondary, target, best)
		}
	}

	return best
}
