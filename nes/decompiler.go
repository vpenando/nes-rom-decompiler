package nes

import (
	"fmt"
)

func nextByte(prg *PrgRom) byte {
	b, err := prg.Next()
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()))
	}
	return b
}

func Decompile(prg *PrgRom) string {
	inst, err := prg.Next()
	if err != nil {
		fmt.Println("EOF")
		return ""
	}
	var stringInst string
	switch inst {
	case AdcImmediate:
		stringInst = fmt.Sprintf("ADC %s", ByteToImmediateValue(nextByte(prg)))
	case AdcZeroPage:
		stringInst = fmt.Sprintf("ADC %s", ByteToZeroPageAddress(nextByte(prg)))
	case AdcZeroPageX:
		stringInst = fmt.Sprintf("ADC %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case AdcAbsolute:
		stringInst = fmt.Sprintf("ADC %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AdcAbsoluteX:
		stringInst = fmt.Sprintf("ADC %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AdcAbsoluteY:
		stringInst = fmt.Sprintf("ADC %s,Y", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AdcIndirectX:
		stringInst = fmt.Sprintf("ADC (%s,X)", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AdcIndirectY:
		stringInst = fmt.Sprintf("ADC (%s,Y)", BytesToAddress(nextByte(prg), nextByte(prg)))
	}
	return fmt.Sprintf("%s\n%s", stringInst, Decompile(prg))
}