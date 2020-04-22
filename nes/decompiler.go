package nes

import (
	"errors"
)

// HexBuffer is a raw buffer containing bytes
type HexBuffer []byte

// PrgRom represents NES ROM PRG
// It holds an internal buffer and an index
type PrgRom struct {
	hexBuffer HexBuffer
	index int
}

func newPrgRom(buffer []byte, startIndex int) *PrgRom {
	return &PrgRom{hexBuffer: buffer, index: startIndex}
}

func ReadNes2PrgRom(rom []byte) *PrgRom {
	prgRomStartIndex := 16 // Header size
	// If bit 2 of Header byte 6 is set, trainer size is 512 bytes
	if rom[6] & 0b00000100 != 0 {
		prgRomStartIndex += 512
	}
	prgRomSize := int(rom[4]) + (int(rom[9] & 0b00001111) << 8)
	return newPrgRom(rom[prgRomStartIndex:prgRomSize], 0)
}

// Next returns the next byte of a PRG ROM, if any.
// If we already have reached the end of the buffer,
// the returned error is not nil.
func (prg *PrgRom) Next() (byte, error) {
	if prg.index == len(prg.hexBuffer) {
		return 0, errors.New("Reached end of buffer")
	}
	next := prg.hexBuffer[prg.index]
	prg.index++
	return next, nil
}
