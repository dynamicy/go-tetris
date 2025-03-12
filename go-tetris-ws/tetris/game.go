package tetris

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"time"
)

// Constants for the game board
const (
	BoardWidth      = 10                     // Number of columns
	BoardHeight     = 20                     // Number of rows
	CellSize        = 24                     // Size of each cell in pixels
	MoveInterval    = 150 * time.Millisecond // Delay between left/right movements
	GravityInterval = 1 * time.Second        // Delay for automatic falling
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

// lockTetromino places the landed Tetromino onto the board and spawns a new one.
func (g *Game) lockTetromino() {
	for _, pos := range TetrominoRotations[g.currentTetromino.shape][g.currentTetromino.rotationState] {
		x, y := g.currentTetromino.x+pos[0], g.currentTetromino.y+pos[1]
		if y >= 0 && y < BoardHeight && x >= 0 && x < BoardWidth {
			g.board[y][x] = true // Correctly mark occupied blocks
		}
	}

	g.clearFullRows()     // Clear full rows
	g.spawnNewTetromino() // Spawn a new Tetromino at the top
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

// Draw renders the game screen, including the score and game over message.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw landed blocks
	for y, row := range g.board {
		for x, occupied := range row {
			if occupied {
				drawCell(screen, x, y, color.RGBA{255, 255, 255, 255}) // White blocks
			}
		}
	}

	// Draw active Tetromino
	if g.currentTetromino != nil {
		g.currentTetromino.Draw(screen)
	}

	// Display Score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, basicfont.Face7x13, 10, 20, color.White)

	// Display Game Over Message
	if g.gameOver {
		text.Draw(screen, "GAME OVER", basicfont.Face7x13, 100, 200, color.RGBA{255, 0, 0, 255})
	}
}

// clearFullRows removes full rows, shifts everything down, and updates the score.
func (g *Game) clearFullRows() {
	newBoard := make([][]bool, BoardHeight)
	for i := range newBoard {
		newBoard[i] = make([]bool, BoardWidth)
	}

	rowIndex := BoardHeight - 1 // Start from the bottom row
	rowsCleared := 0

	// Copy only non-full rows to the new board
	for y := BoardHeight - 1; y >= 0; y-- {
		isFull := true
		for x := 0; x < BoardWidth; x++ {
			if !g.board[y][x] {
				isFull = false
				break
			}
		}

		if isFull {
			rowsCleared++
		} else {
			newBoard[rowIndex] = g.board[y]
			rowIndex--
		}
	}

	g.board = newBoard // Replace old board with the new one

	// Update score based on rows cleared
	if rowsCleared > 0 {
		g.updateScore(rowsCleared)
	}
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

// Layout sets the screen size for the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return BoardWidth * CellSize, BoardHeight * CellSize
}

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
