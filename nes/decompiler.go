package nes

import (
	"errors"
)

// HexBuffer is a raw buffer containing bytes
type HexBuffer []byte

// RomPrg is a raw NES ROM PRG
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

func ByteToHexString(b byte) string {
	return fmt.Sprintf("%x", b)
}

func ByteToImmediateValue(b byte) string {
	hexByte := ByteToHexString(b)
	return fmt.Sprintf("#$%s", hexByte)
}

func BytesToAddress(upper, lower byte) string {
	upperByte := ByteToHexString(upper)
	lowerByte := ByteToHexString(lower)
	return fmt.Sprintf("$%s%s", lowerByte, upperByte)
}