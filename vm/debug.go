// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

import (
	"fmt"
)

// printRegister prints the registers
func (c *CPU) printRegister() {
	// Print the registers
	fmt.Println("Registers:")
	println()

	// Print the RIP in red
	fmt.Printf("\033[31m%-8s %s %8d\033[0m\n", "RIP", uint64toHex(c.RIP), c.RIP)

	for i, register := range c.Registers {
		// if register is 0, color gray
		if register == 0 {
			fmt.Printf("\033[90m%-8s %s %8d\033[0m\n", Register(i), uint64toHex(register), register)
		} else {
			// else color white
			fmt.Printf("%-8s %s %8d\n", Register(i), uint64toHex(register), register)
		}
	}
}

// printStack prints the stack
func (c *CPU) printStack() {
	// Print the stack
	fmt.Println("Stack:")
	println()

	// Get end of stack
	stackEnd := c.Registers[RSP]

	// Get start of stack
	stackStart := c.Memory.Size

	// Print the stack
	for i := stackEnd + 8; i < stackStart; i += 8 {
		// Print the stack if it is not 0
		fmt.Printf("%04X %8d %s\n", i, i, uint64toHex(c.Memory.ReadQ(i)))
	}
}

// printMemory prints the memory
func (c *CPU) printMemory(size uint64) {
	// Print the memory
	fmt.Println("Memory:")
	println()

	// Get the start of the memory
	memoryStart := c.ps

	// Get the end of the memory
	memoryEnd := c.ps + size

	// Print the memory
	for i := memoryStart; i < memoryEnd; i += 8 {
		// Print the memory
		// color yellow if value is not 0
		val := c.Memory.ReadQ(i)
		if val == 0 {
			fmt.Printf("\033[90m%04X %8d %s\033[0m\n", i, i, uint64toHex(val))
		} else {
			fmt.Printf("%04X %8d %s\n", i, i, uint64toHex(val))
		}
	}
}

// printProgram prints the program
func (c *CPU) printProgram() {
	// Print the program
	fmt.Println("Program:")
	println()

	// Get the start of the program
	programStart := uint64(0)

	// Get the end of the program
	programEnd := c.ps

	// Print the program
	// color grey if not current instruction
	for i := programStart; i < programEnd; i += 8 {
		if i == c.RIP-8 {
			fmt.Printf("%04X %8d %s\n", i, i, uint64toHex(c.Memory.ReadQ(i)))
		} else {
			fmt.Printf("\033[90m%04X %8d %s\033[0m\n", i, i, uint64toHex(c.Memory.ReadQ(i)))
		}
	}
}

// printFlags prints the flags
func (c *CPU) printFlags() {
	// Print the flags
	fmt.Println("Flags:")
	println()

	// Print the flags from FlagMap
	for i, flag := range FlagMap {
		// if flag is not set, color gray
		if !c.GetFlag(uint64(i)) {
			fmt.Printf("\033[90m%-8s %t\033[0m\n", flag, c.GetFlag(uint64(i)))
		} else {
			// else color white
			fmt.Printf("%-8s %t\n", flag, c.GetFlag(uint64(i)))
		}
	}
}

// Debug prints the CPU state
func (c *CPU) Debug(showProgram bool) {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	// Print the flags
	c.printFlags()

	println()

	// Print the registers
	c.printRegister()

	println()

	// Print the stack
	c.printStack()

	println()

	// Print the memory
	c.printMemory(128)

	if showProgram {
		println()

		// Print the program
		c.printProgram()
	}
}
