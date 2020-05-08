package nes

import (
	"testing"
)

func TestByteToHexString(t *testing.T) {
	var expectedResults = map[byte]string{
		0: "00",   1: "01",  2: "02",  3: "03",  4: "04",  5: "05",  6: "06",  7: "07",  8: "08",  9: "09",
		10: "0A", 11: "0B", 12: "0C", 13: "0D", 14: "0E", 15: "0F", 16: "10", 17: "11", 18: "12", 19: "13",
		20: "14",
		255: "FF",
	}
	for k, v := range expectedResults {
		if ByteToHexString(k) != v {
			t.Errorf("Invalid value: expected '%s', got '%s'", v, ByteToHexString(k))
		}
	}
}