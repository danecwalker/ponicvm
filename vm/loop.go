// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// LOOP loops a number of times and decrements the loop counter
func (cpu *CPU) LOOP() {
	// Get address to jump to
	address := cpu.Fetch()

	// Decrement loop counter
	cpu.Registers[RCX]--

	// Jump to address if loop counter is not 0
	if cpu.Registers[RCX] != 0 {
		cpu.RIP = address
	}
}
