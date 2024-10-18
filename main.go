package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/pin/tftp/v3"
)

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	fmt.Fprintf(os.Stdout, "read: %v\n", filename)
	file, err := os.Open(fmt.Sprintf("%v/%v", "/etc/tftpgo", filename))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes sent\n", n)
	return nil
}

// writeHandler is called when client starts file upload to server
func writeHandler(filename string, wt io.WriterTo) error {
	err := fmt.Errorf("not implemented")
	fmt.Fprintf(os.Stderr, "%v\n", err)
	return err
}

func main() {
	// use nil in place of handler to disable read or write operations
	s := tftp.NewServer(readHandler, writeHandler)
	s.SetTimeout(5 * time.Second)  // optional
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}
