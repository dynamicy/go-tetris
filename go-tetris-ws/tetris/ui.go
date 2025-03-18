package tetris

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

// drawCell draws a single block at the specified (x, y) position.
func drawCell(screen *ebiten.Image, x, y int, col color.Color) {
	cellSize := CellSize // Use the defined CellSize constant
	cell := ebiten.NewImage(cellSize, cellSize)
	cell.Fill(col)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
	screen.DrawImage(cell, op)
}

// Modify Draw function to accept color input
func (t *Tetromino) Draw(screen *ebiten.Image, col color.RGBA) {
	for _, pos := range TetrominoRotations[t.shape][t.rotationState] {
		drawCell(screen, t.x+pos[0], t.y+pos[1], col)
	}
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

	// Draw ghost piece first (semi-transparent)
	if g.ghostTetromino != nil {
		g.ghostTetromino.Draw(screen, color.RGBA{100, 100, 100, 100}) // Faded color
	}

	// Draw actual Tetromino
	if g.currentTetromino != nil {
		g.currentTetromino.Draw(screen, color.RGBA{0, 255, 0, 255})
	}

	// Display Score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, basicfont.Face7x13, 10, 20, color.White)

	// Display Score, Level, and Lines
	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), basicfont.Face7x13, 10, 20, color.White)
	text.Draw(screen, fmt.Sprintf("Level: %d", g.level), basicfont.Face7x13, 10, 40, color.White)
	text.Draw(screen, fmt.Sprintf("Lines: %d", g.linesCleared), basicfont.Face7x13, 10, 60, color.White)

	// Display Game Over and Restart Instructions
	if g.gameOver {
		text.Draw(screen, "GAME OVER - Press 'R' to Restart", basicfont.Face7x13, 50, 200, color.RGBA{255, 0, 0, 255})
	}
}

// Layout sets the screen size for the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return BoardWidth * CellSize, BoardHeight * CellSize
}
