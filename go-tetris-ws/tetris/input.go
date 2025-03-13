package tetris

import "github.com/hajimehoshi/ebiten/v2"

// isKeyJustPressed checks if a key was just pressed (prevents holding).
func (g *Game) isKeyJustPressed(key ebiten.Key) bool {
	pressed := ebiten.IsKeyPressed(key)
	if pressed && !g.lastKeyState[key] {
		g.lastKeyState[key] = true
		return true
	}
	g.lastKeyState[key] = pressed
	return false
}
