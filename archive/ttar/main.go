package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func main() {
	writeTar()
	readTar()
}

func writeTar() {
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

func readTar() {
	file, err := os.Open("go.tar")
	if err != nil {
		msg := "failed opening archive, run `go run writetar.go` first: %s"
		log.Fatalf(msg, err)
	}
	defer file.Close()

	tr := tar.NewReader(file)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed getting next tar entry %s", err)
		}

		printHeader(hdr)
		printContents(tr, hdr.Size)
	}
}
