package nes

import "fmt"

// PrgRomReader represents NES ROM PRG reader.
// It iterates over an internal buffer.
type PrgRomReader struct {
	rom   []byte
	index int
}

func NewPrgRomReader(buffer []byte) *PrgRomReader {
	return &PrgRomReader{rom: buffer, index: 0}
}

// ReadNesPrgRom returns the PRG ROM of an iNES ROM.
// See https://wiki.nesdev.com/w/index.php/INES#iNES_file_format
func ReadNesPrgRom(rom []byte) *PrgRomReader {
	if !IsNesFile(rom) {
		panic("Not an iNES file!")
	}
	prgRomStartIndex := 16 // Header size
	if rom[6]&0b00000100 != 0 {
		prgRomStartIndex += 512 // Trainer size
	}
	prgRomSize := int(rom[4]) * 16384
	prg := rom[prgRomStartIndex:prgRomSize]
	return NewPrgRomReader(prg)
}

// ReadNes2PrgRom returns the PRG ROM of a NES 2.0 ROM.
// See https://wiki.nesdev.com/w/index.php/NES_2.0#PRG-ROM_Area
func ReadNes2PrgRom(rom []byte) *PrgRomReader {
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
	prg := rom[prgRomStartIndex:prgRomSize]
	return NewPrgRomReader(prg)
}

func (reader *PrgRomReader) next() (byte, bool) {
	if reader.index == len(reader.rom) {
		return 0, false
	}
	nextByte := reader.rom[reader.index]
	reader.index++
	return nextByte, true
}

func nextByte(reader *PrgRomReader) byte {
	b, hasNext := reader.next()
	if !hasNext {
		panic("Reached end of PRG")
	}
	return b
}

