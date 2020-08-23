package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const col = 30
	bar := fmt.Sprintf("[%%-%vs]", col)

	for i := 0; i < col; i++ {
		screen.ClearLine()
		screen.MoveLeft()

		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}

  screen.ClearLine()
	screen.MoveLeft()
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
}
