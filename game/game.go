//Package game is used to define basic game elements -- this is so that more games can be added in the future. Maybe.
package game

// Player is a generic struct to store player information
type Player struct {
	FirstName string
	LastName  string
	Email     string
}

// Game is a generic interface for games
type Game interface {
	String() string
	Move(move interface{}) error
	MoveList() interface{}
}
