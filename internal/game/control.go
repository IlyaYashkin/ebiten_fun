package game

import (
	"ebiten_fun/internal/control"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) handleInputs() {
	controls := &g.controls

	controls.PressedMouseButtons = []ebiten.MouseButton{}
	controls.PressedKeys = inpututil.AppendPressedKeys(controls.PressedKeys[:0])
	x, y := ebiten.CursorPosition()
	controls.CursorPosition.X = x
	controls.CursorPosition.Y = y

	for _, mouseButton := range control.MouseButtons {
		if ebiten.IsMouseButtonPressed(mouseButton) {
			controls.PressedMouseButtons = append(controls.PressedMouseButtons[:0], mouseButton)
		}
	}
}
