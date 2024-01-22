package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err == nil {
		io.Copy(os.Stdout, file)
	} else {
		fmt.Println("File not found")
		os.Exit(1)
	}
}
