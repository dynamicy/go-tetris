package tetris

// Define possible rotations for each Tetromino type
var TetrominoRotations = map[string][][][]int{
	"T": {
		{{-1, 0}, {0, 0}, {1, 0}, {0, -1}}, // Default
		{{0, -1}, {0, 0}, {0, 1}, {1, 0}},  // 90°
		{{-1, 0}, {0, 0}, {1, 0}, {0, 1}},  // 180°
		{{0, -1}, {0, 0}, {0, 1}, {-1, 0}}, // 270°
	},
	"L": {
		{{-1, 0}, {0, 0}, {1, 0}, {1, -1}},  // Default (L shape)
		{{0, -1}, {0, 0}, {0, 1}, {1, 1}},   // 90° rotation (corrected)
		{{-1, 1}, {-1, 0}, {0, 0}, {1, 0}},  // 180° rotation (corrected)
		{{-1, -1}, {0, -1}, {0, 0}, {0, 1}}, // 270° rotation (corrected)
	},
	"J": {
		{{-1, 0}, {0, 0}, {1, 0}, {-1, -1}}, // Default (J shape)
		{{-1, -1}, {0, -1}, {0, 0}, {0, 1}}, // 90° rotation (corrected)
		{{1, 1}, {-1, 0}, {0, 0}, {1, 0}},   // 180° rotation (corrected)
		{{-1, 1}, {0, 1}, {0, 0}, {0, -1}},  // 270° rotation (corrected)
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
