package main

import (
	"encoding/hex"
	"log"
	"os"
)

func dumpFile() {
	data, err := os.ReadFile("main.go")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	dumper := hex.Dumper(os.Stdout)
	dumper.Write(data)
}

func main() {
	hero := []byte("Batman and Robin")
	log.Printf("hero: %s", hero)
	encoded := hex.EncodeToString(hero)
	log.Printf("encoded: %s", encoded)
	decoded, _ := hex.DecodeString(encoded)
	log.Printf("decoded: %s", decoded)
	dumpFile()
}
