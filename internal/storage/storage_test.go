package storage

import (
	"reflect"
	"testing"
)

func TestStorage_Set(t *testing.T) {
	// Initialize storage
	var s Storage = Storage{}

	// Set a new key and value to the storage
	v := s.Set("foo", "bar")

	// Test ...
	if v["foo"] != "bar" {
		t.Error("Expected 'bar', got ", v["foo"])
	}
}

func TestStorage_Get(t *testing.T) {
	// Initialize storage
	var s Storage = Storage{
		"foo": "bar",
	}

	// Get value from storage
	v, _ := s.Get("foo")

	// Test return type from function
	valueType := reflect.TypeOf(v).String()
	if valueType != "storage.Storage" {
		t.Error("Expected 'storage.Storage', got ", valueType)
	}

	// Test return value from function
	if v["foo"] != "bar" {
		t.Error("Expected 'bar', got ", v["foo"])
	}
}

func TestStorage_Flush(t *testing.T) {
	// Initialize storage
	var s Storage = Storage{
		"first":  "foo",
		"second": "bar",
		"third":  "baz",
	}

	// Check map length is empty or not
	if len(s) != 3 {
		t.Error("Expected '3', got ", len(s))
	}

	s.Flush()

	// Check map length is empty or not
	if len(s) != 0 {
		t.Error("Expected '0', got ", len(s))
	}

}
