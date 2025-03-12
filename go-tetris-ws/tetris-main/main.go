package main

import (
	"log"

	"github.com/dynamicy/go-tetris-ws/go-tetris" // Import the Tetris game package
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Initialize the Tetris game
	game := tetris.NewTetrisGame() // Method renamed to `NewTetrisGame()`

	// Set up the game window
	ebiten.SetWindowSize(320, 480)
	ebiten.SetWindowTitle("Tetris in Go")

	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
