package main

import (
	"mbarrera/tic-tac-toe/game"
)

// main is the entry point of the program.
func main() {
	g := game.NewGame() // Create a new instance of the game.

	// Game loop for the Tic Tac Toe game.
	for {
		g.PrintBoard() // Display the current state of the board.

		// Check if the game is over, and if so, break the loop.
		if g.GameOver() {
			break
		}

		g.HumanMove() // Handle the human player's move.

		// Check if the game is over after the human player's move.
		if g.GameOver() {
			break
		}

		g.ComputerMove() // Handle the computer player's move.
	}
}
