package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func printFile(file *zip.File) error {
	frc, err := file.Open()
	if err != nil {
		msg := "failed opening zip entry %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer frc.Close()

	fmt.Fprintf(os.Stdout, "Contents of %s:\n", file.Name)

	copied, err := io.Copy(os.Stdout, frc)
	if err != nil {
		msg := "failed reading zip entry %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}

	if uint64(copied) != file.UncompressedSize64 {
		msg := "read %s bytes of %v but expected to read %v bytes"
		return fmt.Errorf(msg, copied, file.UncompressedSize64)
	}
	fmt.Println()
	return nil
}
