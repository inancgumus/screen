package screen

import (
	"os"

	terminal "golang.org/x/term"
)

// Size returns the width and height of the terminal screen
func Size() (int, int) {
	w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}
	return w, h
}
