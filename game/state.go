package game

import "fmt"

// GameOver checks if the game is over by evaluating if there's a winner or if the board is fully occupied.
func (g *Game) GameOver() bool {
	return g.checkWinner(Human) || g.checkWinner(Computer) || g.boardTilesOccupied()
}

// checkWinner checks if a player has won the game based on the current board configuration.
func (g *Game) checkWinner(player string) bool {
	for i := 0; i < 3; i++ {
		if g.Board[i][0] == player && g.Board[i][1] == player && g.Board[i][2] == player {
			return true // Horizontal win
		}
		if g.Board[0][i] == player && g.Board[1][i] == player && g.Board[2][i] == player {
			return true // Vertical win
		}
	}
	if g.Board[0][0] == player && g.Board[1][1] == player && g.Board[2][2] == player {
		return true // Diagonal win (top-left to bottom-right)
	}
	if g.Board[0][2] == player && g.Board[1][1] == player && g.Board[2][0] == player {
		return true // Diagonal win (top-right to bottom-left)
	}
	return false
}

// boardTilesOccupied checks if all tiles on the board are occupied, implying a tie if true.
func (g *Game) boardTilesOccupied() bool {
	for _, row := range g.Board {
		for _, cell := range row {
			if cell == Empty {
				return false // If any cell is empty, the board is not fully occupied
			}
		}
	}
	return true // All cells are occupied
}

// PrintBoard displays the current state of the board in the console.
func (g *Game) PrintBoard() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s | %s | %s\n", g.Board[i][0], g.Board[i][1], g.Board[i][2])
		if i < 2 {
			fmt.Println("---------")
		}
	}
}
