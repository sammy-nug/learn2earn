package main

import (
	"strings"
	"testing"
)

// Helper function to check substring instead of full match
func contains(output, expected string) bool {
	return strings.Contains(output, expected)
}

func TestGenerateASCII_Basic(t *testing.T) {
	result, err := GenerateASCII("A", "thinkertoy.txt", "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result == "" {
		t.Errorf("Expected output, got empty string")
	}
}

func TestGenerateASCII_Word(t *testing.T) {
	result, err := GenerateASCII("Hello", "thinkertoy.txt", "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) == 0 {
		t.Errorf("Expected ASCII art output")
	}
}

func TestGenerateASCII_NewLine(t *testing.T) {
	result, err := GenerateASCII("Hi\nGo", "thinkertoy.txt", "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	lines := strings.Split(result, "\n")

	if len(lines) < 16 {
		t.Errorf("Expected multiple lines for newline input")
	}
}

func TestGenerateASCII_Color(t *testing.T) {
	result, err := GenerateASCII("A", "thinkertoy.txt", "red")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !contains(result, "\033[31m") {
		t.Errorf("Expected red color code in output")
	}
}

func TestGenerateASCII_InvalidColor(t *testing.T) {
	result, err := GenerateASCII("A", "thinkertoy.txt", "invalid")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if contains(result, "\033[") {
		t.Errorf("Did not expect color codes for invalid color")
	}
}

func TestGenerateASCII_EmptyInput(t *testing.T) {
	result, err := GenerateASCII("", "thinkertoy.txt", "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "" {
		t.Errorf("Expected empty output for empty input")
	}
}

func TestGenerateASCII_InvalidFont(t *testing.T) {
	_, err := GenerateASCII("Hello", "missing.txt", "")
	if err == nil {
		t.Errorf("Expected error for missing font file")
	}
}
