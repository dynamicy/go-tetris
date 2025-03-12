package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// TetrominoShapes defines the structure of different Tetromino pieces.
// Each shape is represented as an array of (x, y) offsets relative to its center.
var TetrominoShapes = map[string][][]int{
	"T": {
		{0, -1}, {0, 0}, {-1, 0}, {1, 0}, // T-shape
	},
	"L": {
		{0, -1}, {0, 0}, {0, 1}, {1, 1}, // L-shape
	},
}

// Tetromino represents a falling piece in the Tetris game.
type Tetromino struct {
	shape string // The type of Tetromino (T, L, Z, etc.)
	x, y  int    // The Tetromino's position on the grid
}

// NewTetromino creates and returns a new Tetromino of the specified shape.
func NewTetromino(shape string) *Tetromino {
	return &Tetromino{
		shape: shape,
		x:     5, // Default spawn position (center)
		y:     0, // Start at the top of the board
	}
}

// Draw renders the Tetromino on the game screen.
func (t *Tetromino) Draw(screen *ebiten.Image) {
	for _, pos := range TetrominoShapes[t.shape] {
		drawCell(screen, t.x+pos[0], t.y+pos[1], color.RGBA{0, 255, 0, 255}) // Green block
	}
}

// drawCell draws a single block at the specified (x, y) position.
func drawCell(screen *ebiten.Image, x, y int, col color.Color) {
	cellSize := 24
	cell := ebiten.NewImage(cellSize, cellSize)
	cell.Fill(col)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
	screen.DrawImage(cell, op)
}

// Move updates the Tetromino's position by (dx, dy).
func (t *Tetromino) Move(dx, dy int) {
	t.x += dx
	t.y += dy
}
