package tetris

import (
	"math/rand"
)

// Tetromino represents a falling piece in the Tetris game.
type Tetromino struct {
	shape         string // The type of Tetromino (T, L, Z, etc.)
	x, y          int    // The Tetromino's position on the grid
	rotationState int    // Current rotation state
}

// NewTetromino creates a new Tetromino with a random shape.
func NewTetromino() *Tetromino {
	shapes := []string{"T", "L", "J", "I", "O", "S", "Z"} // All possible shapes
	randomShape := shapes[rand.Intn(len(shapes))]         // Pick a random shape

	return &Tetromino{
		shape:         randomShape,
		x:             BoardWidth / 2,
		y:             0,
		rotationState: 0, // Default rotation state
	}
}

// Move attempts to move the Tetromino by (dx, dy). Returns false if movement is blocked.
func (t *Tetromino) Move(dx, dy int, board [][]bool) bool {
	for _, pos := range TetrominoRotations[t.shape][t.rotationState] {
		newX, newY := t.x+pos[0]+dx, t.y+pos[1]+dy

		// Check if new position is out of bounds
		if newX < 0 || newX >= BoardWidth || newY >= BoardHeight {
			return false // Block movement if it exceeds board limits
		}

		// Check collision with landed Tetrominoes
		if newY >= 0 && board[newY][newX] {
			return false
		}
	}

	// Move Tetromino only if it's valid
	t.x += dx
	t.y += dy
	return true
}

// RotateClockwise rotates the Tetromino in the clockwise direction.
func (t *Tetromino) RotateClockwise(board [][]bool) {
	newState := (t.rotationState + 1) % len(TetrominoRotations[t.shape])

	if t.canRotate(newState, board) {
		t.rotationState = newState
	}
}

// RotateCounterClockwise rotates the Tetromino counterclockwise.
func (t *Tetromino) RotateCounterClockwise(board [][]bool) {
	newState := (t.rotationState - 1 + len(TetrominoRotations[t.shape])) % len(TetrominoRotations[t.shape])
	if t.canRotate(newState, board) {
		t.rotationState = newState
	}
}

// canRotate checks if the Tetromino can rotate to a new state without colliding.
func (t *Tetromino) canRotate(newState int, board [][]bool) bool {
	for _, pos := range TetrominoRotations[t.shape][newState] {
		x, y := t.x+pos[0], t.y+pos[1]
		if x < 0 || x >= BoardWidth || y >= BoardHeight || (y >= 0 && board[y][x]) {
			return false // Collision detected
		}
	}
	return true
}
