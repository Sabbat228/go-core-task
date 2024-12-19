package main

import (
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	m := newStringIntMap()
	if m == nil {
		t.Fatal("Expected non-nil StringIntMap")
	}
	if len(m.StrInMap) != 0 {
		t.Fatal("Expected empty map")
	}
}

func TestAdd(t *testing.T) {
	m := newStringIntMap()
	m.Add("key1", 10)

	if val, exists := m.Get("key1"); !exists || val != 10 {
		t.Fatalf("Expected key1 to exist with value 10, got %v, exists: %v", val, exists)
	}
}

func TestRemove(t *testing.T) {
	m := newStringIntMap()
	m.Add("key1", 10)
	m.Remove("key1")

	if _, exists := m.Get("key1"); exists {
		t.Fatal("Expected key1 to be removed")
	}
}

func TestCopy(t *testing.T) {
	m := newStringIntMap()
	m.Add("key1", 10)
	m.Add("key2", 20)

	copy := m.Copy()
	if len(copy) != 2 {
		t.Fatalf("Expected copy to have 2 elements, got %d", len(copy))
	}
	if copy["key1"] != 10 || copy["key2"] != 20 {
		t.Fatal("Copy does not match original map")
	}
}

func TestExists(t *testing.T) {
	m := newStringIntMap()
	m.Add("key1", 10)

	if !m.Exists("key1") {
		t.Fatal("Expected key1 to exist")
	}
	if m.Exists("key2") {
		t.Fatal("Expected key2 to not exist")
	}
}

func TestGet(t *testing.T) {
	m := newStringIntMap()
	m.Add("key1", 10)

	if val, exists := m.Get("key1"); !exists || val != 10 {
		t.Fatalf("Expected key1 to return value 10, got %v, exists: %v", val, exists)
	}

	if _, exists := m.Get("key2"); exists {
		t.Fatal("Expected key2 to not exist")
	}
}
