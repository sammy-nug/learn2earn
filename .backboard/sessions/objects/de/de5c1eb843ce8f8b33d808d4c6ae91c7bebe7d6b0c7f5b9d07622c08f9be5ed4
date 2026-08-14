package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// captureOutput intercepts stdout and returns the captured string
func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	f()

	// Close the writer and restore original stdout
	w.Close()
	os.Stdout = oldStdout

	// Read everything written to the pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestMainProgram(t *testing.T) {
	// Save the original arguments so we don't break other tests
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	t.Run("Missing Arguments", func(t *testing.T) {
		os.Args = []string{"program"}
		out := captureOutput(main)
		expected := "Usage: go run . \"text\"\n"
		if out != expected {
			t.Errorf("Expected %q, got %q", expected, out)
		}
	})

	t.Run("Too Many Arguments", func(t *testing.T) {
		os.Args = []string{"program", "arg1", "arg2"}
		out := captureOutput(main)
		expected := "Usage: go run . \"text\"\n"
		if out != expected {
			t.Errorf("Expected %q, got %q", expected, out)
		}
	})

	t.Run("Missing Shadow File", func(t *testing.T) {
		// Ensure file does not exist (we hide it temporarily if it does)
		os.Rename("shadow.txt", "shadow_backup.txt")
		defer os.Rename("shadow_backup.txt", "shadow.txt")

		os.Args = []string{"program", "hello"}
		out := captureOutput(main)
		
		// The error string hard-coded in the provided main() function
		expected := "Error: could not read font file 'standard.txt'\n"
		if out != expected {
			t.Errorf("Expected %q, got %q", expected, out)
		}
	})

	t.Run("Valid Input with Mock", func(t *testing.T) {
		// Mock a shadow.txt file
		// Requires 94 characters × 9 lines + 1 = 847 lines or more.
		// We'll generate 855 lines.
		var sb strings.Builder
		for i := 0; i < 855; i++ {
			if i%9 == 0 {
				sb.WriteString("\n") // separator
			} else {
				sb.WriteString("mock\n")
			}
		}
		
		err := os.WriteFile("shadow.txt", []byte(sb.String()), 0644)
		if err != nil {
			t.Fatalf("Could not create mock shadow.txt: %v", err)
		}
		defer os.Remove("shadow.txt")

		os.Args = []string{"program", "a"}
		out := captureOutput(main)
		
		// Character 'a' has index 97-32 = 65
		// For 8 rows, it prints "mock" each time.
		expected := strings.Repeat("mock\n", 8)
		if out != expected {
			t.Errorf("Expected: \n%v\nGot: \n%v", expected, out)
		}
	})

	t.Run("Incomplete Shadow File", func(t *testing.T) {
		// Create a small shadow.txt
		err := os.WriteFile("shadow.txt", []byte("line1\nline2\n"), 0644)
		if err != nil {
			t.Fatalf("Could not create short shadow.txt: %v", err)
		}
		defer os.Remove("shadow.txt")

		os.Args = []string{"program", "z"}
		out := captureOutput(main)
		
		expected := "Error: font file appears to be incomplete or malformed\n"
		if out != expected {
			t.Errorf("Expected %q, got %q", expected, out)
		}
	})
}
