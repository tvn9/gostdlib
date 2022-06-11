// Creating subtests
// in somme cases, it is useful to create a set of tests that could have a similar setup-up code.
// This could be done without having a separate function for each test.
package main

import (
	"fmt"
	"strconv"
	"testing"
)

var testData = []int{10, 11, 017}

func TestSampleOne(t *testing.T) {
	expected := "10"
	for _, val := range testData {
		tc := val
		t.Run(fmt.Sprintf("input = %d", tc), func(t *testing.T) {
			if expected != strconv.Itoa(tc) {
				t.Fail()
			}
		})
	}
}
