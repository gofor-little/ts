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
		for i, e := range s.Elements {
			if element == e {
				s.Elements[i] = s.Elements[len(s.Elements)-1]
				s.Elements = s.Elements[:len(s.Elements)-1]
			}
		}
	}
}

// Length returns the length of the underlying slice of elements.
func (s *Slice) Length() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.Elements)
}
