package main

import (
	"errors"
	//"github.com/corona10/goimagehash"
	"io/fs"
	"log"
	"testing"
)

func TestIterate(t *testing.T) {
	imageDir := "/"
	_, err := Iterate(imageDir)
	if err != nil {
		t.Fatalf("Iterate(imageDir) has failed: = %v  error", err)
	}
}

func TestIterateEmpty(t *testing.T) {
	imageDir := " "
	_, err := Iterate(imageDir)
	var e *fs.PathError
	if !errors.As(err, &e) {
		t.Fatalf("iterate(' ') test has failed wanted: *fs.PathError, got: %v", err)
	}
}

func TestHashMapEmpty(t *testing.T) {
	foo, err := HashMap("", nil)
	if !errors.Is(err, errInvalid) {
		t.Fatalf("hashmap(' ', nil) test has failed: %v", err)
	}
	log.Println(foo)
}

func TestProcessImageEmpty(t *testing.T) {
	_, err := ProcessImage("")
	var e *fs.PathError
	if !errors.As(err, &e) {
		t.Fatalf("ProcessImage('', '', nil) has failed wanted: *fs.PathError, got: %v", err)
	}
}

func TestProcessImageInvalid(t *testing.T) {
	_, err := ProcessImage("")
	var e *fs.PathError
	if !errors.As(err, &e) {
		t.Fatalf("ProcessImage('') has failed wanted: *fs.PathError, got: %v", err)
	}
}
func TestProcessImageValid(t *testing.T) {
	_, err := ProcessImage("moon1.jpg")
	if err != nil {
		t.Fatalf("ProcessImage('moon1.jpg') has failed wanted: nil, got: %v", err)
	}
}
