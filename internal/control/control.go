package control

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Control struct {
	PressedKeys         []ebiten.Key
	PressedMouseButtons []ebiten.MouseButton
	CursorPosition      struct {
		X, Y int
	}
}

var (
	MouseButtons = []ebiten.MouseButton{
		ebiten.MouseButtonLeft,
		ebiten.MouseButtonMiddle,
		ebiten.MouseButtonRight,
	}
)
