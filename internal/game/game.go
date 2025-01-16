package game

import (
	"ebiten_fun/config"
	"ebiten_fun/internal/control"
	"ebiten_fun/internal/entity"
	"ebiten_fun/internal/structures/kdtree"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	NeedInit      bool
	ticks         []chan struct{}
	workersNumber int
	controls      control.Control
	character     *entity.Character

	entityKDTree    *kdtree.Node
	nearest         entity.Entity
	nearestInRadius []entity.Entity

	entities []entity.Entity
	pixels   []byte
}

func (g *Game) Update() error {
	if g.NeedInit {
		g.initGame()
	}

	g.handleInputs()

	g.callWorkers()

	//for _, ent := range g.entities {
	//	ent.Update(g.controls)
	//}

	//g.character.Update(g.controls)

	//g.entityKDTree = kdtree.BuildKdTree(g.entities, 0)
	//
	//var neighbors []entity.Entity
	//kdtree.FindNeighborsInRadius(g.entityKDTree, g.character, 300, &neighbors)
	//
	//g.nearestInRadius = neighbors

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawPixels(screen)

	//g.drawImages(screen)

	//g.drawKDTree(screen)

	//g.drawLineToNearest(screen)

	//g.drawRadii(screen)

	g.printDebug(screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
