package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	filename := os.Args[1]

	file, error := os.Open(filename)

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, file)

	error = file.Close()

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(2)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
