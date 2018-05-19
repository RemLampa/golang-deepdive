package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Length:", len(bs), lw)

	return len(bs), nil
}
