package main

import "bufio"

func main() {
	// Writing File
	writingFile()

	// Reading
	readingFile()
}

func writingFile() {
	// Writing
	file := fileOpen("bufio.out")
	defer file.Close()

	bw := bufio.NewWriter(file)

	// Remember to Flush!
	defer bw.Flush()

	doWriteByte(bw)
	doWriteRune(bw)
	doWriteString(bw)
}

func readingFile() {
	file := openFile("reading.go")
	defer file.Close()

	br := bufio.NewReader(file)
	doPeek(br)
	doStringRead(br)
	doRuneRead(br)
	doByteRead(br)
	doLineRead(br)
}
