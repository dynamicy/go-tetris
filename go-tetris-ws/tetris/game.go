package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

// Game represents the Tetris game state.
type Game struct {
	currentTetromino *Tetromino
	ghostTetromino   *Tetromino // New ghost piece
	board            [][]bool
	lastFallTime     time.Time
	lastMoveTime     time.Time
	lastKeyState     map[ebiten.Key]bool // Track key states
	score            int
	linesCleared     int
	level            int
	gameOver         bool
	hardDropActive   bool // Track if Hard Drop is in progress
}

// NewTetrisGame initializes and returns a new Tetris game instance.
func NewTetrisGame() *Game {
	board := make([][]bool, BoardHeight)
	for i := range board {
		board[i] = make([]bool, BoardWidth)
	}

	return &Game{
		currentTetromino: NewTetromino(),
		board:            board,
		lastFallTime:     time.Now(),
		lastMoveTime:     time.Now(),
		lastKeyState:     make(map[ebiten.Key]bool), // Initialize key tracking
	}
}

// Update handles game logic, including movement and rotation.
func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.ResetGame() // Restart the game
		}
		return nil // Prevent input if game is over
	}

	g.updateGhostPiece() // Ensure ghost is always updated

	currentTime := time.Now()

	// Handle player inputs
	g.handleInput(currentTime)

	// Process Tetromino movement
	g.processMovement(currentTime)

	return nil
}

// handleInput processes key presses for rotation, movement, and hard drop.
func (g *Game) handleInput(currentTime time.Time) {
	// Rotation (detect key press once per press)
	if g.isKeyJustPressed(ebiten.KeyZ) {
		g.currentTetromino.RotateCounterClockwise(g.board)
	}
	if g.isKeyJustPressed(ebiten.KeyX) {
		g.currentTetromino.RotateClockwise(g.board)
	}

	// Hard Drop (instant fall)
	if g.isKeyJustPressed(ebiten.KeySpace) {
		g.hardDropActive = true
	}
}

// processMovement updates Tetromino position based on inputs and gravity.
func (g *Game) processMovement(currentTime time.Time) {
	if g.hardDropActive {
		// Move downward until collision
		for g.currentTetromino.Move(0, 1, g.board) {
		}
		g.lockTetromino()
		g.hardDropActive = false
		return
	}

	// Handle left/right movement with DAS (Delayed Auto Shift)
	moveLeft := ebiten.IsKeyPressed(ebiten.KeyLeft)
	moveRight := ebiten.IsKeyPressed(ebiten.KeyRight)

	if moveLeft || moveRight {
		if g.lastMoveTime.IsZero() || currentTime.Sub(g.lastMoveTime) > InitialMoveDelay {
			if currentTime.Sub(g.lastMoveTime) > MoveRepeatRate {
				if moveLeft {
					g.currentTetromino.Move(-1, 0, g.board)
				} else if moveRight {
					g.currentTetromino.Move(1, 0, g.board)
				}
				g.lastMoveTime = currentTime
			}
		}
	} else {
		g.lastMoveTime = time.Time{} // Reset move timer if no key is pressed
	}

	// Handle downward movement (gravity and soft drop)
	shouldLock := false
	if ebiten.IsKeyPressed(ebiten.KeyDown) || currentTime.Sub(g.lastFallTime) > GravityInterval {
		if !g.currentTetromino.Move(0, 1, g.board) {
			shouldLock = true
		}
		g.lastFallTime = currentTime
	}

	if shouldLock {
		g.lockTetromino()
	}
}

// updateScore increases the score based on the number of rows cleared.
func (g *Game) updateScore(rowsCleared int) {
	if points, exists := PointsPerLine[rowsCleared]; exists {
		g.score += points
	}

	// Track total lines cleared
	g.linesCleared += rowsCleared

	// Increase level every 10 lines
	if g.linesCleared/10 > g.level {
		g.level++
	}
}

// ResetGame resets the game state, allowing for a fresh start.
func (g *Game) ResetGame() {
	// Clear the game board
	g.board = make([][]bool, BoardHeight)
	for i := range g.board {
		g.board[i] = make([]bool, BoardWidth)
	}

	// Reset score and game status
	g.score = 0
	g.gameOver = false
	g.hardDropActive = false

	// Spawn a new Tetromino
	g.currentTetromino = NewTetromino()

	// Reset timing variables
	g.lastFallTime = time.Now()
	g.lastMoveTime = time.Now()
	g.lastKeyState = make(map[ebiten.Key]bool) // Reset key tracking
}
