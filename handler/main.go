package handler
package main

import "os"

func main() {
	args := os.Args[1:]

	for _, a := range args {
		os.Stdout.Write([]byte(a))
		os.Stdout.Write([]byte("\n"))
	}
}