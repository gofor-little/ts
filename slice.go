package ts

import (
	"sync"
)

// Slice wraps a standard slice and a sync.Mutex.
type Slice struct {
	Elements []interface{}
	mutex    sync.Mutex
}

// Add locks the mutex, adds the passed elements from the underlying slice and finally unlocks the mutex.
func (s *Slice) Add(elements ...interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Elements = append(s.Elements, elements...)
}

// Remove locks the mutex, removes the passed elements from the underlying slice and finally unlocks the mutex.
func (s *Slice) Remove(elements ...interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, element := range elements {
		for index, e := range s.Elements {
			if element == e {
				s.Elements[index] = s.Elements[len(s.Elements)-1]
				s.Elements = s.Elements[:len(s.Elements)-1]
			}
		}
	}
}

// Exists locks the mutex, checks if the passed element exists, returning the result and unlocking the mutex.
func (s *Slice) Exists(element interface{}) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, e := range s.Elements {
		if e == element {
			return true
		}
	}

	return false
}

// Length returns the length of the underlying slice of elements.
func (s *Slice) Length() int {
	return len(s.Elements)
}
