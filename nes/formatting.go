package nes

import (
	"fmt"
	"strings"
)

// ByteToHexString returns the hex equivalent
// of a given byte.
// Example: ByteToHexString(10) == "0A"
func ByteToHexString(b byte) string {
	return strings.ToUpper(fmt.Sprintf("%x", b))
}

// ByteToImmediateValue returns the 6502 equivalent
// of a literal value.
// Example: ByteToImmediateValue(10) == "#$0A"
func ByteToImmediateValue(b byte) string {
	hexByte := ByteToHexString(b)
	return fmt.Sprintf("#$%s", hexByte)
}

// ByteToZeroPageAddress turns a byte into
// a 6502 Zero Page address.
// Example: ByteToZeroPageAddress(18) == "$12"
func ByteToZeroPageAddress(b byte) string {
	address := ByteToHexString(b)
	return fmt.Sprintf("$%s", address)
}

// BytesToAddress turns a lower and upper bytes into
// a 6502 address (big endian).
// Example: BytesToAddress(52, 18) == "$1234"
func BytesToAddress(upper, lower byte) string {
	upperByte := ByteToHexString(upper)
	lowerByte := ByteToHexString(lower)
	return fmt.Sprintf("$%s%s", lowerByte, upperByte)
}