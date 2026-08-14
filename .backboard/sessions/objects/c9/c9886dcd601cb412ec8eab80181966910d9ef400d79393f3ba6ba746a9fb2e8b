package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	// Read from stdin if no arguments
	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printErr(err)
			os.Exit(1)
		}
		return
	}

	// Read each file
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printErr(err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			printErr(err)
			file.Close()
			os.Exit(1)
		}

		file.Close()
	}
}