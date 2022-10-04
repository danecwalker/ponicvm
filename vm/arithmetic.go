// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// DEC_R decrements a register
func (cpu *CPU) DEC_R() {
	// Read the operands
	dst := cpu.Fetch()

	// Decrement the register
	cpu.Registers[dst]--
}

// INC_R increments a register
func (cpu *CPU) INC_R() {
	// Read the operands
	dst := cpu.Fetch()

	// Increment the register
	cpu.Registers[dst]++
}

// ADD_RR adds two registers
func (cpu *CPU) ADD_RR() {
	// Read the operands
	dst := cpu.Fetch()
	src := cpu.Fetch()

	// Add the registers
	cpu.Registers[dst] += cpu.Registers[src]
}

// ADD_RI adds a register and an immediate
func (cpu *CPU) ADD_RI() {
	// Read the operands
	dst := cpu.Fetch()
	src := cpu.Fetch()

	// Add the register and immediate
	cpu.Registers[dst] += src
}

// SUB_RR subtracts two registers
func (cpu *CPU) SUB_RR() {
	// Read the operands
	dst := cpu.Fetch()
	src := cpu.Fetch()

	// Subtract the registers
	cpu.Registers[dst] -= cpu.Registers[src]
}

// SUB_RI subtracts a register and an immediate
func (cpu *CPU) SUB_RI() {
	// Read the operands
	dst := cpu.Fetch()
	src := cpu.Fetch()

	// Subtract the register and immediate
	cpu.Registers[dst] -= src
}
