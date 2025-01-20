package game

import (
	"ebiten_fun/config"
	"ebiten_fun/internal/control"
	"ebiten_fun/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	NeedInit      bool
	ticks         []chan struct{}
	workersNumber int
	controls      control.Control
	character     *entity.Character

	searchStructure entity.SearchStructure

	entities []entity.Entity
	pixels   []byte
}

func (g *Game) Update() error {
	if g.NeedInit {
		g.initGame()
	}

	g.handleInputs()

	//g.character.Update(g.controls, g.searchStructure)
	//
	//var controls control.Control

	//for _, ent := range g.entities {
	//	//if ent == g.character {
	//	//	continue
	//	//}
	//	ent.Update(g.controls, &structures.Mock{})
	//}

	g.callWorkers()

	//g.searchStructure.Update(g.entities)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawPixels(screen)

	//g.drawImages(screen)

	//g.drawKDTree(screen)

	//g.drawRadii(screen)

	g.printDebug(screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
