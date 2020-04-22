package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	_ "strings"
	_ "github.com/vpenando/nes-rom-decompiler/nes"
)

const (
	empty = ""
)

var (
	inputFile *string
)

func init() {
	fmt.Println("Parsing args...")
	inputFile = flag.String("input", empty, "The ROM")
	flag.Parse()
	if !checkInputFile() {
		printUsage()
		os.Exit(0)
	}
}

func checkInputFile() bool {
	if *inputFile == empty {
		fmt.Println("Error: No input file.")
		return false
	}
	return true
}

func printUsage() {

}

func main() {
	content, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(fmt.Sprintf("Panic: Failed to read '%s'. Aborting.", *inputFile))
	}

	fmt.Println(nes.IsNesFile(content))
}