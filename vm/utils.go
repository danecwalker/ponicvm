// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

import (
	"fmt"
)

// uint64toHex converts a uint64 to a hex string
func uint64toHex(n uint64) string {
	hex := fmt.Sprintf("%016X", n)
	final := ""
	for i := 0; i < len(hex); i += 2 {
		final += fmt.Sprintf("%s ", hex[i:i+2])
	}

	return final[:len(final)-1]
}

// throw throws an error
func throw(err string) {
	// Print the error message with label ERROR in red
	fmt.Printf("\033[31mERROR\033[0m: %s\n", err)
}
