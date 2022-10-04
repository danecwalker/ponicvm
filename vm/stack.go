// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// Push pushes a value to the stack
func (c *CPU) Push(val uint64) {
	// Write the value to the stack
	c.Memory.WriteQ(c.Registers[RSP], val)
	// Decrement the stack pointer
	c.Registers[RSP] -= 8
}

// Pop pops a value from the stack
func (c *CPU) Pop() uint64 {
	// Increment the stack pointer
	c.Registers[RSP] += 8
	// Read the value from the stack
	return c.Memory.ReadQ(c.Registers[RSP])
}

// PUSH_R pushes a value from a register to the stack
func (c *CPU) PUSH_R() {
	// Read the operands
	src := c.Fetch()

	// Push the value
	c.Push(c.Registers[src])
}

// PUSH_I pushes a value from an immediate to the stack
func (c *CPU) PUSH_I() {
	// Read the operands
	imm := c.Fetch()

	// Push the value
	c.Push(imm)
}

// PUSH_M pushes a value from memory to the stack
func (c *CPU) PUSH_M() {
	// Read the operands
	addr := c.Fetch()

	// Push the value
	c.Push(c.Memory.ReadQ(addr))
}

// POP pops a value from the stack to a register
func (c *CPU) POP() {
	// Read the operands
	dst := c.Fetch()

	// Pop the value
	c.Registers[dst] = c.Pop()
}
