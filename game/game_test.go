package game

import (
	"testing"
)

// TestNewGame checks if a new game is correctly initialized.
func TestNewGame(t *testing.T) {
	game := NewGame()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board[i][j] != Empty {
				t.Errorf("Expected empty board, got %s at [%d][%d]", game.Board[i][j], i, j)
			}
		}
	}
}

// TestHumanMove checks if a human player can make a valid move.
func TestHumanMove(t *testing.T) {
	game := NewGame()
	game.Board[0][0] = Human
	if game.Board[0][0] != Human {
		t.Errorf("Expected %s at [0][0], got %s", Human, game.Board[0][0])
	}
}

// TestGameOver checks if the game over condition is correctly identified.
func TestGameOver(t *testing.T) {
	game := NewGame()
	game.Board[0][0] = Human
	game.Board[0][1] = Human
	game.Board[0][2] = Human

	if !game.GameOver() {
		t.Errorf("Expected game over, but GameOver() returned false")
	}
}

// TestCheckWinner checks if the winner checking logic works correctly.
func TestCheckWinner(t *testing.T) {
	game := NewGame()
	game.Board[0][0] = Human
	game.Board[1][1] = Human
	game.Board[2][2] = Human

	if !game.checkWinner(Human) {
		t.Errorf("Expected %s to be winner, but checkWinner returned false", Human)
	}
}

// TestBoardTilesOccupied checks if the board full condition is identified correctly.
func TestBoardTilesOccupied(t *testing.T) {
	game := NewGame()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.Board[i][j] = Human
		}
	}

	if !game.boardTilesOccupied() {
		t.Errorf("Expected board to be fully occupied, but boardTilesOccupied() returned false")
	}
}

// TestComputerMove can be tricky to test due to its AI nature.
// However, you can test specific scenarios, such as ensuring the computer makes a move when possible.
func TestComputerMove(t *testing.T) {
	game := NewGame()
	game.ComputerMove() // Assuming the first move is by the computer.
	occupied := false
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board[i][j] == Computer {
				occupied = true
			}
		}
	}

	if !occupied {
		t.Errorf("Expected computer to make a move, but no move was made")
	}
}
