# ASCII Art Generator

A command-line application written in Go that converts text strings into an ASCII graphical representation using a specified font file (`shadow.txt`).

## Features
- Converts standard text into large ASCII art characters.
- Supports handling literal newline sequences (`\n`).
- Includes a complete unit test suite to verify application behavior and edge cases.

## Requirements
- [Go](https://golang.org/) 1.16 or higher
- A standard `shadow.txt, standard.txt and thinkertoy.txt` font file located in the same directory as the executable.

## Usage

Run the program by supplying exactly one command-line argument:

```bash
go run . "Hello World"
```

To incorporate newlines, pass the literal `\n` sequence inside the string argument:

```bash
go run . "Line 1\nLine 2"
```

### Font File Format (`shadow.txt`)
- The font file is required to start formatting from the space character (ASCII 32) up to the `~` character (ASCII 126).
- Every character representation occupies exactly **9 lines**: 
  - 1 blank separator line 
  - 8 lines depicting the graphical ascii-art rows.

## Testing

The project is fully tested via Go's standard `testing` package. The test suite avoids relying on a real `shadow.txt` by mocking standard output streams and injecting mocked filesystem responses.

To run the test suite:

```bash
go test -v ./...
```

The testing covers branches including:
- Missing and excessive command-line arguments.
- Handling of a missing or incomplete `shadow.txt` file, standard.txt file and thinkertoy.txt
- Correct execution output with mocked font data.
- Handling escaped newlines correctly.

## Authors
- Rachael  Mfon
- Olayinka Samuel Ojo
- Samuel Emmanuel
