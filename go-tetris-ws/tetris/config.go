package tetris

import "time"

// Constants for the game board
const (
	BoardWidth  = 10 // Number of columns
	BoardHeight = 20 // Number of rows
	CellSize    = 30 // Size of each cell in pixels

	GravityInterval  = 1 * time.Second        // Delay for automatic falling
	InitialMoveDelay = 100 * time.Millisecond // Delay before repeat
	MoveRepeatRate   = 30 * time.Millisecond  // Faster movement after holding
)

// PointsPerLine defines the score for each number of rows cleared.
var PointsPerLine = map[int]int{
	1: 100, // Single row
	2: 300, // Double row
	3: 500, // Triple row
	4: 800, // Tetris (4 rows)
}
