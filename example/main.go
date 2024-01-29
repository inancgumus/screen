package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		w, h := screen.Size()
		fmt.Printf("Width: %d Height: %d\n", w, h)

		fmt.Println(time.Now())

		time.Sleep(time.Second)

		err := screen.MoveCursor(0, 0)
		if err == nil {

			time.Sleep(time.Second)
		}
	}
}
