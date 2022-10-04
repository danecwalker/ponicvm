// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/danecwalker/ponicvm/devices"
	"github.com/danecwalker/ponicvm/vm"
)

func main() {
	// Create a new io device
	io := vm.NewIO()

	// Create 1kb of memory
	memory := vm.NewMemory(352)

	// Create a screen device
	screen := devices.NewScreenDevice(80, 13)

	// Add the memory to the io device
	io.AddDevice(memory)
	io.AddDevice(screen)

	// Create a new CPU
	cpu := vm.NewCPU(io)

	// Program
	program := []uint64{
		vm.NOP,

		vm.MOV_MI, 0x160, 0x48,
		vm.MOV_MI, 0x161, 0x65,
		vm.MOV_MI, 0x162, 0x6c,
		vm.MOV_MI, 0x163, 0x6c,
		vm.MOV_MI, 0x164, 0x6f,
		vm.MOV_MI, 0x165, 0x20,
		vm.MOV_MI, 0x166, 0x57,
		vm.MOV_MI, 0x167, 0x6f,
		vm.MOV_MI, 0x168, 0x72,
		vm.MOV_MI, 0x169, 0x6c,
		vm.MOV_MI, 0x16a, 0x64,
		vm.MOV_MI, 0x16b, 0x21,
		vm.MOV_MI, 0x16c, 0x0a,

		vm.HLT,
	}

	// Load the program into the cpu
	cpu.LoadProgram(program)

	// // Step the CPU
	// for !cpu.Step() {
	// 	cpu.Debug(false)

	// 	// Wait for user input
	// 	fmt.Scanln()
	// }

	// Run the CPU
	cpu.Run()
}
