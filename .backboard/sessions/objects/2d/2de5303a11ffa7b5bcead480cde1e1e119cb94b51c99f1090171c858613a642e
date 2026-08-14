package main

import "testing"

func TestAsciiRender(t *testing.T) {

	banner, _ := LoadBanner("standard.txt")

	lines := ParseInput("A")

	if len(banner) == 0 {
		t.Error("Banner failed to load")
	}

	if len(lines) != 1 {
		t.Error("Input parsing failed")
	}
}


func TestLoadBanner(t *testing.T) {

	banner, err := LoadBanner("standard.txt")

	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	if len(banner) == 0 {
		t.Error("Banner should not be empty")
	}
}

func TestParseInput(t *testing.T) {

	result := ParseInput("Hello\\nWorld")

	if len(result) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(result))
	}

	if result[0] != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", result[0])
	}

	if result[1] != "World" {
		t.Errorf("Expected 'World', got '%s'", result[1])
	}
}

func TestParseEmptyInput(t *testing.T) {

	result := ParseInput("")

	if len(result) != 1 {
		t.Errorf("Expected 1 empty line, got %d", len(result))
	}
}

func TestMultipleNewLines(t *testing.T) {

	result := ParseInput("Hello\\n\\nWorld")

	if len(result) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(result))
	}

	if result[1] != "" {
		t.Error("Expected empty line between Hello and World")
	}
}

func TestIsPrintable(t *testing.T) {

	if !IsPrintable('A') {
		t.Error("'A' should be printable")
	}

	if !IsPrintable(' ') {
		t.Error("Space should be printable")
	}

	if IsPrintable(31) {
		t.Error("ASCII 31 should not be printable")
	}

	if IsPrintable(127) {
		t.Error("ASCII 127 should not be printable")
	}
}

func TestGetCharIndex(t *testing.T) {

	index := GetCharIndex('A')

	expected := (65-32)*9 + 1

	if index != expected {
		t.Errorf("Expected %d, got %d", expected, index)
	}
}

func TestBannerStructure(t *testing.T) {

	banner, err := LoadBanner("standard.txt")

	if err != nil {
		t.Fatal(err)
	}

	if len(banner)%9 != 0 {
		t.Error("Banner structure invalid")
	}
}