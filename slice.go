package ts

import (
	"sync"
)

// Slice wraps a standard slice and a sync.Mutex.
type Slice struct {
	elements []interface{}
	mutex    sync.RWMutex
}

// Add locks the mutex, adds the passed elements from the underlying slice and finally unlocks the mutex.
func (s *Slice) Add(elements ...interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.elements = append(s.elements, elements...)
}

// Remove locks the mutex, removes the passed elements from the underlying slice and finally unlocks the mutex.
func (s *Slice) Remove(elements ...interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, element := range elements {
		for i, e := range s.elements {
			if element == e {
				s.elements[i] = s.elements[len(s.elements)-1]
				s.elements = s.elements[:len(s.elements)-1]
			}
		}
	}
}

func (s *Slice) GetElement(index int) interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.elements[index]
}

func (s *Slice) GetElements() []interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.elements
}

// Length returns the length of the underlying slice of elements and finally unlocks the mutex.
func (s *Slice) Length() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.elements)
}
