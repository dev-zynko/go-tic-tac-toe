package game

const (
	Empty    = " "
	Human    = "X"
	Computer = "O"
)

// Game represents the state of the Tic Tac Toe game.
type Game struct {
	Board [3][3]string // Board represents the 3x3 grid for the game.
}

// NewGame creates a new instance of the Tic Tac Toe game and initializes the board with empty tiles.
func NewGame() *Game {
	return &Game{
		Board: [3][3]string{
			{Empty, Empty, Empty},
			{Empty, Empty, Empty},
			{Empty, Empty, Empty},
		},
	}
}
