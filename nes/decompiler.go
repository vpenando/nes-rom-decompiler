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
		stringInst = fmt.Sprintf("ADC (%s),Y", BytesToAddress(nextByte(prg), nextByte(prg)))

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
		stringInst = fmt.Sprintf("AND (%s),Y", BytesToAddress(nextByte(prg), nextByte(prg)))

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

	// Branches
	case Bpl:
		stringInst = fmt.Sprintf("BPL %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case Bmi:
		stringInst = fmt.Sprintf("BMI %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case Bvc:
		stringInst = fmt.Sprintf("BVC %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case Bvs:
		stringInst = fmt.Sprintf("BVS %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case Bcc:
		stringInst = fmt.Sprintf("BCC %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case Bcs:
		stringInst = fmt.Sprintf("BCS %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case Bne:
		stringInst = fmt.Sprintf("BNE %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case Beq:
		stringInst = fmt.Sprintf("BEQ %s", BytesToAddress(nextByte(prg), nextByte(prg)))

	// BRK
	case Brk:
		stringInst = "BRK"

	// CMP
	case CmpImmediate:
		stringInst = fmt.Sprintf("CMP %s", ByteToImmediateValue(nextByte(prg)))
	case CmpZeroPage:
		stringInst = fmt.Sprintf("CMP %s", ByteToZeroPageAddress(nextByte(prg)))
	case CmpZeroPageX:
		stringInst = fmt.Sprintf("CMP %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case CmpAbsolute:
		stringInst = fmt.Sprintf("CMP %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case CmpAbsoluteX:
		stringInst = fmt.Sprintf("CMP %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))
	case CmpAbsoluteY:
		stringInst = fmt.Sprintf("CMP %s,Y", BytesToAddress(nextByte(prg), nextByte(prg)))
	case CmpIndirectX:
		stringInst = fmt.Sprintf("CMP (%s,X)", BytesToAddress(nextByte(prg), nextByte(prg)))
	case CmpIndirectY:
		stringInst = fmt.Sprintf("CMP (%s),Y", BytesToAddress(nextByte(prg), nextByte(prg)))

	// CPX
	case CpxImmediate:
		stringInst = fmt.Sprintf("CPX %s", ByteToImmediateValue(nextByte(prg)))
	case CpxZeroPage:
		stringInst = fmt.Sprintf("CPX %s", ByteToZeroPageAddress(nextByte(prg)))
	case CpxAbsolute:
		stringInst = fmt.Sprintf("CPX %s", BytesToAddress(nextByte(prg), nextByte(prg)))

	// CPY
	case CpyImmediate:
		stringInst = fmt.Sprintf("CPY %s", ByteToImmediateValue(nextByte(prg)))
	case CpyZeroPage:
		stringInst = fmt.Sprintf("CPY %s", ByteToZeroPageAddress(nextByte(prg)))
	case CpyAbsolute:
		stringInst = fmt.Sprintf("CPY %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	
	// DEC
	case DecZeroPage:
		stringInst = fmt.Sprintf("DEC %s", ByteToZeroPageAddress(nextByte(prg)))
	case DecZeroPageX:
		stringInst = fmt.Sprintf("DEC %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case DecAbsolute:
		stringInst = fmt.Sprintf("DEC %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case DecAbsoluteX:
		stringInst = fmt.Sprintf("DEC %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))

	// EOR
	case EorImmediate:
		stringInst = fmt.Sprintf("EOR %s", ByteToImmediateValue(nextByte(prg)))
	case EorZeroPage:
		stringInst = fmt.Sprintf("EOR %s", ByteToZeroPageAddress(nextByte(prg)))
	case EorZeroPageX:
		stringInst = fmt.Sprintf("EOR %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case EorAbsolute:
		stringInst = fmt.Sprintf("EOR %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case EorAbsoluteX:
		stringInst = fmt.Sprintf("EOR %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))
	case EorAbsoluteY:
		stringInst = fmt.Sprintf("EOR %s,Y", BytesToAddress(nextByte(prg), nextByte(prg)))
	case EorIndirectX:
		stringInst = fmt.Sprintf("EOR (%s,X)", BytesToAddress(nextByte(prg), nextByte(prg)))
	case EorIndirectY:
		stringInst = fmt.Sprintf("EOR (%s),Y", BytesToAddress(nextByte(prg), nextByte(prg)))

	// Processor status
	case Clc:
		stringInst = "CLC"
	case Sec:
		stringInst = "SEC"
	case Cli:
		stringInst = "CLI"
	case Sei:
		stringInst = "SEI"
	case Clv:
		stringInst = "CLV"
	case Cld:
		stringInst = "CLD"
	case Sed:
		stringInst = "SED"

	// INC
	case IncZeroPage:
		stringInst = fmt.Sprintf("INC %s", ByteToZeroPageAddress(nextByte(prg)))
	case IncZeroPageX:
		stringInst = fmt.Sprintf("INC %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case IncAbsolute:
		stringInst = fmt.Sprintf("INC %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case IncAbsoluteX:
		stringInst = fmt.Sprintf("INC %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))

	// JMP
	case JmpAbsolute:
		stringInst = fmt.Sprintf("JMP %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case JmpIndirect:
		stringInst = fmt.Sprintf("JMP (%s)", BytesToAddress(nextByte(prg), nextByte(prg)))

	// JSR
	case JsrAbsolute:
		stringInst = fmt.Sprintf("JSR %s", BytesToAddress(nextByte(prg), nextByte(prg)))

	// LDA
	case LdaImmediate:
		stringInst = fmt.Sprintf("LDA %s", ByteToImmediateValue(nextByte(prg)))
	case LdaZeroPage:
		stringInst = fmt.Sprintf("LDA %s", ByteToZeroPageAddress(nextByte(prg)))
	case LdaZeroPageX:
		stringInst = fmt.Sprintf("LDA %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case LdaAbsolute:
		stringInst = fmt.Sprintf("LDA %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case LdaAbsoluteX:
		stringInst = fmt.Sprintf("LDA %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))
	case LdaAbsoluteY:
		stringInst = fmt.Sprintf("LDA %s,Y", BytesToAddress(nextByte(prg), nextByte(prg)))
	case LdaIndirectX:
		stringInst = fmt.Sprintf("LDA (%s,X)", BytesToAddress(nextByte(prg), nextByte(prg)))
	case LdaIndirectY:
		stringInst = fmt.Sprintf("LDA (%s),Y", BytesToAddress(nextByte(prg), nextByte(prg)))

	// LDX
	case LdxImmediate:
		stringInst = fmt.Sprintf("LDX %s", ByteToImmediateValue(nextByte(prg)))
	case LdxZeroPage:
		stringInst = fmt.Sprintf("LDX %s", ByteToZeroPageAddress(nextByte(prg)))
	case LdxZeroPageY:
		stringInst = fmt.Sprintf("LDX %s,Y", ByteToZeroPageAddress(nextByte(prg)))
	case LdxAbsolute:
		stringInst = fmt.Sprintf("LDX %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case LdxAbsoluteY:
		stringInst = fmt.Sprintf("LDX %s,Y", BytesToAddress(nextByte(prg), nextByte(prg)))

	// LDY
	case LdyImmediate:
		stringInst = fmt.Sprintf("LDY %s", ByteToImmediateValue(nextByte(prg)))
	case LdyZeroPage:
		stringInst = fmt.Sprintf("LDY %s", ByteToZeroPageAddress(nextByte(prg)))
	case LdyZeroPageX:
		stringInst = fmt.Sprintf("LDY %s,X", ByteToZeroPageAddress(nextByte(prg)))
	case LdyAbsolute:
		stringInst = fmt.Sprintf("LDY %s", BytesToAddress(nextByte(prg), nextByte(prg)))
	case LdyAbsoluteX:
		stringInst = fmt.Sprintf("LDY %s,X", BytesToAddress(nextByte(prg), nextByte(prg)))

	}
	
	return fmt.Sprintf("%s\n%s", stringInst, Decompile(prg))
}