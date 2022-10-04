// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// CMP_RR compares two registers
func (c *CPU) CMP_RR() {
	// Read the operands
	dst := c.Fetch()
	src := c.Fetch()

	// set the flags

	// zero flag
	if c.Registers[dst] == c.Registers[src] {
		c.SetFlag(ZF)
		c.ClearFlag(CF)
	} else {
		c.ClearFlag(ZF)
	}

	// carry flag
	if c.Registers[dst] < c.Registers[src] {
		c.SetFlag(CF)
	} else {
		c.ClearFlag(CF)
	}
}
