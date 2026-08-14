package main

import (
	"os"
)

func main() {
	path := os.Args[0]
	name := ""

	// Extract program name manually (everything after the last '/')
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			name = path[i+1:]
			break
		}
	}

	// If no '/' was found, the name is the whole path
	if name == "" {
		name = path
	}

	os.Stdout.Write([]byte(name))
	os.Stdout.Write([]byte("\n"))
}
