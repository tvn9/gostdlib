// Sorting slices
// sorting of data is a very common task. Go standard library simplifies the sorting by
// the sort package. This example gives a brief look at how to use it.
package main

import (
	"fmt"
	"sort"
)

type Gopher struct {
	Name string
	Age  int
}

var data = []Gopher{
	{"Daniel", 25},
	{"Tom", 19},
	{"Murthy", 33},
}

type Gophers []Gopher

func (g Gophers) Len() int {
	return len(g)
}

func (g Gophers) Less(i, j int) bool {
	return g[i].Age > g[j].Age
}

func (g Gophers) Swap(i, j int) {
	tmp := g[j]
	g[j] = g[i]
	g[i] = tmp
}

func main() {
	sort.Slice(data, func(i, j int) bool {
		return sort.StringsAreSorted([]string{data[i].Name, data[j].Name})
	})

	fmt.Printf("Sorted by name: %v\n", data)

	gophers := Gophers(data)
	sort.Sort(gophers)

	fmt.Printf("Sorted by age: %v\n", data)
}
