package main

import "testing"

func TestExampleFunction(t *testing.T) {
<<<<<<< HEAD
	// TODO: Implement tests for the student's code
	// Example test (adjust according to the actual assignment)
	expected := "Hello!"
	if got := HelloFunction(); got != expected {
		t.Errorf("HelloFunction() = %v, want %v", got, expected)
	}
=======
    // TODO: Implement tests for the student's code
    // Example test (adjust according to the actual assignment)
    expected := "Hello"
    if got := HelloFunction(); got != expected {
        t.Errorf("HelloFunction() = %v, want %v", got, expected)
    }
>>>>>>> c14f48423ddfe158113d1d2431e9774091475058
}

func TestAdd(t *testing.T) {
    expected := 5
    if got := Add(2, 3); got != expected {
        t.Errorf("Add(2, 3) = %v, want %v", got, expected)
    }
}