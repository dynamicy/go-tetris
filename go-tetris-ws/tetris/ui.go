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

// Draw renders the Tetromino correctly after rotation.
func (t *Tetromino) Draw(screen *ebiten.Image) {
	for _, pos := range TetrominoRotations[t.shape][t.rotationState] {
		drawCell(screen, t.x+pos[0], t.y+pos[1], color.RGBA{0, 255, 0, 255}) // Green block
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

// Layout sets the screen size for the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return BoardWidth * CellSize, BoardHeight * CellSize
}
