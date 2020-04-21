package nes

const (
	// ADC

	AdcImmediate   byte = 0x69
	AdcZeroPage    byte = 0x65
	AdcZeroPageX   byte = 0x75
	AdcAbsolute    byte = 0x6D
	AdcAbsoluteX   byte = 0x7D
	AdcAbsoluteY   byte = 0x79
	AdcIndirectX   byte = 0x61
	AdcIndirectY   byte = 0x71

	// AND

	AndImmediate   byte = 0x29
	AndZeroPage    byte = 0x25
	AndZeroPageX   byte = 0x35
	AndAbsolute    byte = 0x2D
	AndAbsoluteX   byte = 0x3D
	AndAbsoluteY   byte = 0x39
	AndIndirectX   byte = 0x21
	AndIndirectY   byte = 0x31

	// ASL

	AslImmediate   byte = 0x0A
	AslZeroPage    byte = 0x06
	AslZeroPageX   byte = 0x16
	AslAbsolute    byte = 0x0E
	AslAbsoluteX   byte = 0x1E

	// BIT

	BitZeroPage    byte = 0x24
	BitAbsolute    byte = 0x2C

	// Branches

	Bpl            byte = 0x10
	Bmi            byte = 0x30
	Bvc            byte = 0x50
	Bvs            byte = 0x70
	Bcc            byte = 0x90
	Bcs            byte = 0xB0
	Bne            byte = 0xD0
	Beq            byte = 0xF0

	// BRK

	Brk            byte = 0x00

	// CMP

	CmpImmediate   byte = 0xC9
	CmpZeroPage    byte = 0xC5
	CmpZeroPageX   byte = 0xD5
	CmpAbsolute    byte = 0xCD
	CmpAbsoluteX   byte = 0xDD
	CmpAbsoluteY   byte = 0xD9
	CmpIndirectX   byte = 0xC1
	CmpIndirectY   byte = 0xD1

	// CPX

	CpxImmediate   byte = 0xE0
	CpxZeroPage    byte = 0xE4
	CpxAbsolute    byte = 0xEC

	// CPY

	CpyImmediate   byte = 0xC0
	CpyZeroPage    byte = 0xC4
	CpyAbsolute    byte = 0xCC

	// DEC

	DecZeroPage    byte = 0xC6
	DecZeroPageX   byte = 0xD6
	DecAbsolute    byte = 0xCE
	DecAbsoluteX   byte = 0xDE

	// EOR

	EorImmediate   byte = 0x49
	EorZeroPage    byte = 0x45
	EorZeroPageX   byte = 0x55
	EorAbsolute    byte = 0x4D
	EorAbsoluteX   byte = 0x5D
	EorAbsoluteY   byte = 0x59
	EorIndirectX   byte = 0x41
	EorIndirectY   byte = 0x51

	// Processor status

	Clc            byte = 0x18
	Sec            byte = 0x38
	Cli            byte = 0x58
	Sei            byte = 0x78
	Clv            byte = 0xB8
	Cld            byte = 0xD8
	Sed            byte = 0xF8

	// INC

	IncZeroPage    byte = 0xE6
	IncZeroPageX   byte = 0xF6
	IncAbsolute    byte = 0xEE
	IncAbsoluteX   byte = 0xFE
	
	// JMP

	JmpAbsolute    byte = 0x4C
	JmpIndirect    byte = 0x6C

	// JSR

	JsrAbsolute    byte = 0x20

	// LDA

	LdaImmediate   byte = 0xA9
	LdaZeroPage    byte = 0xA5
	LdaZeroPageX   byte = 0xB5
	LdaAbsolute    byte = 0xAD
	LdaAbsoluteX   byte = 0xBD
	LdaAbsoluteY   byte = 0xB9
	LdaIndirectX   byte = 0xA1
	LdaIndirectY   byte = 0xB1

	// LDX

	LdxImmediate   byte = 0xA2
	LdxZeroPage    byte = 0xA6
	LdxZeroPageY   byte = 0xB6
	LdxAbsolute    byte = 0xAE
	LdxAbsoluteY   byte = 0xBE

	// LDY

	LdyImmediate   byte = 0xA0
	LdyZeroPage    byte = 0xA4
	LdyZeroPageX   byte = 0xB4
	LdyAbsolute    byte = 0xAC
	LdyAbsoluteX   byte = 0xBC

	// LSR

	LsrAccumulator byte = 0x4A
	LsrZeroPage    byte = 0x46
	LsrZeroPageX   byte = 0x56
	LsrAbsolute    byte = 0x4E
	LsrAbsoluteX   byte = 0x5E

	// NOP

	NopImplied     byte = 0xEA

	// ORA

	OraImmediate   byte = 0x09
	OraZeroPage    byte = 0x05
	OraZeroPageX   byte = 0x15
	OraAbsolute    byte = 0x0D
	OraAbsoluteX   byte = 0x1D
	OraAbsoluteY   byte = 0x19
	OraIndirectX   byte = 0x01
	OraIndirectY   byte = 0x11

	// Register transfers

	Tax            byte = 0xAA
	Txa            byte = 0x8A
	Dex            byte = 0xCA
	Inx            byte = 0xE8
	Tay            byte = 0xA8
	Tya            byte = 0x98
	Dey            byte = 0x88
	Iny            byte = 0xC8
	
	// ROL

	RolAccumulator byte = 0x2A
	RolZeroPage    byte = 0x26
	RolZeroPageX   byte = 0x36
	RolAbsolute    byte = 0x2E
	RolAbsoluteX   byte = 0x3E

	// ROR

	RorAccumulator byte = 0x6A
	RorZeroPage    byte = 0x66
	RorZeroPageX   byte = 0x76
	RorAbsolute    byte = 0x6E
	RorAbsoluteX   byte = 0x7E

	// RTI

	RtiImplied     byte = 0x40

	// RTS

	RtsImplied     byte = 0x60

	// SBC

	SbcImmediate   byte = 0xE9
	SbcZeroPage    byte = 0xE5
	SbcZeroPageX   byte = 0xF5
	SbcAbsolute    byte = 0xED
	SbcAbsoluteX   byte = 0xFD
	SbcAbsoluteY   byte = 0xF9
	SbcIndirectX   byte = 0xE1
	SbcIndirectY   byte = 0xF1

	// STA

	StaZeroPage    byte = 0x85
	StaZeroPageX   byte = 0x95
	StaAbsolute    byte = 0x8D
	StaAbsoluteX   byte = 0x9D
	StaAbsoluteY   byte = 0x99
	StaIndirectX   byte = 0x81
	StaIndirectY   byte = 0x91

	// Stack

	Txs            byte = 0x9A
	Tsx            byte = 0xBA
	Pha            byte = 0x48
	Pla            byte = 0x68
	Php            byte = 0x08
	Plp            byte = 0x28

	// STX

	StxZeroPage    byte = 0x86
	StxZeroPageY   byte = 0x96
	StxAbsolute    byte = 0x8E

	// STY
	
	StyZero Page   byte = 0x84
	StyZeroPageX   byte = 0x94
	StyAbsolute    byte = 0x8C
)
