package main

import (
	"ebiten_fun/internal/character"
	"ebiten_fun/internal/geo"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

type Game struct {
	needInit  bool
	character character.Character
}

func (g *Game) Update() error {
	if g.needInit {
		g.init()
	}

	g.character.HandleInputs()
	g.character.Move()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	image, options := g.character.GetImage()

	screen.DrawImage(image, options)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) init() {
	initialPosition := geo.Vector{X: screenWidth / 2, Y: screenHeight / 2}
	g.character = character.GetCharacter(initialPosition)

	g.needInit = false
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(120)
	if err := ebiten.RunGame(&Game{needInit: true}); err != nil {
		log.Fatal(err)
	}
}
