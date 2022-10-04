// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// Opcode is a 64-bit representation of an opcode
type Opcode = uint64

// Opcodes
const (
	// NOP does nothing
	NOP Opcode = iota
	// HLT halts the CPU
	HLT

	// MOV_RR moves a value from one register to another
	MOV_RR
	// MOV_RI moves a value from an immediate to a register
	MOV_RI
	// MOV_RM moves a value from memory to a register
	MOV_RM
	// MOV_MR moves a value from a register to memory
	MOV_MR
	// MOV_MI moves a value from an immediate to memory
	MOV_MI

	// PUSH_R pushes a value from a register to the stack
	PUSH_R
	// PUSH_I pushes a value from an immediate to the stack
	PUSH_I
	// PUSH_M pushes a value from memory to the stack
	PUSH_M
	// POP pops a value from the stack to a register
	POP

	// CALL_R calls a function at a register
	CALL_R
	// CALL_I calls a function at an immediate
	CALL_I
	// RET returns from a function
	RET

	// CMP_RR compares two registers
	CMP_RR

	// JEQ jumps to address if the last comparison was equal
	JEQ
	// JNE jumps to address if the last comparison was not equal
	JNE

	// LOOP loops a number of times
	LOOP

	// DEC_R decrements a register
	DEC_R
	// INC_R increments a register
	INC_R

	// ADD_RR adds two registers
	ADD_RR
	// ADD_RI adds a immediate and a register
	ADD_RI
	// SUB_RR subtracts two registers
	SUB_RR
	// SUB_RI subtracts a immediate and a register
	SUB_RI
)
