// Parsing large XML file effectively
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title       string  `xml:"title"`
	Author      string  `xml:"author"`
	Genre       string  `xml:"genre"`
	Price       float64 `xml:"price"`
	PubDate     string  `xml:"pubdate"`
	Descprition string  `xml:"descprition"`
}

func main() {
	f, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decode := xml.NewDecoder(f)

	// Read the book one by one
	books := make([]Book, 0)

	for {
		tok, _ := decode.Token()
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "book" {
				// Decode the element to struct
				var b Book
				decode.DecodeElement(&b, &tp)
				books = append(books, b)
			}
		}
	}
	fmt.Println(books)
}
