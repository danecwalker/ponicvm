// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"time"

	"github.com/danecwalker/ponicvm/vm"
)

func main() {
	// Create a new io device
	io := vm.NewIO()

	// Create 1kb of memory
	memory := vm.NewMemory(1024)

	// Add the memory to the io device
	io.AddDevice(memory)

	// Create a new CPU
	cpu := vm.NewCPU(io)

	// Program
	program := []uint64{
		vm.NOP,

		vm.MOV_RI, vm.RAX.U64(), 2,
		vm.MOV_RI, vm.RBX.U64(), 3,
		vm.ADD_RR, vm.RAX.U64(), vm.RBX.U64(),
		vm.MOV_MR, 0x198, vm.RAX.U64(),
		vm.PUSH_I, 0x1,
		vm.PUSH_I, 0x4,
		vm.MOV_RM, vm.R8.U64(), 0x198,
		vm.POP, vm.R9.U64(),
		vm.POP, vm.R10.U64(),
		vm.ADD_RR, vm.R9.U64(), vm.R10.U64(),
		vm.ADD_RR, vm.R8.U64(), vm.R9.U64(),
		vm.MOV_RR, vm.RAX.U64(), vm.R8.U64(),
		vm.NOP,
		vm.INC_R, vm.RAX.U64(),
		vm.CMP_RI, vm.RAX.U64(), 12,
		vm.JNE, 0x108,

		vm.MOV_RI, vm.RCX.U64(), 0x6,
		vm.NOP,
		vm.SUB_RI, vm.RAX.U64(), 0x2,
		vm.LOOP, 0x160,
		vm.HLT,
	}

	// Load the program into the cpu
	cpu.LoadProgram(program)

	// Step the CPU
	for !cpu.Step() {
		cpu.Debug(false)

		// Wait for user input
		// fmt.Scanln()

		time.Sleep(200 * time.Millisecond)
	}

	// Run the CPU
	// cpu.Run()
}
