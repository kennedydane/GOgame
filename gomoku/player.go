package gomoku

import "fmt"

// Player type for keeping information we care about
type Player struct {
	firstName string
	lastName  string
	email     string
}

// New is for creating a new player and checking their info makes senses
func (player *Player) New() {

}

// Name is for printing the player's name
func (player *Player) Name() string {
	return fmt.Sprintf("%v %v", player.firstName, player.lastName)
}

// Email is for printing the player's email address
func (player *Player) Email() string {
	return fmt.Sprintf("%v", player.email)
}
