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
	}
	return fmt.Sprintf("%s\n%s", stringInst, Decompile(prg))
}