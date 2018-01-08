package gomoku

import "fmt"

// Error for gomoku-specific errors
type Error struct {
	errCode int
}

func (e Error) Error() string {
	return fmt.Sprintf("%v", e.errCode)
}
