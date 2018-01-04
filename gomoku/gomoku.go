package main

import (
	"fmt"

	"gitlab.com/kennedy.dane/GoMoku.git/game"
)

type goMoku struct {
	theBoard int
}

type goMokuMove struct {
	x, y int
}

func (gomoku *goMoku) String() string {
	return "Woot"
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
}
