//go:build windows
// +build windows

package screen

import (
	"os"
	"syscall"
	"unsafe"
)

// Clear clears the screen
func Clear() {
	var (
		cursor coord
		w      dword
		h      = getScreen()
	)

	total := dword(h.size.x * h.size.y)

	xFillConsoleOutputCharacter.Call(
		uintptr(h.handle),
		uintptr(' '),
		uintptr(total),
		*(*uintptr)(unsafe.Pointer(&cursor)),
		uintptr(unsafe.Pointer(&w)),
	)

	xFillConsoleOutputAttribute.Call(
		uintptr(h.handle),
		uintptr(h.attributes),
		uintptr(total), *(*uintptr)(unsafe.Pointer(&cursor)),
		uintptr(unsafe.Pointer(&w)),
	)
}

// MoveTopLeft moves the cursor to the top left position of the screen
func MoveTopLeft() {
	h := getScreen()

	xSetConsoleCursorPosition.Call(
		uintptr(h.handle),
		*(*uintptr)(unsafe.Pointer(&coord{})),
	)
}

// MoveCursor moves the cursor anywhere on the screen
// returns outOfRange error if (x, y) is bigger than screen.Size()
func MoveCursor(x, y uint16) error {
	if a, b := Size(); int(x) > a || int(y) > b {
		return outOfRange{}
	}

	h := getScreen()

	xSetConsoleCursorPosition.Call(
		uintptr(h.handle),
		*(*uintptr)(unsafe.Pointer(&coord{x, y})),
	)

	return nil
}

func getScreen() consoleScreenBufferInfoHandle {
	h := consoleScreenBufferInfoHandle{
		handle: syscall.Handle(os.Stdout.Fd()),
	}

	xGetConsoleScreenBufferInfo.Call(
		uintptr(h.handle),
		uintptr(unsafe.Pointer(&h.consoleScreenBufferInfo)),
	)

	return h
}

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	xGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	xSetConsoleCursorPosition   = kernel32.NewProc("SetConsoleCursorPosition")
	xFillConsoleOutputCharacter = kernel32.NewProc("FillConsoleOutputCharacterW")
	xFillConsoleOutputAttribute = kernel32.NewProc("FillConsoleOutputAttribute")
)

type (
	// wchar uint16
	short int16
	dword uint32
	word  uint16
)

type coord struct {
	x uint16
	y uint16
}

type smallRect struct {
	left   short
	top    short
	right  short
	bottom short
}

type consoleScreenBufferInfo struct {
	size              coord
	cursorPosition    coord
	attributes        word
	window            smallRect
	maximumWindowSize coord
}

type consoleScreenBufferInfoHandle struct {
	handle syscall.Handle
	consoleScreenBufferInfo
}

type outOfRange struct{}

func (o outOfRange) Error() string {
	return "(x, y) out of range"
}
