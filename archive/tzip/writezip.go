package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

var files = []string{"writezip.go", "readzip.go"}

func addFile(filename string, zw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed opening %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zw.Create(filename)
	if err != nil {
		msg := "failed creating entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}
	// Not checking how many bytes copied,
	// since we don't know the file size without doing more work
	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("failed writing %s to zip: %s", filename, err)
	}
	return nil
}
