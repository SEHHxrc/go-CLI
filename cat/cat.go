package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// get file names
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Usage: cat [files]\nRead all files and writes the result to standard output")
		os.Exit(0)
	}
	for _, file := range files {
		cat(file) // cat every single file
	}
}

func cat(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Failed to open file: %v\n", file)
	}
	fmt.Printf("%s", data)
}
