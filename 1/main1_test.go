package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestCombineVariables(t *testing.T) {
	result := CombineVariables(42, 052, 0x2A, 3.14, "Golang", true, 1+2i)
	expected := "42 42 42 3.140000 Golang true (1+2i)"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
func TestHashWithSalt(t *testing.T) {
	input := "Hello, World!"
	salt := "go-2024"
	saltedInput := input[:len(input)/2] + salt + input[len(input)/2:]
	expectedHash := sha256.Sum256([]byte(saltedInput))
	expectedHex := hex.EncodeToString(expectedHash[:])

	result := HashWithSalt(input, salt)

	if result != expectedHex {
		t.Errorf("Expected hash %s but got %s", expectedHex, result)
	}
}
func TestPrint(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("printTypes panicked: %v", r)
		}
	}()
	print(42, 052, 0x2A, 3.14, "Golang", true, 1+2i)
}
