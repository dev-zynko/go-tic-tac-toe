package game

import (
	"fmt"
	"math"
)

// HumanMove handles the player's move, prompting for the row and column, validating the move, and updating the board.
func (g *Game) HumanMove() {
	var row, col int
	fmt.Println("Enter row (0-2) and column (0-2) for your move (e.g., 1 1 for middle): ")
	fmt.Scanln(&row, &col)

	if g.Board[row][col] != Empty {
		fmt.Println("Invalid move. Try again.")
		g.HumanMove()
		return
	}
	g.Board[row][col] = Human
}

// ComputerMove calculates the best move for the computer using the minimax algorithm and updates the board.
func (g *Game) ComputerMove() {
	bestScore := math.Inf(-1)
	var moveRow, moveCol int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == Empty {
				g.Board[i][j] = Computer
				score := g.minimax(0, false)
				g.Board[i][j] = Empty
				if score > bestScore {
					bestScore = score
					moveRow = i
					moveCol = j
				}
			}
		}
	}

	g.Board[moveRow][moveCol] = Computer
	fmt.Println("AI's move:")
}

// minimax implements the minimax algorithm for optimal decision-making in the game.
func (g *Game) minimax(depth int, isMaximizing bool) float64 {
	if g.GameOver() {
		if g.checkWinner(Computer) {
			return 1
		} else if g.checkWinner(Human) {
			return -1
		}
		return 0
	}

	if isMaximizing {
		bestScore := math.Inf(-1)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if g.Board[i][j] == Empty {
					g.Board[i][j] = Computer
					score := g.minimax(depth+1, false)
					g.Board[i][j] = Empty
					bestScore = math.Max(bestScore, score)
				}
			}
		}
		return bestScore
	}

	bestScore := math.Inf(1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == Empty {
				g.Board[i][j] = Human
				score := g.minimax(depth+1, true)
				g.Board[i][j] = Empty
				bestScore = math.Min(bestScore, score)
			}
		}
	}
	return bestScore
}
