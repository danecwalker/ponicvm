// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package devices

import "fmt"

// ScreenDevice is a device that represents the screen
type ScreenDevice struct {
	// The screen width
	Width uint64
	// The screen height
	Height uint64
}

// NewScreenDevice creates a new screen device
func NewScreenDevice(width, height uint64) *ScreenDevice {
	return &ScreenDevice{
		Width:  width,
		Height: height,
	}
}

// ReadB reads a byte from the screen
func (s *ScreenDevice) ReadB(addr uint64) byte {
	return 0
}

// ReadW reads a word from the screen
func (s *ScreenDevice) ReadW(addr uint64) uint16 {
	return 0
}

// ReadL reads a double word from the screen
func (s *ScreenDevice) ReadL(addr uint64) uint32 {
	return 0
}

// ReadQ reads a quad word from the screen
func (s *ScreenDevice) ReadQ(addr uint64) uint64 {
	return 0
}

// WriteB writes a byte to the screen
func (s *ScreenDevice) WriteB(addr uint64, val byte)   {}
func (s *ScreenDevice) WriteW(addr uint64, val uint16) {}
func (s *ScreenDevice) WriteL(addr uint64, val uint32) {}
func (s *ScreenDevice) WriteQ(addr uint64, val uint64) {
	// Get the x and y coordinates
	x := addr%s.Width + 1
	y := addr/s.Height + 1

	// moveto x, y
	s.moveTo(x, y)
	// println(x, y)

	// Print the ASCII character
	s.convertToASCII(val)
}

// moveTo moves the cursor to the given x and y coordinates
func (s *ScreenDevice) moveTo(x, y uint64) {
	// Print the escape sequence
	print("\033[")
	print(y)
	print(";")
	print(x)
	print("H")
}

// convertToASCII converts a byte to an ASCII character
func (s *ScreenDevice) convertToASCII(b uint64) {
	// Check if the byte is a printable character
	fmt.Printf("%c", b)
}

// MSize returns the size of the screen device
func (s *ScreenDevice) MSize() uint64 {
	return s.Width * s.Height * 8
}
