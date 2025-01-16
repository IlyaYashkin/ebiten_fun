package game

import (
	"ebiten_fun/config"
	"ebiten_fun/internal/entity"
	"sync"
)

func (g *Game) callWorkers() {
	for _, tick := range g.ticks {
		select {
		case tick <- struct{}{}:
		}
	}
}

func (g *Game) startWorkers() {
	g.workersNumber = config.WorkersNumber
	var wg sync.WaitGroup

	for i := 0; i < g.workersNumber; i++ {
		g.ticks = append(g.ticks, make(chan struct{}))
	}

	batchSize := len(g.entities) / g.workersNumber
	remainder := len(g.entities) % batchSize

	for i := 0; i < g.workersNumber; i++ {
		wg.Add(1)

		offset := batchSize * i

		if i == g.workersNumber-1 {
			go g.worker(i, g.entities[offset:offset+batchSize+remainder], &wg, g.ticks[i])
			continue
		}

		go g.worker(i, g.entities[offset:offset+batchSize], &wg, g.ticks[i])
	}
}

func (g *Game) worker(id int, entities []entity.Entity, wg *sync.WaitGroup, tick <-chan struct{}) {
	defer wg.Done()

	for range tick {
		for _, ent := range entities {
			ent.Update(g.controls)
		}
	}
}
