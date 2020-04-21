package nes

import (
	"errors"
)

// HexBuffer is a raw buffer containing bytes
type HexBuffer []byte

// RomPrg represents NES ROM PRG
// It holds an internal buffer and an index
type RomPrg struct {
	hexBuffer HexBuffer
	index int
}

// Next returns the next byte of a PRG ROM, if any.
// If we already have reached the end of the buffer,
// the returned error is not nil.
func (prg *RomPrg) Next() (byte, error) {
	if prg.index == len(prg.hexBuffer) {
		return 0, errors.New("Reached end of buffer")
	}
	next := prg.hexBuffer[prg.index]
	prg.index++
	return next, nil
}