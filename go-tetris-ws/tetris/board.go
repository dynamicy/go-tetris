package tetris

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
