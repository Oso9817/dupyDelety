package main

import (
	"errors"
	"io/fs"
	"log"
	"testing"
)

func TestIterate(t *testing.T) {
	imageDir := "C:/Users/Alonzo/Programming/DisArchived/DisArchived/fullData/"
	_, err := iterate(imageDir)
	if err != nil {
		t.Fatalf("iterate(imageDir) has failed: = %v  error", err)
	}
}

func TestIterateEmpty(t *testing.T) {
	imageDir := " "
	_, err := iterate(imageDir)
	var e *fs.PathError
	if !errors.As(err, &e) {
		t.Fatalf("iterate(' ') test has failed: wanted: *fs.PathError, got: %v", err)
	}
}

func TestHashMapEmpty(t *testing.T) {
	foo, err := hashMap("", nil)
	if !errors.Is(err, errInvalid) {
		t.Fatalf("hashmap(' ', nil) test has failed: %v", err)
	}
	log.Println(foo)
}
