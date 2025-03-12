package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
	"time"
)

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
	shape string // The type of Tetromino (T, L, Z, etc.)
	x, y  int    // The Tetromino's position on the grid
}

// NewTetromino creates and returns a new Tetromino of a random shape.
func NewTetromino() *Tetromino {
	rand.Seed(time.Now().UnixNano()) // Ensure different random shapes per run
	shapes := []string{"T", "L", "J", "I", "O", "S", "Z"}
	randomShape := shapes[rand.Intn(len(shapes))]

	return &Tetromino{
		shape: randomShape,
		x:     BoardWidth / 2, // Spawn in the center
		y:     0,              // Start at the top
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

// Draw renders the Tetromino on the game screen.
func (t *Tetromino) Draw(screen *ebiten.Image) {
	for _, pos := range TetrominoShapes[t.shape] {
		drawCell(screen, t.x+pos[0], t.y+pos[1], color.RGBA{0, 255, 0, 255}) // Green block
	}
}

// Move updates the Tetromino's position but prevents it from colliding with landed blocks.
func (t *Tetromino) Move(dx, dy int, board [][]bool) bool {
	for _, pos := range TetrominoShapes[t.shape] {
		newX, newY := t.x+pos[0]+dx, t.y+pos[1]+dy

		// Check if moving out of bounds
		if newX < 0 || newX >= BoardWidth || newY >= BoardHeight {
			return false // Collision detected
		}

		// Check if moving into an occupied cell
		if newY >= 0 && board[newY][newX] {
			return false // Collision with landed block
		}
	}

	// Move is valid
	t.x += dx
	t.y += dy
	return true
}
