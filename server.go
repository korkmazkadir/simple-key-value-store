package simplekeyvalue

import (
	"fmt"
	"log"
)

// Server implements a signle threaded RPC server for keyvalue store
type Server struct {
	store      *Store
	usedMemory int
	keyCount   int
}

// KeyValuePair defines a key value pair
type KeyValuePair struct {
	Key   string
	Value []byte
}

// NewServer creates an instance of the server
func NewServer() *Server {
	s := &Server{}
	s.store = NewKeyvalueStore()
	return s
}

// Put invokes Put method of keyvalue store
func (s *Server) Put(kvPair KeyValuePair, reply *bool) error {

	if kvPair.Key == "" || len(kvPair.Value) == 0 {
		return fmt.Errorf("key or Value can not be empty")
	}

	*reply = s.store.Put(kvPair.Key, kvPair.Value)

	if *reply {
		s.usedMemory += len(kvPair.Value)
		s.keyCount++
		log.Printf("the number of keys is %d Used memory is %d bytes\n", s.keyCount, s.usedMemory)
	}

	return nil
}

// Get invokes Get method of keyvalue store
func (s *Server) Get(key string, reply *[]byte) error {

	if key == "" {
		return fmt.Errorf("key can not be empty")
	}

	*reply = s.store.Get(key)
	return nil
}

// Exists invokes Exists method of keyvalue store
func (s *Server) Exists(key string, reply *bool) error {

	if key == "" {
		return fmt.Errorf("key can not be empty")
	}

	*reply = s.store.Exists(key)
	return nil
}
