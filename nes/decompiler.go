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
		// We have reached the end of the PRG ROM
		return ""
	}
	var stringInst string
	switch inst {
	// ADC
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

	// AND
	case AndImmediate:
		stringInst = fmt.Sprintf("AND %s", ByteToImmediateValue(nextByte(prg)))
	case AndZeroPage:
		stringInst = fmt.Sprintf("AND %s", ByteToZeroPageAddress(nextByte(prg)))
	case AndZeroPageX:
		stringInst = fmt.Sprintf("AND %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case AndAbsolute:
		stringInst = fmt.Sprintf("AND %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AndAbsoluteX:
		stringInst = fmt.Sprintf("AND %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AndAbsoluteY:
		stringInst = fmt.Sprintf("AND %s,Y", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AndIndirectX:
		stringInst = fmt.Sprintf("AND (%s,X)", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AndIndirectY:
		stringInst = fmt.Sprintf("AND (%s,Y)", BytesToAddress(nextByte(prg), nextByte(prg)))

	// ASL
	case AslImmediate:
		stringInst = fmt.Sprintf("ASL %s", ByteToImmediateValue(nextByte(prg)))
	case AslZeroPage:
		stringInst = fmt.Sprintf("ASL %s", ByteToZeroPageAddress(nextByte(prg)))
	case AslZeroPageX:
		stringInst = fmt.Sprintf("ASL %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case AslAbsolute:
		stringInst = fmt.Sprintf("ASL %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case AslAbsoluteX:
		stringInst = fmt.Sprintf("ASL %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))

	// BIT
	case BitZeroPage:
		stringInst = fmt.Sprintf("BIT %s", ByteToZeroPageAddress(nextByte(prg)))
	case BitAbsolute:
		stringInst = fmt.Sprintf("BIT %s", BytesToAddress(nextByte(prg), nextByte(prg)))

	}
	
	return fmt.Sprintf("%s\n%s", stringInst, Decompile(prg))
}