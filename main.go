package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vpenando/nes-rom-decompiler/nes"
)

var (
	inputFile  *string
	outputFile *string
)

func init() {
	inputFile = flag.String("i", "", "Input file (*.nes)")
	outputFile = flag.String("o", "", "Output file (*.s / *.asm)")
	flag.Parse()
	if !checkInputFile() {
		printUsage()
		os.Exit(0)
	}
}

func checkInputFile() bool {
	if *inputFile == "" {
		fmt.Println("Error: No input file.")
		return false
	}
	return true
}

func printUsage() {
	fmt.Println("Options:")
	pattern := "  -%s: %s"

	inputFileFlag := flag.Lookup("i")
	fmt.Println(fmt.Sprintf(pattern, inputFileFlag.Name, inputFileFlag.Usage))

	outputFileFlag := flag.Lookup("o")
	fmt.Println(fmt.Sprintf(pattern, outputFileFlag.Name, outputFileFlag.Usage))

	fmt.Println("Example:")
	fmt.Println("  ./decompiler -i XXX.nes [-o YYY.asm]")
}

func tryReadRom() []byte {
	rom, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to read '%s'. Aborting.", *inputFile))
	}
	if !nes.IsNesFile(rom) {
		panic("Not a NES ROM.")
	}
	return rom
}

func writePrg(reader *nes.PrgRomReader) error {
	asm := reader.Decompile()
	if *outputFile == "" {
		fmt.Println(asm)
		return nil
	}
	output, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer output.Close()
	output.Write([]byte(asm))
	return nil
}

func main() {
	rom := tryReadRom()
	var reader *nes.NewPrgRomReader
	if nes.IsNes2File(rom) {
		reader = nes.ReadNes2PrgRom(rom)
	} else if nes.IsNesFile(rom) {
		reader = nes.ReadNesPrgRom(rom)
	}
	writePrg(reader)
}
