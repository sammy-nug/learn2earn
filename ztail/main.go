package main

import "os"

func printStr(s string) {
	for i := 0; i < len(s); i++ {
		os.Stdout.Write([]byte{s[i]})
	}
}

func parseInt(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return -1
		}
		n = n*10 + int(s[i]-'0')
	}
	return n
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 || args[0] != "-c" {
		printStr("Usage: ./ztail -c <num> <file1> [file2 ...]\n")
		os.Exit(1)
	}

	count := parseInt(args[1])
	if count < 1 {
		printStr("Invalid number\n")
		os.Exit(1)
	}

	files := args[2:]
	if len(files) == 0 {
		printStr("No files provided\n")
		os.Exit(1)
	}

	exitStatus := 0
	multi := len(files) > 1

	for i, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			// print error first, no header
			printStr("open " + fname + ": no such file or directory\n")
			exitStatus = 1
			continue
		}

		// print header only if file opened successfully and multiple files
		if multi {
			if i > 0 {
				printStr("\n")
			}
			printStr("==> " + fname + " <==\n")
		}

		// read entire file
		info, _ := f.Stat()
		size := int(info.Size())
		buf := make([]byte, size)
		f.Read(buf)
		f.Close()

		start := 0
		if count < size {
			start = size - count
		}

		for j := start; j < size; j++ {
			os.Stdout.Write([]byte{buf[j]})
		}
	}

	os.Exit(exitStatus)
}
