package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		// Piscine does NOT want custom error messages
		// So just return silently.
		return
	}
	defer file.Close()

	// Copy file content to stdout
	io.Copy(os.Stdout, file)
}