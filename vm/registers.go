// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// Register is a 64-bit representation of a register
type Register uint64

// Registers
const (
	RSP Register = iota
	RBP
	RAX
	RCX
	RDX
	RBX
	RSI
	RDI
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
)

// String returns the string representation of a register
func (r Register) String() string {
	switch r {
	case RAX:
		return "RAX"
	case RCX:
		return "RCX"
	case RDX:
		return "RDX"
	case RBX:
		return "RBX"
	case RSI:
		return "RSI"
	case RDI:
		return "RDI"
	case RSP:
		return "RSP"
	case RBP:
		return "RBP"
	case R8:
		return "R8"
	case R9:
		return "R9"
	case R10:
		return "R10"
	case R11:
		return "R11"
	case R12:
		return "R12"
	case R13:
		return "R13"
	case R14:
		return "R14"
	case R15:
		return "R15"
	}

	return "R??"
}

// u64 returns the uint64 representation of a register
func (r Register) U64() uint64 {
	return uint64(r)
}
