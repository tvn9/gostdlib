package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {

	// Writezip
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("go.zip", flags, 0644)
	if err != nil {
		if err != nil {
			log.Fatalf("failed opening zip for writing: %s", err)
		}
	}
	defer file.Close()

	zw := zip.NewWriter(file)
	defer zw.Close()

	for _, filename := range files {
		if err := addFile(filename, zw); err != nil {
			log.Fatalf("failed adding file %s to zip: %s", filename, err)
		}
	}

	// ReadZip
	rc, err := zip.OpenReader("go.zip")
	if err != nil {
		msg := "failed opening archive, run `go run writezip.go` first: %s"
		log.Fatalf(msg, err)
	}
	defer rc.Close()
	for _, file := range rc.File {
		if err := printFile(file); err != nil {
			log.Fatalf("failed reading %s from zip %s", file.Name, err)
		}
	}

}
