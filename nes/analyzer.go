package nes

// IsNesFile returns true if the raw ROM is a NES ROM,
// i.e. `rom` starts with "NES" followed by 0x1A.
func IsNesFile(rom []byte) bool {
	return len(rom) >= 4 && string(rom[0:3]) == "NES" && rom[3] == 0x1A
}

func IsNes2File(rom []byte) bool {
	return IsNesFile(rom) && len(rom) >= 8 (rom[7] & 0x0C) == 0x08
}