package tetris

// clearFullRows removes full rows, shifts everything down, and updates the score.
func (g *Game) clearFullRows() {
	rowsCleared := 0

	// Check rows from bottom to top
	for y := BoardHeight - 1; y >= 0; y-- {
		isFull := true
		for x := 0; x < BoardWidth; x++ {
			if !g.board[y][x] {
				isFull = false
				break
			}
		}

		// If a row is full, shift everything above it down
		if isFull {
			rowsCleared++
			for shiftY := y; shiftY > 0; shiftY-- {
				g.board[shiftY] = append([]bool(nil), g.board[shiftY-1]...) // Copy above row
			}
			g.board[0] = make([]bool, BoardWidth) // Clear the top row
			y++                                   // Re-check this row after shifting
		}
	}

	// Update score based on rows cleared
	if rowsCleared > 0 {
		g.updateScore(rowsCleared)
	}
}

// lockTetromino places the Tetromino onto the board and spawns a new one.
func (g *Game) lockTetromino() {
	for _, pos := range TetrominoRotations[g.currentTetromino.shape][g.currentTetromino.rotationState] {
		x, y := g.currentTetromino.x+pos[0], g.currentTetromino.y+pos[1]
		if x >= 0 && x < BoardWidth && y >= 0 && y < BoardHeight {
			g.board[y][x] = true
		}
	}

	g.clearFullRows()     // Remove completed rows
	g.spawnNewTetromino() // Generate new piece
}

// canMove checks if the Tetromino can move without colliding.
func (g *Game) canMove(dx, dy int) bool {
	for _, pos := range TetrominoRotations[g.currentTetromino.shape][g.currentTetromino.rotationState] {
		x, y := g.currentTetromino.x+pos[0]+dx, g.currentTetromino.y+pos[1]+dy
		if x < 0 || x >= BoardWidth || y >= BoardHeight || (y >= 0 && g.board[y][x]) {
			return false
		}
	}
	return true
}

// spawnNewTetromino creates a new Tetromino and checks for game over.
func (g *Game) spawnNewTetromino() {
	newTetromino := NewTetromino()

	// Check immediate collision for game over
	for _, pos := range TetrominoRotations[newTetromino.shape][newTetromino.rotationState] {
		x, y := newTetromino.x+pos[0], newTetromino.y+pos[1]
		if y >= 0 && g.board[y][x] {
			g.gameOver = true
			return
		}
	}

	g.currentTetromino = newTetromino
}

// updateGhostPiece calculates where the ghost piece should land
func (g *Game) updateGhostPiece() {
	if g.currentTetromino == nil {
		return
	}

	// Create a new Tetromino for the ghost piece
	g.ghostTetromino = &Tetromino{
		shape:         g.currentTetromino.shape,
		x:             g.currentTetromino.x,
		y:             g.currentTetromino.y,
		rotationState: g.currentTetromino.rotationState,
	}

	// Move the ghost piece down **until it reaches a collision**
	for g.canMoveTetromino(g.ghostTetromino, 0, 1) {
		g.ghostTetromino.y++
	}
}
