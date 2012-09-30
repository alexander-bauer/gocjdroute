package main

import "testing"

func TestPass(t *testing.T) {
	for i := 0; i < 4; i++ {
		print(GenPass("test", 16) + "\n")
	}
}
