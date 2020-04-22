package nes

import (
	"fmt"
	"strings"
)

// byteToHexString returns the hex equivalent
// of a given byte.
// Example: byteToHexString(10) == "0A"
func byteToHexString(b byte) string {
	return strings.ToUpper(fmt.Sprintf("%x", b))
}

// ByteToImmediateValue returns the 6502 equivalent
// of a literal value.
// Example: ByteToImmediateValue(10) == "#$0A"
func ByteToImmediateValue(b byte) string {
	hexByte := byteToHexString(b)
	return fmt.Sprintf("#$%s", hexByte)
}

// ByteToZeroPageAddress turns a byte into
// a 6502 Zero Page address.
// Example: ByteToZeroPageAddress(18) == "$12"
func ByteToZeroPageAddress(b byte) string {
	address := byteToHexString(b)
	return fmt.Sprintf("$%s", address)
}

// BytesToAddress turns a lower and upper bytes into
// a 6502 address (big endian).
// Example: BytesToAddress(52, 18) == "$1234"
func BytesToAddress(upper, lower byte) string {
	upperByte := byteToHexString(upper)
	lowerByte := byteToHexString(lower)
	return fmt.Sprintf("$%s%s", lowerByte, upperByte)
}