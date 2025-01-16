package main

import (
	"ebiten_fun/config"
	"ebiten_fun/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("particles")
	ebiten.SetFullscreen(true)
	ebiten.SetTPS(60)
	if err := ebiten.RunGame(&game.Game{NeedInit: true}); err != nil {
		log.Fatal(err)
	}
}
