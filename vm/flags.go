// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// Flags indexes
const (
	ZF = iota
	CF
)

// FlagMap is a map of flags to their indexes
var FlagMap = [...]string{
	ZF: "ZF",
	CF: "CF",
}
