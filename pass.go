package main

import (
	"crypto/rand"
	"io"
)

//
func GenPass(tag string, length int) string {
	p := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, p); err != nil {
		panic(err)
	}
	for i := range p {
		p[i] &= 0x3f
		p[i] += 0x20
	}
	password := tag + "_" + string(p)
	return password
}
