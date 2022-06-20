package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

var files = []string{"writetar.go", "readtar.go"}

func addFile(filename string, tw *tar.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed opening %s: %s", filename, err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed file stat for %s: %s", filename, err)
	}

	hdr := &tar.Header{
		ModTime: stat.ModTime(),
		Name:    filename,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		msg := "failed writing tar header for %s: %s"
		return fmt.Errorf(msg, filename, err)
	}

	copied, err := io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("failed writing %s to tar: %s", filename, err)
	}

	// Check copied, since we have the file stat with its size
	if copied < stat.Size() {
		msg := "wrote %d bytes of %s, expected to write %d"
		return fmt.Errorf(msg, copied, filename, stat.Size())
	}
	return nil
}
