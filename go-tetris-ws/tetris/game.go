package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

// Game represents the Tetris game state.
type Game struct {
	currentTetromino *Tetromino
	board            [][]bool
	lastFallTime     time.Time
	lastMoveTime     time.Time
	lastKeyState     map[ebiten.Key]bool // Track key states
	score            int
	gameOver         bool
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
	currentTime := time.Now()

	// Handle rotation (detect key press only once per press)
	if g.isKeyJustPressed(ebiten.KeyZ) {
		g.currentTetromino.RotateCounterClockwise(g.board)
	}
	if g.isKeyJustPressed(ebiten.KeyX) {
		g.currentTetromino.RotateClockwise(g.board)
	}

	// Handle left/right movement with delay
	if currentTime.Sub(g.lastMoveTime) > MoveInterval {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			g.currentTetromino.Move(-1, 0, g.board)
			g.lastMoveTime = currentTime
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			g.currentTetromino.Move(1, 0, g.board)
			g.lastMoveTime = currentTime
		}
	}

	// Soft drop (manual down movement)
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if !g.currentTetromino.Move(0, 1, g.board) {
			g.lockTetromino()
		}
	}

	// Automatic falling (Gravity)
	if currentTime.Sub(g.lastFallTime) > GravityInterval {
		if !g.currentTetromino.Move(0, 1, g.board) {
			g.lockTetromino()
		}
		g.lastFallTime = currentTime
	}

	return nil
}

// canMove checks if the Tetromino can move without colliding with other blocks.
func (g *Game) canMove(dx, dy int) bool {
	for _, pos := range TetrominoShapes[g.currentTetromino.shape] {
		x, y := g.currentTetromino.x+pos[0]+dx, g.currentTetromino.y+pos[1]+dy
		if x < 0 || x >= BoardWidth || y >= BoardHeight || (y >= 0 && g.board[y][x]) {
			return false // Collision detected
		}
	}
	return true
}

// spawnNewTetromino creates a new Tetromino and checks for game over.
func (g *Game) spawnNewTetromino() {
	newTetromino := NewTetromino()

	// Check if the new Tetromino collides immediately (Game Over)
	for _, pos := range TetrominoShapes[newTetromino.shape] {
		x, y := newTetromino.x+pos[0], newTetromino.y+pos[1]
		if y >= 0 && g.board[y][x] {
			g.gameOver = true
			return
		}
	}

	g.currentTetromino = newTetromino
}

// updateScore increases the score based on the number of rows cleared.
func (g *Game) updateScore(rowsCleared int) {
	points := map[int]int{
		1: 100, // Single row
		2: 300, // Double row
		3: 500, // Triple row
		4: 800, // Tetris (4 rows)
	}
	g.score += points[rowsCleared]
}
