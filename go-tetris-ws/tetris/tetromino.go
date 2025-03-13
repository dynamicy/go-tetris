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

// RotateClockwise rotates the Tetromino in the clockwise direction with Wall Kick.
func (t *Tetromino) RotateClockwise(board [][]bool) {
	newState := (t.rotationState + 1) % len(TetrominoRotations[t.shape])
	if newX, newY, success := t.wallKickTest(newState, board); success {
		t.rotationState = newState
		t.x, t.y = newX, newY
	}
}

// RotateCounterClockwise rotates the Tetromino counterclockwise with Wall Kick.
func (t *Tetromino) RotateCounterClockwise(board [][]bool) {
	newState := (t.rotationState - 1 + len(TetrominoRotations[t.shape])) % len(TetrominoRotations[t.shape])
	if newX, newY, success := t.wallKickTest(newState, board); success {
		t.rotationState = newState
		t.x, t.y = newX, newY
	}
}

// canRotate checks if the Tetromino can rotate without colliding.
func (t *Tetromino) canRotate(newX, newY, newState int, board [][]bool) bool {
	for _, pos := range TetrominoRotations[t.shape][newState] {
		x, y := newX+pos[0], newY+pos[1]

		// Out of bounds check
		if x < 0 || x >= BoardWidth || y >= BoardHeight {
			return false
		}

		// Collision check with existing blocks
		if y >= 0 && board[y][x] {
			return false
		}
	}
	return true // Rotation is valid
}

// wallKickTest applies correct wall kick offsets for rotation.
func (t *Tetromino) wallKickTest(newState int, board [][]bool) (int, int, bool) {
	kickType := "default"
	if t.shape == "I" {
		kickType = "I"
	}

	// Get the correct offsets
	kickOffsets := wallKickOffsets[kickType]

	for _, offset := range kickOffsets {
		newX := t.x + offset[0]
		newY := t.y + offset[1]

		// Check if the Tetromino can rotate with this offset
		if t.canRotate(newX, newY, newState, board) {
			return newX, newY, true // Return new position if rotation works
		}
	}
	return t.x, t.y, false // Rotation fails, stay in the same place
}
