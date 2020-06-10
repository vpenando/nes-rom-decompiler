package nes

import (
	"testing"
)

func TestIsNesFile(t *testing.T) {
	validBytes := []byte("NES\x1A")
	invalidBytes := []byte("NES")

	if !IsNesFile(validBytes) {
		t.Errorf("Not a NES file")
	}

	if IsNesFile(invalidBytes) {
		t.Errorf("Not supposed to be a NES file")
	}
}

func TestIsNes2File(t *testing.T) {
	validBytes := []byte("NES\x1A\x00\x00\x00\x08")
	invalidBytes := []byte("NES\x1A\x00\x00\x00\x00")

	if !IsNes2File(validBytes) {
		t.Errorf("Not a NES 2.0 file")
	}

	if IsNes2File(invalidBytes) {
		t.Errorf("Not supposed to be a NES 2.0 file")
	}
}
