package game

import (
	"ebiten_fun/config"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

func (g *Game) drawPixels(screen *ebiten.Image) {
	g.pixels = make([]byte, config.ScreenWidth*config.ScreenHeight*4)

	for _, char := range g.entities {
		x, y := char.GetObject().Position.GetValues()

		if x > 0 && int(x) < config.ScreenWidth && y > 0 && int(y) < config.ScreenHeight {
			ind := (int(y)*config.ScreenWidth + int(x)) * 4

			g.pixels[ind] = colornames.Red.R
			g.pixels[ind+1] = colornames.Red.G
			g.pixels[ind+2] = colornames.Red.B
			g.pixels[ind+3] = colornames.Red.A
		}
	}

	screen.WritePixels(g.pixels)
}

func (g *Game) drawImages(screen *ebiten.Image) {
	for _, char := range g.entities {
		image, options := char.GetImage()

		screen.DrawImage(image, options)
	}
}

func (g *Game) printDebug(screen *ebiten.Image) {
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			"TPS: %0.2f\nFPS: %0.2f\n%+v\n%+v\n%+v",
			ebiten.ActualTPS(),
			ebiten.ActualFPS(),
			g.controls.PressedKeys,
			g.controls.PressedMouseButtons,
			g.character.GetObject().Velocity.GetMagnitude(),
		),
	)
}

func (g *Game) drawRadii(screen *ebiten.Image) {
	for _, ent := range g.entities {
		vector.StrokeCircle(
			screen,
			float32(ent.GetObject().Position.X),
			float32(ent.GetObject().Position.Y),
			float32(ent.GetObject().GetRadius()),
			1,
			colornames.White,
			false,
		)
	}
}
