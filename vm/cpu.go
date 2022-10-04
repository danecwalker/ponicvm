// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// CPU is a 64-bit CPU
type CPU struct {
	// The memory space
	Memory *IO
	// The registers
	Registers [16]uint64
	// The instruction pointer
	RIP uint64
	// The Program Size
	ps uint64
	// Flags
	Flags uint64
}

// NewCPU creates a new CPU
func NewCPU(memory *IO) *CPU {
	c := &CPU{
		Memory: memory,
	}

	// Set the stack pointer to the end of the memory space
	c.Registers[RSP] = memory.Size - 8
	// Set the base pointer to the end of the memory space
	c.Registers[RBP] = memory.Size - 8

	return c
}

// GetFlag gets a flag at a given index
func (c *CPU) GetFlag(index uint64) bool {
	return c.Flags&(1<<index) != 0
}

// SetFlag sets a flag at a given index to true
func (c *CPU) SetFlag(index uint64) {
	c.Flags |= 1 << index
}

// Clear flag clears a flag at a given index to false
func (c *CPU) ClearFlag(index uint64) {
	c.Flags &= ^(1 << index)
}

// Fetch fetches an instruction from memory
func (c *CPU) Fetch() uint64 {
	// Fetch the instruction
	instruction := c.Memory.ReadQ(c.RIP)
	// Increment the instruction pointer
	c.RIP += 8

	return instruction
}

// Execute executes an instruction
func (c *CPU) Execute(instruction uint64) {
	switch instruction {
	case NOP:
		// Do nothing
	case HLT:
		// Halt the CPU

	// MOV_RR moves a value from one register to another
	case MOV_RR:
		c.MOV_RR()
	// MOV_RI moves a value from an immediate to a register
	case MOV_RI:
		c.MOV_RI()
	// MOV_RM moves a value from memory to a register
	case MOV_RM:
		c.MOV_RM()
	// MOV_MR moves a value from a register to memory
	case MOV_MR:
		c.MOV_MR()
	// MOV_MI moves a value from an immediate to memory
	case MOV_MI:
		c.MOV_MI()

	// PUSH_R pushes a value from a register to the stack
	case PUSH_R:
		c.PUSH_R()
	// PUSH_I pushes a value from an immediate to the stack
	case PUSH_I:
		c.PUSH_I()
	// PUSH_M pushes a value from memory to the stack
	case PUSH_M:
		c.PUSH_M()
	// POP pops a value from the stack to a register
	case POP:
		c.POP()

	// CALL_R calls a function at a register
	case CALL_R:
		c.CALL_R()
	// CALL_I calls a function at an immediate
	case CALL_I:
		c.CALL_I()
	// RET returns from a function
	case RET:
		c.RET()

	// CMP_RR compares two registers
	case CMP_RR:
		c.CMP_RR()

	// JEQ jumps to address if the last comparison was equal
	case JEQ:
		c.JEQ()
	// JNE jumps to address if the last comparison was not equal
	case JNE:
		c.JNE()

	// LOOP loops a given number of times
	case LOOP:
		c.LOOP()

	// DEC_R decrements a register
	case DEC_R:
		c.DEC_R()
	// INC_R increments a register
	case INC_R:
		c.INC_R()

	// ADD_RR adds two registers
	case ADD_RR:
		c.ADD_RR()
	// ADD_RI adds a register and an immediate
	case ADD_RI:
		c.ADD_RI()
	// SUB_RR subtracts two registers
	case SUB_RR:
		c.SUB_RR()
	// SUB_RI subtracts a register and an immediate
	case SUB_RI:
		c.SUB_RI()
	}
}

// Step steps the CPU
func (c *CPU) Step() bool {
	// Fetch the instruction
	instruction := c.Fetch()
	// Execute the instruction
	c.Execute(instruction)
	return instruction == HLT
}

// LoadProgram loads a program into memory
func (c *CPU) LoadProgram(program []uint64) {
	// Set the program size
	c.ps = uint64(len(program) * 8)

	// Write the program to memory
	for i, instruction := range program {
		c.Memory.WriteQ(uint64(i)*8, instruction)
	}
}

// Run runs the CPU
func (c *CPU) Run() {
	// Run the CPU until it halts
	for !c.Step() {
	}
}
