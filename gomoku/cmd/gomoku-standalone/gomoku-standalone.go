package main

import (
	"fmt"

	"github.com/kennedydane/GOgame/game"
	"github.com/kennedydane/GOgame/gomoku"
)

func main() {
	fmt.Println("moop")
	var theGame game.Game
	var gm *gomoku.Game
	theGame = gm
	fmt.Println(theGame)
	ok := theGame.Move(1)
	fmt.Println("!!!!", ok)
	var dummy gomoku.Move
	someMove := &dummy
	dummy.New(1, 10)
	ok = theGame.Move(someMove)
	fmt.Println("!!!!", ok)
	fmt.Println(string('A' + 1))
}
