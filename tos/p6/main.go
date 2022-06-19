package main

import (
	"fmt"
	"os"
)

func main() {
	mapper := func(str string) string {
		switch str {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}
		return ""
	}
	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))
}
