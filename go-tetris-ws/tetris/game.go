package tetris

import "github.com/hajimehoshi/ebiten/v2"

// Game represents the Tetris game.
type Game struct{}

// NewTetrisGame creates a new instance of the Tetris game.
func NewTetrisGame() *Game {
	return &Game{}
}

// Update handles game logic updates (TODO: Implement game state updates).
func (g *Game) Update() error {
	// TODO: Handle block falling, collision detection, and line clearing.
	return nil
}

// Draw renders the game screen (TODO: Draw blocks and background).
func (g *Game) Draw(screen *ebiten.Image) {
	// TODO: Draw Tetrominoes on the screen.
}

// Layout sets the screen size for the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 480
}
