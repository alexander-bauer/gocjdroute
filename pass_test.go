package main

import "testing"

func TestPass(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GenPass("test") + "\n")
	}
}
