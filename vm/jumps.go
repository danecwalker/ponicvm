// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// JEQ jumps if the zero flag is set
func (c *CPU) JEQ() {
	// Read the operands
	imm := c.Fetch()

	// Jump if the zero flag is set
	if c.GetFlag(ZF) {
		c.RIP = imm
	}
}

// JNE jumps if the zero flag is not set
func (c *CPU) JNE() {
	// Read the operands
	imm := c.Fetch()

	// Jump if the zero flag is not set
	if !c.GetFlag(ZF) {
		c.RIP = imm
	}
}
