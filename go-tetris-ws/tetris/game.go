package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

// Game represents the Tetris game state.
type Game struct {
	currentTetromino *Tetromino
	lastFallTime     time.Time // Last time the Tetromino fell
}

// NewTetrisGame initializes a new Tetris game.
func NewTetrisGame() *Game {
	return &Game{
		currentTetromino: NewTetromino("T"), // Default Tetromino
		lastFallTime:     time.Now(),        // Initialize timer
	}
}

// Update handles game logic, including automatic falling and player input.
func (g *Game) Update() error {
	// Handle player movement
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.currentTetromino.Move(-1, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.currentTetromino.Move(1, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.currentTetromino.Move(0, 1) // Soft drop
	}

	// Automatic falling (Gravity)
	if time.Since(g.lastFallTime) > time.Second { // 1-second interval
		g.currentTetromino.Move(0, 1) // Move down
		g.lastFallTime = time.Now()   // Reset timer
	}

	return nil
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.currentTetromino.Draw(screen) // Draw the active Tetromino
}

// Layout sets the screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 480
}
