// Running a code block only once
// In situatios when multiple goroutines run the same code and there is a code block that initializes,
// for example, shared resource, the Go standard library offers the solution, which will be described
// further.

package main

import (
	"fmt"
	"sync"
)

var names = []interface{}{"Alan", "Joe", "Jack", "Ben", "Ellen", "Lisa", "Carl", "Steve", "Anton", "Yo"}

type Source struct {
	m    *sync.Mutex
	o    *sync.Once
	data []interface{}
}

func (s *Source) Pop() (interface{}, error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.o.Do(func() {
		s.data = names
		fmt.Println("Data has been loaded.")
	})
	if len(s.data) > 0 {
		res := s.data[0]
		s.data = s.data[1:]
		return res, nil
	}
	return nil, fmt.Errorf("no data available")
}

func main() {
	s := &Source{&sync.Mutex{}, &sync.Once{}, nil}
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			// This code block is done only once
			if val, err := s.Pop(); err == nil {
				fmt.Printf("Pop %d returned: %s\n", idx, val)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
