package main

import (
	"ebiten_fun/internal/character"
	"ebiten_fun/internal/geo"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"log"
	"math/rand/v2"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
	needInit   bool
	keys       []ebiten.Key
	characters []*character.Character
	pixels     []byte
}

func (g *Game) Update() error {
	if g.needInit {
		g.initGame()
	}

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	g.HandleInputs()

	for _, char := range g.characters {
		char.Move()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.drawImages(screen)

	g.drawPixels(screen)

	g.printDebug(screen)
}

func (g *Game) drawPixels(screen *ebiten.Image) {
	g.pixels = make([]byte, screenWidth*screenHeight*4)

	for _, char := range g.characters {
		x, y := char.GetRawPosition()

		if x > 0 && x < screenWidth && y > 0 && y < screenHeight {
			ind := (int(y)*screenWidth + int(x)) * 4

			g.pixels[ind] = colornames.Red.R
			g.pixels[ind+1] = colornames.Red.G
			g.pixels[ind+2] = colornames.Red.B
			g.pixels[ind+3] = colornames.Red.A
		}
	}

	screen.WritePixels(g.pixels)
}

func (g *Game) drawImages(screen *ebiten.Image) {
	for _, char := range g.characters {
		options := char.GetPosition()

		screen.DrawImage(char.Image, options)
	}
}

func (g *Game) printDebug(screen *ebiten.Image) {
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			"TPS: %0.2f\nFPS: %0.2f\n %+v\n",
			ebiten.ActualTPS(),
			ebiten.ActualFPS(),
			g.keys,
		),
	)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) HandleInputs() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		destination := geo.Vector{X: float64(x), Y: float64(y)}

		for _, char := range g.characters {
			char.Object.Destination = destination
		}

		return
	}

	for _, char := range g.characters {
		destination := geo.Vector{X: 0, Y: 0}

		destination.Add(char.Object.Position)

		char.Object.Destination = destination
	}

}

func (g *Game) initGame() {
	g.pixels = make([]byte, screenWidth*screenHeight*4)

	image := character.GetImage()

	for i := 0; i < 1_000_000; i++ {
		x := rand.IntN(screenWidth)
		y := rand.IntN(screenHeight)

		initialPosition := geo.Vector{X: float64(x), Y: float64(y)}
		char := character.GetCharacter(initialPosition, image)
		g.characters = append(g.characters, &char)
		//g.character = character.GetCharacter(initialPosition)
	}

	g.needInit = false
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("particles")
	ebiten.SetFullscreen(true)
	ebiten.SetTPS(60)
	if err := ebiten.RunGame(&Game{needInit: true}); err != nil {
		log.Fatal(err)
	}
}
