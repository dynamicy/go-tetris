package tetris

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

// Define possible rotations for each Tetromino type
var TetrominoRotations = map[string][][][]int{
	"T": {
		{{-1, 0}, {0, 0}, {1, 0}, {0, -1}}, // Default
		{{0, -1}, {0, 0}, {0, 1}, {1, 0}},  // 90°
		{{-1, 0}, {0, 0}, {1, 0}, {0, 1}},  // 180°
		{{0, -1}, {0, 0}, {0, 1}, {-1, 0}}, // 270°
	},
	"L": {
		{{-1, 0}, {0, 0}, {1, 0}, {1, -1}},  // Default
		{{0, -1}, {0, 0}, {0, 1}, {1, 1}},   // 90°
		{{-1, 1}, {-1, 0}, {0, 0}, {1, 0}},  // 180°
		{{-1, -1}, {0, -1}, {0, 0}, {0, 1}}, // 270°
	},
	"J": {
		{{-1, 0}, {0, 0}, {1, 0}, {-1, -1}}, // Default
		{{-1, -1}, {0, -1}, {0, 0}, {0, 1}}, // 90°
		{{1, 1}, {-1, 0}, {0, 0}, {1, 0}},   // 180°
		{{-1, 1}, {0, 1}, {0, 0}, {0, -1}},  // 270°
	},
	"I": {
		{{-2, 0}, {-1, 0}, {0, 0}, {1, 0}}, // Default (Horizontal)
		{{0, -1}, {0, 0}, {0, 1}, {0, 2}},  // 90° (Vertical)
		{{-2, 0}, {-1, 0}, {0, 0}, {1, 0}}, // 180° (Horizontal)
		{{0, -1}, {0, 0}, {0, 1}, {0, 2}},  // 270° (Vertical)
	},
	"O": {
		{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, // Square block (No rotation)
	},
	"S": {
		{{-1, 0}, {0, 0}, {0, 1}, {1, 1}},
		{{0, -1}, {0, 0}, {-1, 0}, {-1, 1}},
	},
	"Z": {
		{{-1, 1}, {0, 1}, {0, 0}, {1, 0}},
		{{0, -1}, {0, 0}, {1, 0}, {1, 1}},
	},
}

// TetrominoShapes defines the structure of different Tetromino pieces.
var TetrominoShapes = map[string][][]int{
	"T": {{0, -1}, {0, 0}, {-1, 0}, {1, 0}}, // T-shape
	"L": {{0, -1}, {0, 0}, {0, 1}, {1, 1}},  // L-shape
	"J": {{0, -1}, {0, 0}, {0, 1}, {-1, 1}}, // J-shape
	"I": {{0, -2}, {0, -1}, {0, 0}, {0, 1}}, // I-shape
	"O": {{0, 0}, {1, 0}, {0, 1}, {1, 1}},   // O-shape
	"S": {{-1, 0}, {0, 0}, {0, 1}, {1, 1}},  // S-shape
	"Z": {{-1, 1}, {0, 1}, {0, 0}, {1, 0}},  // Z-shape
}

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

// drawCell draws a single block at the specified (x, y) position.
func drawCell(screen *ebiten.Image, x, y int, col color.Color) {
	cellSize := CellSize // Use the defined CellSize constant
	cell := ebiten.NewImage(cellSize, cellSize)
	cell.Fill(col)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
	screen.DrawImage(cell, op)
}

// Draw renders the Tetromino correctly after rotation.
func (t *Tetromino) Draw(screen *ebiten.Image) {
	for _, pos := range TetrominoRotations[t.shape][t.rotationState] {
		drawCell(screen, t.x+pos[0], t.y+pos[1], color.RGBA{0, 255, 0, 255}) // Green block
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
	//newState := (t.rotationState + 1) % len(TetrominoRotations[t.shape])
	//if t.canRotate(newState, board) {
	//	t.rotationState = newState
	//}
	newState := (t.rotationState + 1) % len(TetrominoRotations[t.shape])

	if t.canRotate(newState, board) {
		t.rotationState = newState
		fmt.Println("CW Rotation Successful! New Shape Positions:")
		for _, pos := range TetrominoRotations[t.shape][newState] {
			fmt.Println("Block at:", t.x+pos[0], t.y+pos[1])
		}
	} else {
		fmt.Println("CW Rotation Blocked!")
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
