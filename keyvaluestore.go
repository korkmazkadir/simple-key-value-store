package simplekeyvalue

import "sync"

// Store implements an inmemory key value store
type Store struct {
	// key-values are kept in map
	objects sync.Map
}

// NewKeyvalueStore createss an instance of KeyvalueStore
func NewKeyvalueStore() *Store {
	s := &Store{}
	return s
}

// Put puts a key value pair to the store if the key doesnot exists
func (s *Store) Put(key string, value []byte) bool {

	if _, ok := s.objects.Load(key); !ok {
		s.objects.Store(key, value)
		return true
	}

	return false
}

// Get returns the value of the key
func (s *Store) Get(key string) []byte {

	val, _ := s.objects.Load(key)

	return val.([]byte)
}

// Exists returns true if an object with a given name is available in the store
func (s *Store) Exists(key string) bool {

	if _, ok := s.objects.Load(key); ok {
		return true
	}

	return false
}
