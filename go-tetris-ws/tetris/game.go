package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game represents the Tetris game state.
type Game struct {
	currentTetromino *Tetromino // The currently falling Tetromino
}

// NewTetrisGame initializes and returns a new Tetris game instance.
func NewTetrisGame() *Game {
	return &Game{
		currentTetromino: NewTetromino("T"), // Default Tetromino at the start
	}
}

// Update processes game logic, including keyboard input for movement.
func (g *Game) Update() error {
	// Move Tetromino left
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.currentTetromino.Move(-1, 0)
	}

	// Move Tetromino right
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.currentTetromino.Move(1, 0)
	}

	// Move Tetromino down faster (soft drop)
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.currentTetromino.Move(0, 1)
	}

	return nil
}

// Draw renders the current game state onto the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.currentTetromino.Draw(screen) // Draw the active Tetromino
}

// Layout defines the logical screen size for rendering.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 480
}
