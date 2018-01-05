package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	if false {
		t.Fatalf("Meh")
	}
}

func TestNew(t *testing.T) {
	var (
		//		tmp       goMoku
		moop      *goMoku
		playerOne = goMokuPlayer{firstName: "Dane", lastName: "Kennedy", email: "kennedy.dane@gmail.com"}
		playerTwo = goMokuPlayer{firstName: "Mike", lastName: "McMikeFace", email: "mike.mcmikeface@gmail.com"}
	)
	moop = new(goMoku)
	moop.New(15, 15, &playerOne, &playerTwo)
	//moop = &tmp
	fmt.Printf("hmm -- %v\n", moop.theBoard)
	fmt.Println("The board...")
	fmt.Println(moop)
}
