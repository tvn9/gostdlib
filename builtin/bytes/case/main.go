package main

import (
	"bytes"
	"log"
	"unicode"
)

func main() {
	quickBrownFox := []byte("The quick brown fox jumped over the lazy dog")

	title := bytes.ToTitle(quickBrownFox)
	log.Printf("ToTitle turned %q into %q", quickBrownFox, title)

	allTitle := bytes.ToUpper(quickBrownFox)
	log.Printf("ToUpper turned %q into %q", quickBrownFox, allTitle)

	allTitleTurkish := bytes.ToTitleSpecial(unicode.TurkishCase, quickBrownFox)
	log.Printf("ToTitleSpecial turned %q into %q using the Turkish case rules", quickBrownFox, allTitleTurkish)

	lower := bytes.ToLower(title)
	log.Printf("ToLower turned %q into %q", title, lower)

	vietNamese := []byte("Thành Phố Hồ Chí Minh")
	toLowerSpecialVN := bytes.ToLowerSpecial(unicode.SpecialCase{}, vietNamese)
	log.Printf("ToLowerSpecial turned %q into %q using the Turkish case rules", vietNamese, toLowerSpecialVN)
}
