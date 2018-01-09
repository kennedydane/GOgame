package gomoku

import (
	"fmt"
)

// Move struct contains all the information we need for a move
type Move struct {
	row, column int
}

func (theMove Move) String() string {
	return fmt.Sprintf("%v %v", theMove.row, theMove.column)
}

//NewMove function to create new move
func NewMove(row, column int) *Move {
	theMove := new(Move)
	theMove.row = row
	theMove.column = column
	return theMove
}
