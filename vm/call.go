// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// CALL_R calls a function at a register
func (c *CPU) CALL_R() {
	// Read the operands
	src := c.Fetch()

	addr := c.Registers[src]

	// Push the return address
	c.Push(c.RIP)

	// Push the base pointer
	c.Push(c.Registers[RBP])

	// Set base pointer to stack pointer
	c.Registers[RBP] = c.Registers[RSP] + 8

	// Set the instruction pointer to the address
	c.RIP = addr
}

// CALL_I calls a function at an immediate
func (c *CPU) CALL_I() {
	// Read the operands
	imm := c.Fetch()

	// Push the return address
	c.Push(c.RIP)

	// Push the base pointer
	c.Push(c.Registers[RBP])

	// Set base pointer to stack pointer
	c.Registers[RBP] = c.Registers[RSP] + 8

	// Set the instruction pointer to the address
	c.RIP = imm
}

// RET returns from a function
func (c *CPU) RET() {
	// Pop the base pointer
	c.Registers[RBP] = c.Pop()

	// Pop the return address
	c.RIP = c.Pop()
}
