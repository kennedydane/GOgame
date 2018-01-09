package gomoku

import (
	"bytes"
	"fmt"
	"strings"
)

// Game struct is for storing all the game-state information we need
type Game struct {
	theBoard             [][]int
	moves                []Move
	rows, columns        int
	playerOne, playerTwo *Player
}

func columnHeaderString(columns int) string {
	cols := make([]string, columns)
	for i := 0; i < columns; i++ {
		cols[i] = string('A' + i)
	}
	return strings.Join(cols, " ")
}

func gomokuBoardRowString(theRow []int) string {
	cols := make([]string, len(theRow))
	for i := 0; i < len(theRow); i++ {
		switch who := theRow[i]; who {
		case 0:
			cols[i] = "."
		case 1:
			cols[i] = "X"
		case 2:
			cols[i] = "O"
		}
	}
	return strings.Join(cols, " ")
}

// String is a stringer for the current game state
func (gomoku *Game) String() string {
	var buffer bytes.Buffer

	for row := 0; row < gomoku.rows; row++ {
		if row == 0 {
			buffer.WriteString("   ")
			buffer.WriteString(columnHeaderString(gomoku.columns))
			buffer.WriteString("\n")
		}
		buffer.WriteString(fmt.Sprintf("%2d ", row+1))
		buffer.WriteString(gomokuBoardRowString(gomoku.theBoard[row]))
		buffer.WriteString(fmt.Sprintf(" %d", row+1))

		buffer.WriteString("\n")
		if row == gomoku.rows-1 {
			buffer.WriteString("   ")
			buffer.WriteString(columnHeaderString(gomoku.columns))
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

// New method for Game to set up a new game
func (gomoku *Game) New(rows, columns int, playerOne, playerTwo *Player) {

	theBoard := make([][]int, rows, rows) // maximum board size will be 19x19
	for i := range theBoard {
		theBoard[i] = make([]int, columns, columns)
		for j := range theBoard[i] {
			theBoard[i][j] = 0
		}
	}
	gomoku.theBoard = theBoard[:columns][:rows]
	gomoku.rows = rows
	gomoku.columns = columns
	fmt.Println(cap(gomoku.theBoard), len(gomoku.theBoard))
	gomoku.playerOne = playerOne
	gomoku.playerTwo = playerOne
	gomoku.moves = make([]Move, 0, 10)
}

// Move method to add a new move to the game state
func (gomoku *Game) Move(move *Move) error {
	return nil
}

// MoveList returns a list of the move made during a game
func (gomoku *Game) MoveList() interface{} {
	return "12345"
}
