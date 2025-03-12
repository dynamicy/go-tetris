package tetris

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Constants for the game board
const (
	BoardWidth  = 10 // Number of columns
	BoardHeight = 20 // Number of rows
	CellSize    = 24 // Size of each cell in pixels
)

// Game represents the Tetris game state.
type Game struct {
	currentTetromino *Tetromino
	board            [][]bool  // Stores landed Tetrominoes
	lastFallTime     time.Time // Last time the Tetromino fell
}

// NewTetrisGame initializes a new Tetris game instance.
func NewTetrisGame() *Game {
	// Initialize board as empty
	board := make([][]bool, BoardHeight)
	for i := range board {
		board[i] = make([]bool, BoardWidth)
	}

	return &Game{
		currentTetromino: NewTetromino(), // Spawn a random Tetromino
		board:            board,
		lastFallTime:     time.Now(),
	}
}

// Update handles game logic, including automatic falling and player input.
func (g *Game) Update() error {
	// Handle player movement
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.currentTetromino.Move(-1, 0, g.board)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.currentTetromino.Move(1, 0, g.board)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if !g.currentTetromino.Move(0, 1, g.board) {
			g.lockTetromino() // Store the Tetromino in the board and spawn a new one
		}
	}

	// Automatic falling (Gravity)
	if time.Since(g.lastFallTime) > time.Second {
		if !g.currentTetromino.Move(0, 1, g.board) {
			g.lockTetromino() // Store the Tetromino in the board and spawn a new one
		}
		g.lastFallTime = time.Now()
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

// lockTetromino places the landed Tetromino onto the board and spawns a new one.
func (g *Game) lockTetromino() {
	for _, pos := range TetrominoShapes[g.currentTetromino.shape] {
		x, y := g.currentTetromino.x+pos[0], g.currentTetromino.y+pos[1]
		if y >= 0 && y < BoardHeight && x >= 0 && x < BoardWidth {
			g.board[y][x] = true // Mark the grid as occupied
		}
	}

	g.clearFullRows()     // Clear full rows after locking a Tetromino
	g.spawnNewTetromino() // Spawn a new Tetromino
}

// spawnNewTetromino creates a new random Tetromino and checks for game over.
func (g *Game) spawnNewTetromino() {
	newTetromino := NewTetromino()

	// Check if the new Tetromino collides immediately (Game Over)
	for _, pos := range TetrominoShapes[newTetromino.shape] {
		x, y := newTetromino.x+pos[0], newTetromino.y+pos[1]
		if y >= 0 && g.board[y][x] {
			// TODO: Implement Game Over handling
			return
		}
	}

	g.currentTetromino = newTetromino
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw landed blocks
	for y, row := range g.board {
		for x, occupied := range row {
			if occupied {
				drawCell(screen, x, y, color.RGBA{255, 255, 255, 255}) // White blocks for landed Tetrominoes
			}
		}
	}

	// Draw active Tetromino
	if g.currentTetromino != nil {
		g.currentTetromino.Draw(screen)
	}
}

// Layout sets the screen size for the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return BoardWidth * CellSize, BoardHeight * CellSize
}

// clearFullRows removes full rows and shifts everything down.
func (g *Game) clearFullRows() {
	newBoard := make([][]bool, BoardHeight)
	for i := range newBoard {
		newBoard[i] = make([]bool, BoardWidth)
	}

	rowIndex := BoardHeight - 1 // Start from the bottom row

	// Copy only non-full rows to the new board
	for y := BoardHeight - 1; y >= 0; y-- {
		isFull := true
		for x := 0; x < BoardWidth; x++ {
			if !g.board[y][x] {
				isFull = false
				break
			}
		}

		if !isFull { // Keep non-full rows
			newBoard[rowIndex] = g.board[y]
			rowIndex--
		}
	}

	g.board = newBoard // Replace old board with the new one
}
