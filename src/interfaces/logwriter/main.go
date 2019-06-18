package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// io.Copy takes the first paramter Writer interface
	// Writer interface has an implementation of Write function
	// In this program we are basically creating our own version of write
	// and adding a reciever to it. Hence by passing lw the copy function still works
	io.Copy(lw, resp.Body)
}

// Interfaces only provide guidelines for the implementation.
// It doesn't actually check anything so we can do anything inside the func
// Our responsibility to implement correctly
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
