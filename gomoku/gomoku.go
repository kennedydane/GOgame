package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kennedydane/GOgame/game"
)

type goMokuPlayer struct {
	firstName string
	lastName  string
	email     string
}

func (player *goMokuPlayer) Name() string {
	return fmt.Sprintf("%v %v", player.firstName, player.lastName)
}

func (player *goMokuPlayer) Email() string {
	return fmt.Sprintf("%v", player.email)
}

type goMoku struct {
	theBoard             [][]int
	rows, columns        int
	playerOne, playerTwo *goMokuPlayer
}

type goMokuMove struct {
	row, column int
}

func columnHeaderString(columns int) string {
	/*var buf bytes.Buffer
	for col := 0; col < columns; col++ {
		buf.WriteString(string('A'+col) + " ")
	}
	return buf.String()*/
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

func (gomoku *goMoku) String() string {
	var buffer bytes.Buffer

	for row := 0; row < gomoku.rows; row++ {
		if row == 0 {
			buffer.WriteString(columnHeaderString(gomoku.columns))
			buffer.WriteString("\n")
		}
		buffer.WriteString(gomokuBoardRowString(gomoku.theBoard[row]))

		buffer.WriteString("\n")
		if row == gomoku.rows-1 {
			buffer.WriteString(columnHeaderString(gomoku.columns))
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
	//return "WTF"
}

func (gomoku *goMoku) New(rows, columns int, playerOne, playerTwo *goMokuPlayer) {

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
}

type gomokuError struct {
	errCode int
}

func (e gomokuError) Error() string {
	return fmt.Sprintf("%v", e.errCode)
}

func (gomoku *goMoku) Move(move interface{}) error {
	theMove, ok := move.(goMokuMove)
	fmt.Println("---", theMove)
	if !ok {
		return gomokuError{1}
	}
	return nil
}

func (gomoku *goMoku) MoveList() interface{} {
	return "12345"
}

func main() {
	fmt.Println("moop")
	var theGame game.Game
	var gm *goMoku
	theGame = gm
	fmt.Println(theGame)
	ok := theGame.Move(1)
	fmt.Println("!!!!", ok)
	someMove := goMokuMove{1, 10}
	ok = theGame.Move(someMove)
	fmt.Println("!!!!", ok)
	fmt.Println(string('A' + 1))
}
