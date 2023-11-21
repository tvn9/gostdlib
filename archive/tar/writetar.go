package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
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

	if copied < stat.Size() {
		msg := "write %d bytes of %s, expected to write %d"
		return fmt.Errorf(msg, copied, filename, stat.Size())
	}
	return nil
}

func main() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("go.tar", flags, 0644)
	if err != nil {
		log.Fatalf("failed opening tar for writing: %s", err)
	}
	defer file.Close()

	tw := tar.NewWriter(file)
	defer tw.Close()

	for _, filename := range files {
		if err := addFile(filename, tw); err != nil {
			log.Fatalf("failed adding file %s to tar: %s", filename, err)
		}
	}
}
