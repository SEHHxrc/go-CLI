package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// get files
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Usage: cat files\nInvert each specified file line by line and write to standard output")
		os.Exit(0)
	}
	for _, file := range files {
		tac(file)
	}
}

func tac(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	buf := bufio.NewReader(file)

	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			defer file.Close()
			return
		}
		if err != nil {
			defer file.Close()
			fmt.Printf("Failed to read file: %v\n", err)
			return
		}
		// reverse lines
		defer func(s string) { // defer should not be used here
			fmt.Print(s)
		}(line)
	}
}
