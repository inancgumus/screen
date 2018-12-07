// +build !windows

package screen

import "fmt"

// Clear clears the screen and moves the cursor to the top-left
// corner of the screen
func Clear() {
	fmt.Println("\033[2J")
}

// MoveTopLeft moves the cursor to the top left position of the screen
func MoveTopLeft() {
	fmt.Println("\033[H")
}
