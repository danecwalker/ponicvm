// Copyright (c) 2022 DevDane <dane@danecwalker.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package vm

// IO is an interface for IO devices
type IO struct {
	Regions []*Region
	Size    uint64
}

// Region is a region of memory
type Region struct {
	Start  uint64
	End    uint64
	Device Device
}

// Device is an interface for devices
type Device interface {
	// Size returns the size of the device
	MSize() uint64

	// ReadB reads a byte from the device
	ReadB(addr uint64) byte
	// ReadW reads a word from the device
	ReadW(addr uint64) uint16
	// ReadL reads a double word from the device
	ReadL(addr uint64) uint32
	// ReadQ reads a quad word from the device
	ReadQ(addr uint64) uint64

	// WriteB writes a byte to the device
	WriteB(addr uint64, val byte)
	// WriteW writes a word to the device
	WriteW(addr uint64, val uint16)
	// WriteL writes a double word to the device
	WriteL(addr uint64, val uint32)
	// WriteQ writes a quad word to the device
	WriteQ(addr uint64, val uint64)
}

// NewIO creates a new IO device
func NewIO() *IO {
	return &IO{
		Regions: make([]*Region, 0),
		Size:    0,
	}
}

// AddDevice adds a device to the IO device
func (io *IO) AddDevice(device Device) func() {

	// Add the device to the start of the regions
	r := &Region{Device: device}

	if len(io.Regions) == 0 {
		r.Start = 0
	} else {
		r.Start = io.Regions[0].End
	}

	r.End = r.Start + device.MSize()

	io.Regions = append([]*Region{r}, io.Regions...)
	io.Size += device.MSize()

	// Return a function to remove the device
	return func() {
		for i, region := range io.Regions {
			if region == r {
				io.Regions = append(io.Regions[:i], io.Regions[i+1:]...)
				break
			}
		}
	}
}

// getRegion gets the region at a given address
func (io *IO) getRegion(addr uint64) *Region {
	for _, region := range io.Regions {
		if addr >= region.Start && addr <= region.End {
			return region
		}
	}

	throw("invalid memory read")
	return nil
}

// ReadB reads a byte from the IO device
func (io *IO) ReadB(addr uint64) byte {
	// Get the region
	region := io.getRegion(addr)
	if region == nil {
		throw("invalid memory read")
		return 0
	}

	// Remap the address
	addr -= region.Start

	// Read the byte
	return region.Device.ReadB(addr)
}

// ReadW reads a word from the IO device
func (io *IO) ReadW(addr uint64) uint16 {
	// Get the region
	region := io.getRegion(addr)
	if region == nil {
		throw("invalid memory read")
		return 0
	}

	// Remap the address
	addr -= region.Start

	// Read the word
	return region.Device.ReadW(addr)
}

// ReadL reads a double word from the IO device
func (io *IO) ReadL(addr uint64) uint32 {
	// Get the region
	region := io.getRegion(addr)
	if region == nil {
		throw("invalid memory read")
		return 0
	}

	// Remap the address
	addr -= region.Start

	// Read the double word
	return region.Device.ReadL(addr)
}

// ReadQ reads a quad word from the IO device
func (io *IO) ReadQ(addr uint64) uint64 {
	// Get the region
	region := io.getRegion(addr)
	if region == nil {
		throw("invalid memory read")
		return 0
	}

	// Remap the address
	addr -= region.Start

	// Read the quad word
	return region.Device.ReadQ(addr)
}

// WriteB writes a byte to the IO device
func (io *IO) WriteB(addr uint64, val byte) {
	// Get the region
	region := io.getRegion(addr)
	if region == nil {
		throw("invalid memory write")
		return
	}

	// Remap the address
	addr -= region.Start

	// Write the byte
	region.Device.WriteB(addr, val)
}

// WriteW writes a word to the IO device
func (io *IO) WriteW(addr uint64, val uint16) {
	// Get the region
	region := io.getRegion(addr)
	if region == nil {
		throw("invalid memory write")
		return
	}

	// Remap the address
	addr -= region.Start

	// Write the word
	region.Device.WriteW(addr, val)
}

// WriteL writes a double word to the IO device
func (io *IO) WriteL(addr uint64, val uint32) {
	// Get the region
	region := io.getRegion(addr)
	if region == nil {
		throw("invalid memory write")
		return
	}

	// Remap the address
	addr -= region.Start

	// Write the double word
	region.Device.WriteL(addr, val)
}

// WriteQ writes a quad word to the IO device
func (io *IO) WriteQ(addr uint64, val uint64) {
	// Get the region
	region := io.getRegion(addr)
	if region == nil {
		throw("invalid memory write")
		return
	}

	// fmt.Printf("%x\n", region.Start)
	// Remap the address
	addr -= region.Start

	// Write the quad word
	region.Device.WriteQ(addr, val)
}
