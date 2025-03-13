package tetris

import "time"

// Constants for the game board
const (
	BoardWidth      = 10                     // Number of columns
	BoardHeight     = 20                     // Number of rows
	CellSize        = 24                     // Size of each cell in pixels
	MoveInterval    = 150 * time.Millisecond // Delay between left/right movements
	GravityInterval = 1 * time.Second        // Delay for automatic falling
)
