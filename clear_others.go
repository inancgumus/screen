//go:build !windows
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

// MoveCursor moves the cursor anywhere on the screen
// returns outOfRange error if (x, y) is bigger than screen.Size()
func MoveCursor(x, y uint16) error {
	if a, b := Size(); int(x) > a || int(y) > b {
		return outOfRange{x, y}
	}

	fmt.Printf("\033[%d;%dH", x, y)
}
