package utils

import "testing"

func TestGenerateKey(t *testing.T) {
	length := 10
	got := GenerateKey(uint(length))

	if len(got) != length {
		t.Errorf("GenerateKey() = %d, want %d", len(got), length)
	}
}
