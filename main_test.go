package main

import "testing"

func TestHello(t *testing.T) {
	expected := "hello world"
	result := Hello()

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
