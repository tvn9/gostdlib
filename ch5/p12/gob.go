// Serializing objects to binary format
// Besides the well known JSON and XML, Go also offers the binary format, gob.
// This example code goes through the basic concept of how to use gob package.

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type User struct {
	Fname  string
	Lname  string
	Age    int
	Active bool
}

func (u User) String() string {
	return fmt.Sprintf(`{Fname":%s, "Lname":%s, "Age":%d, "Active":%v}`, u.Fname, u.Lname, u.Age, u.Active)
}

type SimpleUser struct {
	Fname string
	Lname string
}

func (u SimpleUser) String() string {
	return fmt.Sprintf(`{"Fname":%s, "Lname":%s}`, u.Fname, u.Lname)
}

func main() {
	var buff bytes.Buffer

	// Encode value
	enc := gob.NewEncoder(&buff)
	user := User{
		"Radomir",
		"Sohlich",
		30,
		true,
	}

	enc.Encode(user)
	fmt.Printf("%X\n", buff.Bytes())

	// Decode value
	out := User{}
	dec := gob.NewDecoder(&buff)
	dec.Decode(&out)
	fmt.Println(out.String())

	enc.Encode(user)
	out2 := SimpleUser{}
	dec.Decode(&out2)
	fmt.Println(out2.String())
}
