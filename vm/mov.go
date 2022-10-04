// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// MOV_RR moves a value from one register to another
func (c *CPU) MOV_RR() {
	// Read the operands
	dst := c.Fetch()
	src := c.Fetch()

	// Move the value
	c.Registers[dst] = c.Registers[src]
}

// MOV_RI moves a value from an immediate to a register
func (c *CPU) MOV_RI() {
	// Read the operands
	dst := c.Fetch()
	imm := c.Fetch()

	// Move the value
	c.Registers[dst] = imm
}

// MOV_RM moves a value from memory to a register
func (c *CPU) MOV_RM() {
	// Read the operands
	dst := c.Fetch()
	addr := c.Fetch()

	// Move the value
	c.Registers[dst] = c.Memory.ReadQ(addr)
}

// MOV_MR moves a value from a register to memory
func (c *CPU) MOV_MR() {
	// Read the operands
	addr := c.Fetch()
	src := c.Fetch()

	// Move the value
	c.Memory.WriteQ(addr, c.Registers[src])
}

// MOV_MI moves a value from an immediate to memory
func (c *CPU) MOV_MI() {
	// Read the operands
	addr := c.Fetch()
	imm := c.Fetch()

	// Move the value
	c.Memory.WriteQ(addr, imm)
}
