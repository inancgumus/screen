# Screen

The screen package provides an easy way for clearing the screen, moving the cursor and getting the size of the current window in a cross-platform way (Linux, OS X and Windows).

## Installation

```
$ go get -u github.com/inancgumus/screen
```

## Clearing the Screen

You can clear the screen and move the cursor to the top-left corner of the screen. This is good enough to create an animated program (such as an always updating clock or a progress bar).

```go
package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// Clear all the characters on the screen
	screen.Clear()

	for {
		// Moves the cursor to the top-left position of the screen
		screen.MoveTopLeft()

		// Animate the time always in the same position
		fmt.Println(time.Now())

		time.Sleep(time.Second)
	}
}
```

## Getting the Width and Height

You can get the current terminal width and height. Actually, you don't have to use this package to do that. It's just a simple wrapper over [terminal](https://godoc.org/golang.org/x/crypto/ssh/terminal) package.

```go
package main

import (
	"fmt"
	"github.com/inancgumus/screen"
)

func main() {
	w, h := screen.Size()
    
	fmt.Printf("Width: %d Height: %d\n", w, h)
}
```

---

_MIT License_