package encryptLR

import (
	"testing"
)

func TestDecodeLR(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		{"RRL=R", "012001"},
	}

	for _, tt := range tests {
		result := decodeLR(tt.input)
		if result != tt.expected {
			t.Errorf("decodeLR(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}
