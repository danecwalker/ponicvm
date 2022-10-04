// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// Memory is a 64-bit addressable memory space
type Memory struct {
	// The memory space
	Memory []byte
	// The size of the memory space
	Size uint64
}

// NewMemory creates a new memory space
func NewMemory(size uint64) *Memory {
	return &Memory{
		Memory: make([]byte, size),
		Size:   size,
	}
}

// Size returns the size of the memory space
func (m *Memory) MSize() uint64 {
	return m.Size
}

// ReadB reads a byte from memory
func (m *Memory) ReadB(addr uint64) byte {
	return m.Memory[addr]
}

// ReadW reads a word from memory
func (m *Memory) ReadW(addr uint64) uint16 {
	return uint16(m.Memory[addr]) | uint16(m.Memory[addr+1])<<8
}

// ReadL reads a double world from memory
func (m *Memory) ReadL(addr uint64) uint32 {
	return uint32(m.Memory[addr]) | uint32(m.Memory[addr+1])<<8 | uint32(m.Memory[addr+2])<<16 | uint32(m.Memory[addr+3])<<24
}

// ReadQ reads a quad word from memory
func (m *Memory) ReadQ(addr uint64) uint64 {
	return uint64(m.Memory[addr]) | uint64(m.Memory[addr+1])<<8 | uint64(m.Memory[addr+2])<<16 | uint64(m.Memory[addr+3])<<24 | uint64(m.Memory[addr+4])<<32 | uint64(m.Memory[addr+5])<<40 | uint64(m.Memory[addr+6])<<48 | uint64(m.Memory[addr+7])<<56
}

// WriteB writes a byte to memory
func (m *Memory) WriteB(addr uint64, val byte) {
	m.Memory[addr] = val
}

// WriteW writes a word to memory
func (m *Memory) WriteW(addr uint64, val uint16) {
	m.Memory[addr] = byte(val)
	m.Memory[addr+1] = byte(val >> 8)
}

// WriteL writes a double word to memory
func (m *Memory) WriteL(addr uint64, val uint32) {
	m.Memory[addr] = byte(val)
	m.Memory[addr+1] = byte(val >> 8)
	m.Memory[addr+2] = byte(val >> 16)
	m.Memory[addr+3] = byte(val >> 24)
}

// WriteQ writes a quad word to memory
func (m *Memory) WriteQ(addr uint64, val uint64) {
	m.Memory[addr] = byte(val)
	m.Memory[addr+1] = byte(val >> 8)
	m.Memory[addr+2] = byte(val >> 16)
	m.Memory[addr+3] = byte(val >> 24)
	m.Memory[addr+4] = byte(val >> 32)
	m.Memory[addr+5] = byte(val >> 40)
	m.Memory[addr+6] = byte(val >> 48)
	m.Memory[addr+7] = byte(val >> 56)
}
