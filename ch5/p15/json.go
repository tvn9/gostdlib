// Extrating data from and incomplete JSON array
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const js = `
[
	{
		"name":"Alex",
		"lname":"Jone"
	},
	{
		"name":"Mike",
		"lname":"Hale"
	},
	{
		"name":"Tim",
		"lname":"Burton"
`

type User struct {
	Name  string `json:"name"`
	Lname string `json:"lname"`
}

func main() {
	userSlice := make([]User, 10)

	r := strings.NewReader(js)
	dec := json.NewDecoder(r)

	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}

		if tok == nil {
			break
		}

		switch tp := tok.(type) {
		case json.Delim:
			str := tp.String()
			if str == "[" || str == "{" {
				for dec.More() {
					u := User{}
					err := dec.Decode(&u)
					if err == nil {
						userSlice = append(userSlice, u)
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Println(userSlice)
}
