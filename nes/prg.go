package nes

import (
	"errors"
)

// PrgRom represents NES ROM PRG.
// It holds an internal buffer and an index.
type PrgRom struct {
	bytes []byte
	index int
}

func newPrgRom(buffer []byte, startIndex int) *PrgRom {
	return &PrgRom{bytes: buffer, index: startIndex}
}

// ReadNesPrgRom returns the PRG ROM of an iNES ROM.
// See https://wiki.nesdev.com/w/index.php/INES#iNES_file_format
func ReadNesPrgRom(rom []byte) *PrgRom {
	if !IsNesFile(rom) {
		panic("Not an iNES file!")
	}
	prgRomStartIndex := 16 // Header size
	if rom[6]&0b00000100 != 0 {
		prgRomStartIndex += 512 // Trainer size
	}
	prgRomSize := int(rom[4]) * 16384
	return newPrgRom(rom[prgRomStartIndex:prgRomSize], 0)
}

// ReadNes2PrgRom returns the PRG ROM of a NES 2.0 ROM.
// See https://wiki.nesdev.com/w/index.php/NES_2.0#PRG-ROM_Area
func ReadNes2PrgRom(rom []byte) *PrgRom {
	if !IsNes2File(rom) || len(rom) < 10 {
		panic("Not a NES 2.0 file!")
	}
	prgRomStartIndex := 16 // Header size
	// If bit 2 of Header byte 6 is set, trainer size is 512 bytes
	if rom[6]&0b00000100 != 0 {
		prgRomStartIndex += 512
	}
	prgRomSize := int(rom[4]) + (int(rom[9]&0b00001111) << 8)
	if len(rom) < prgRomStartIndex+prgRomSize {
		panic("Invalid ROM length")
	}
	return newPrgRom(rom[prgRomStartIndex:prgRomSize], 0)
}

func (prg *PrgRom) next() (byte, error) {
	if prg.index == len(prg.bytes) {
		return 0, errors.New("Reached end of buffer")
	}
	nextByte := prg.bytes[prg.index]
	prg.index++
	return nextByte, nil
}

// Size method returns the size of the internal buffer
// of `prg`.
func (prg PrgRom) Size() int {
	return len(prg.bytes)
}
