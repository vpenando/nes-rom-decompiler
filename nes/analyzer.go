package nes

// IsNesFile returns true if the raw ROM is a NES ROM,
// i.e. `rom` starts with "NES" followed by 0x1A.
func IsNesFile(rom []byte) bool {
	return len(rom) >= 4 && string(rom[0:3]) == "NES" && rom[3] == 0x1A
}

// IsNes2File returns true if the raw ROM is a NES 2.0 ROM,
// i.e. `rom` is a NES ROM and flag 7 if correctly set.
func IsNes2File(rom []byte) bool {
	return IsNesFile(rom) && len(rom) >= 8 && (rom[7] & 0x0C) == 0x08
}