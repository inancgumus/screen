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

		fmt.Println(time.Now())

		time.Sleep(time.Second)
	}
}
