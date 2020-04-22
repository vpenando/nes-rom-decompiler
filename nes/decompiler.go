package nes

import (
	"fmt"
)

func Decompile(prg *PrgRom) string {
	inst, err := prg.Next()
	if err != nil {
		fmt.Println("EOF")
		return ""
	}
	var stringInst string
	switch inst {
	case AdcImmediate:
		stringInst = fmt.Sprintf("ADC %s", ByteToImmediateValue(prg.Next()))
	}
	return fmt.Sprintf("%s\n%s", stringInst, Decompile(prg))
}