// +build !windows

package screen

import "fmt"

// Clear clears the screen
func Clear() {
	fmt.Print("\033[2J")
}

// MoveTopLeft moves the cursor to the top left position of the screen
func MoveTopLeft() {
	fmt.Print("\033[H")
}
