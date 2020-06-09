package nes

import (
	"testing"
)

func TestIsNesFile(t *testing.T) {
	validBytes := []byte{'N', 'E', 'S', 0x1A}
	invalidBytes := []byte{'N', 'E', 'S'}

	if !IsNesFile(validBytes) {
		t.Errorf("Not a NES file")
	}

	if IsNesFile(invalidBytes) {
		t.Errorf("Not supposed to be a NES file")
	}
}

func TestIsNes2File(t *testing.T) {
	validBytes := []byte{'N', 'E', 'S', 0x1A, 0, 0, 0, 0x08}
	invalidBytes := []byte{'N', 'E', 'S', 0x1A, 0, 0, 0, 0}

	if !IsNes2File(validBytes) {
		t.Errorf("Not a NES 2.0 file")
	}

	if IsNes2File(invalidBytes) {
		t.Errorf("Not supposed to be a NES 2.0 file")
	}
}