// Decompile returns a raw PRG ROM's ASM content.
// Each unknown byte is written in commentary.
func (reader *PrgRomReader) Decompile() string {
	inst, hasNext := reader.next()
	if !hasNext {
		// We have reached the end of the PRG ROM
		return "; EOF"
	}
	var stringInst string
	switch inst {
	// ADC
	case AdcImmediate:
		stringInst = fmt.Sprintf("ADC %s", ByteToImmediateValue(nextByte(reader)))
	case AdcZeroPage:
		stringInst = fmt.Sprintf("ADC %s", ByteToZeroPageAddress(nextByte(reader)))
	case AdcZeroPageX:
		stringInst = fmt.Sprintf("ADC %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case AdcAbsolute:
		stringInst = fmt.Sprintf("ADC %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AdcAbsoluteX:
		stringInst = fmt.Sprintf("ADC %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AdcAbsoluteY:
		stringInst = fmt.Sprintf("ADC %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AdcIndirectX:
		stringInst = fmt.Sprintf("ADC (%s,X)", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AdcIndirectY:
		stringInst = fmt.Sprintf("ADC (%s),Y", BytesToAddress(nextByte(reader), nextByte(reader)))

	// AND
	case AndImmediate:
		stringInst = fmt.Sprintf("AND %s", ByteToImmediateValue(nextByte(reader)))
	case AndZeroPage:
		stringInst = fmt.Sprintf("AND %s", ByteToZeroPageAddress(nextByte(reader)))
	case AndZeroPageX:
		stringInst = fmt.Sprintf("AND %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case AndAbsolute:
		stringInst = fmt.Sprintf("AND %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AndAbsoluteX:
		stringInst = fmt.Sprintf("AND %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AndAbsoluteY:
		stringInst = fmt.Sprintf("AND %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AndIndirectX:
		stringInst = fmt.Sprintf("AND (%s,X)", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AndIndirectY:
		stringInst = fmt.Sprintf("AND (%s),Y", BytesToAddress(nextByte(reader), nextByte(reader)))

	// ASL
	case AslImmediate:
		stringInst = fmt.Sprintf("ASL %s", ByteToImmediateValue(nextByte(reader)))
	case AslZeroPage:
		stringInst = fmt.Sprintf("ASL %s", ByteToZeroPageAddress(nextByte(reader)))
	case AslZeroPageX:
		stringInst = fmt.Sprintf("ASL %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case AslAbsolute:
		stringInst = fmt.Sprintf("ASL %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case AslAbsoluteX:
		stringInst = fmt.Sprintf("ASL %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))

	// BIT
	case BitZeroPage:
		stringInst = fmt.Sprintf("BIT %s", ByteToZeroPageAddress(nextByte(reader)))
	case BitAbsolute:
		stringInst = fmt.Sprintf("BIT %s", BytesToAddress(nextByte(reader), nextByte(reader)))

	// Branches
	case Bpl:
		stringInst = fmt.Sprintf("BPL %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case Bmi:
		stringInst = fmt.Sprintf("BMI %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case Bvc:
		stringInst = fmt.Sprintf("BVC %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case Bvs:
		stringInst = fmt.Sprintf("BVS %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case Bcc:
		stringInst = fmt.Sprintf("BCC %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case Bcs:
		stringInst = fmt.Sprintf("BCS %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case Bne:
		stringInst = fmt.Sprintf("BNE %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case Beq:
		stringInst = fmt.Sprintf("BEQ %s", BytesToAddress(nextByte(reader), nextByte(reader)))

	// BRK
	case Brk:
		stringInst = "BRK"

	// CMP
	case CmpImmediate:
		stringInst = fmt.Sprintf("CMP %s", ByteToImmediateValue(nextByte(reader)))
	case CmpZeroPage:
		stringInst = fmt.Sprintf("CMP %s", ByteToZeroPageAddress(nextByte(reader)))
	case CmpZeroPageX:
		stringInst = fmt.Sprintf("CMP %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case CmpAbsolute:
		stringInst = fmt.Sprintf("CMP %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case CmpAbsoluteX:
		stringInst = fmt.Sprintf("CMP %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))
	case CmpAbsoluteY:
		stringInst = fmt.Sprintf("CMP %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))
	case CmpIndirectX:
		stringInst = fmt.Sprintf("CMP (%s,X)", BytesToAddress(nextByte(reader), nextByte(reader)))
	case CmpIndirectY:
		stringInst = fmt.Sprintf("CMP (%s),Y", BytesToAddress(nextByte(reader), nextByte(reader)))

	// CPX
	case CpxImmediate:
		stringInst = fmt.Sprintf("CPX %s", ByteToImmediateValue(nextByte(reader)))
	case CpxZeroPage:
		stringInst = fmt.Sprintf("CPX %s", ByteToZeroPageAddress(nextByte(reader)))
	case CpxAbsolute:
		stringInst = fmt.Sprintf("CPX %s", BytesToAddress(nextByte(reader), nextByte(reader)))

	// CPY
	case CpyImmediate:
		stringInst = fmt.Sprintf("CPY %s", ByteToImmediateValue(nextByte(reader)))
	case CpyZeroPage:
		stringInst = fmt.Sprintf("CPY %s", ByteToZeroPageAddress(nextByte(reader)))
	case CpyAbsolute:
		stringInst = fmt.Sprintf("CPY %s", BytesToAddress(nextByte(reader), nextByte(reader)))

	// DEC
	case DecZeroPage:
		stringInst = fmt.Sprintf("DEC %s", ByteToZeroPageAddress(nextByte(reader)))
	case DecZeroPageX:
		stringInst = fmt.Sprintf("DEC %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case DecAbsolute:
		stringInst = fmt.Sprintf("DEC %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case DecAbsoluteX:
		stringInst = fmt.Sprintf("DEC %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))

	// EOR
	case EorImmediate:
		stringInst = fmt.Sprintf("EOR %s", ByteToImmediateValue(nextByte(reader)))
	case EorZeroPage:
		stringInst = fmt.Sprintf("EOR %s", ByteToZeroPageAddress(nextByte(reader)))
	case EorZeroPageX:
		stringInst = fmt.Sprintf("EOR %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case EorAbsolute:
		stringInst = fmt.Sprintf("EOR %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case EorAbsoluteX:
		stringInst = fmt.Sprintf("EOR %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))
	case EorAbsoluteY:
		stringInst = fmt.Sprintf("EOR %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))
	case EorIndirectX:
		stringInst = fmt.Sprintf("EOR (%s,X)", BytesToAddress(nextByte(reader), nextByte(reader)))
	case EorIndirectY:
		stringInst = fmt.Sprintf("EOR (%s),Y", BytesToAddress(nextByte(reader), nextByte(reader)))

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
		stringInst = fmt.Sprintf("INC %s", ByteToZeroPageAddress(nextByte(reader)))
	case IncZeroPageX:
		stringInst = fmt.Sprintf("INC %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case IncAbsolute:
		stringInst = fmt.Sprintf("INC %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case IncAbsoluteX:
		stringInst = fmt.Sprintf("INC %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))

	// JMP
	case JmpAbsolute:
		stringInst = fmt.Sprintf("JMP %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case JmpIndirect:
		stringInst = fmt.Sprintf("JMP (%s)", BytesToAddress(nextByte(reader), nextByte(reader)))

	// JSR
	case JsrAbsolute:
		stringInst = fmt.Sprintf("JSR %s", BytesToAddress(nextByte(reader), nextByte(reader)))

	// LDA
	case LdaImmediate:
		stringInst = fmt.Sprintf("LDA %s", ByteToImmediateValue(nextByte(reader)))
	case LdaZeroPage:
		stringInst = fmt.Sprintf("LDA %s", ByteToZeroPageAddress(nextByte(reader)))
	case LdaZeroPageX:
		stringInst = fmt.Sprintf("LDA %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case LdaAbsolute:
		stringInst = fmt.Sprintf("LDA %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case LdaAbsoluteX:
		stringInst = fmt.Sprintf("LDA %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))
	case LdaAbsoluteY:
		stringInst = fmt.Sprintf("LDA %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))
	case LdaIndirectX:
		stringInst = fmt.Sprintf("LDA (%s,X)", BytesToAddress(nextByte(reader), nextByte(reader)))
	case LdaIndirectY:
		stringInst = fmt.Sprintf("LDA (%s),Y", BytesToAddress(nextByte(reader), nextByte(reader)))

	// LDX
	case LdxImmediate:
		stringInst = fmt.Sprintf("LDX %s", ByteToImmediateValue(nextByte(reader)))
	case LdxZeroPage:
		stringInst = fmt.Sprintf("LDX %s", ByteToZeroPageAddress(nextByte(reader)))
	case LdxZeroPageY:
		stringInst = fmt.Sprintf("LDX %s,Y", ByteToZeroPageAddress(nextByte(reader)))
	case LdxAbsolute:
		stringInst = fmt.Sprintf("LDX %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case LdxAbsoluteY:
		stringInst = fmt.Sprintf("LDX %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))

	// LDY
	case LdyImmediate:
		stringInst = fmt.Sprintf("LDY %s", ByteToImmediateValue(nextByte(reader)))
	case LdyZeroPage:
		stringInst = fmt.Sprintf("LDY %s", ByteToZeroPageAddress(nextByte(reader)))
	case LdyZeroPageX:
		stringInst = fmt.Sprintf("LDY %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case LdyAbsolute:
		stringInst = fmt.Sprintf("LDY %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case LdyAbsoluteX:
		stringInst = fmt.Sprintf("LDY %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))

	// LSR
	case LsrAccumulator:
		stringInst = "LSR A"
	case LsrZeroPage:
		stringInst = fmt.Sprintf("LSR %s", ByteToZeroPageAddress(nextByte(reader)))
	case LsrZeroPageX:
		stringInst = fmt.Sprintf("LSR %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case LsrAbsolute:
		stringInst = fmt.Sprintf("LSR %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case LsrAbsoluteX:
		stringInst = fmt.Sprintf("LSR %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))

	// NOP
	case NopImplied:
		stringInst = "NOP"

	// ORA
	case OraImmediate:
		stringInst = fmt.Sprintf("ORA %s", ByteToImmediateValue(nextByte(reader)))
	case OraZeroPage:
		stringInst = fmt.Sprintf("ORA %s", ByteToZeroPageAddress(nextByte(reader)))
	case OraZeroPageX:
		stringInst = fmt.Sprintf("ORA %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case OraAbsolute:
		stringInst = fmt.Sprintf("ORA %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case OraAbsoluteX:
		stringInst = fmt.Sprintf("ORA %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))
	case OraAbsoluteY:
		stringInst = fmt.Sprintf("ORA %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))
	case OraIndirectX:
		stringInst = fmt.Sprintf("ORA (%s,X)", BytesToAddress(nextByte(reader), nextByte(reader)))
	case OraIndirectY:
		stringInst = fmt.Sprintf("ORA (%s),Y", BytesToAddress(nextByte(reader), nextByte(reader)))

	// Register transfers
	case Tax:
		stringInst = "TAX"
	case Txa:
		stringInst = "TXA"
	case Dex:
		stringInst = "DEX"
	case Inx:
		stringInst = "INX"
	case Tay:
		stringInst = "TAY"
	case Tya:
		stringInst = "TYA"
	case Dey:
		stringInst = "DEY"
	case Iny:
		stringInst = "INY"

	// ROL
	case RolAccumulator:
		stringInst = "ROL A"
	case RolZeroPage:
		stringInst = fmt.Sprintf("ROL %s", ByteToZeroPageAddress(nextByte(reader)))
	case RolZeroPageX:
		stringInst = fmt.Sprintf("ROL %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case RolAbsolute:
		stringInst = fmt.Sprintf("ROL %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case RolAbsoluteX:
		stringInst = fmt.Sprintf("ROL %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))

	// ROR
	case RorAccumulator:
		stringInst = "ROR A"
	case RorZeroPage:
		stringInst = fmt.Sprintf("ROR %s", ByteToZeroPageAddress(nextByte(reader)))
	case RorZeroPageX:
		stringInst = fmt.Sprintf("ROR %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case RorAbsolute:
		stringInst = fmt.Sprintf("ROR %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case RorAbsoluteX:
		stringInst = fmt.Sprintf("ROR %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))

	// RTI
	case RtiImplied:
		stringInst = "RTI"

	// RTS
	case RtsImplied:
		stringInst = "RTS"

	// SBC
	case SbcImmediate:
		stringInst = fmt.Sprintf("SBC %s", ByteToImmediateValue(nextByte(reader)))
	case SbcZeroPage:
		stringInst = fmt.Sprintf("SBC %s", ByteToZeroPageAddress(nextByte(reader)))
	case SbcZeroPageX:
		stringInst = fmt.Sprintf("SBC %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case SbcAbsolute:
		stringInst = fmt.Sprintf("SBC %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case SbcAbsoluteX:
		stringInst = fmt.Sprintf("SBC %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))
	case SbcAbsoluteY:
		stringInst = fmt.Sprintf("SBC %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))
	case SbcIndirectX:
		stringInst = fmt.Sprintf("SBC (%s,X)", BytesToAddress(nextByte(reader), nextByte(reader)))
	case SbcIndirectY:
		stringInst = fmt.Sprintf("SBC (%s),Y", BytesToAddress(nextByte(reader), nextByte(reader)))

	// STA
	case StaZeroPage:
		stringInst = fmt.Sprintf("STA %s", ByteToZeroPageAddress(nextByte(reader)))
	case StaZeroPageX:
		stringInst = fmt.Sprintf("STA %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case StaAbsolute:
		stringInst = fmt.Sprintf("STA %s", BytesToAddress(nextByte(reader), nextByte(reader)))
	case StaAbsoluteX:
		stringInst = fmt.Sprintf("STA %s,X", BytesToAddress(nextByte(reader), nextByte(reader)))
	case StaAbsoluteY:
		stringInst = fmt.Sprintf("STA %s,Y", BytesToAddress(nextByte(reader), nextByte(reader)))
	case StaIndirectX:
		stringInst = fmt.Sprintf("STA (%s,X)", BytesToAddress(nextByte(reader), nextByte(reader)))
	case StaIndirectY:
		stringInst = fmt.Sprintf("STA (%s),Y", BytesToAddress(nextByte(reader), nextByte(reader)))

	// Stack
	case Txs:
		stringInst = "TXS"
	case Tsx:
		stringInst = "TSX"
	case Pha:
		stringInst = "PHA"
	case Pla:
		stringInst = "PLA"
	case Php:
		stringInst = "PHP"
	case Plp:
		stringInst = "PLP"

	// STX
	case StxZeroPage:
		stringInst = fmt.Sprintf("STX %s", ByteToZeroPageAddress(nextByte(reader)))
	case StxZeroPageY:
		stringInst = fmt.Sprintf("STX %s,Y", ByteToZeroPageAddress(nextByte(reader)))
	case StxAbsolute:
		stringInst = fmt.Sprintf("STX %s", BytesToAddress(nextByte(reader), nextByte(reader)))

	// STY
	case StyZeroPage:
		stringInst = fmt.Sprintf("STY %s", ByteToZeroPageAddress(nextByte(reader)))
	case StyZeroPageX:
		stringInst = fmt.Sprintf("STY %s,X", ByteToZeroPageAddress(nextByte(reader)))
	case StyAbsolute:
		stringInst = fmt.Sprintf("STY %s", BytesToAddress(nextByte(reader), nextByte(reader)))

	default:
		stringInst = fmt.Sprintf("; Unknown opcode %s", ByteToHexString(inst))
	}

	return fmt.Sprintf("%s\n%s", stringInst, reader.Decompile())
}
