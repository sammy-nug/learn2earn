package main

import "os"

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}

func printErr(err error) {
	os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
}
