package storage

import (
	"errors"
	"strings"
)

// Storage Define "Storage" type to store keys and values
type Storage map[string]string

// New Initialize new Storage
func New() Storage {
	return Storage{}
}

// Set a new key and value in storage
func (s Storage) Set(key, value string) Storage {
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)
	s[key] = value

	return Storage{key: value}
}

// Get a value from storage by given key
func (s Storage) Get(key string) (Storage, error) {
	key = strings.TrimSpace(key)
	value, ok := s[key]
	if !ok {
		return Storage{}, errors.New("this key is not exist in storage")
	}

	return Storage{key: value}, nil
}

// Flush all data in storage
func (s Storage) Flush() {
	for k := range s {
		delete(s, k)
	}
}
