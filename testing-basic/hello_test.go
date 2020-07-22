package main

import "testing"

func TestGetName(t *testing.T) {
	name := getName()
	if name != "there" {
		t.Error("Unexpected value from getName()")
	}
}

func TestGetNameFail(t *testing.T) {
	name := getName()
	if name != "unexpected" {
		t.Error("Unexpected value from getName()")
	}
}
