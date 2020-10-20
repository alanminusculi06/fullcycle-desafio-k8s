package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	value := "Test"
	res := Greeting(value)
	expected := "<b>" + "Test" + "</b>"
	if res != expected {
		t.Errorf("Got: %s; Expected: %s", res, expected)
	}
}