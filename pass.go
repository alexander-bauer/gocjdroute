package main

import (
	"crypto/rand"
	"math/big"
)

//
func GenPass(tag string) string {
	p := ""           //initialize an empty string for the password
	for len(p) < 16 { //getting 128 bits
		c, _ := rand.Int(rand.Reader, big.NewInt(int64(0xffffffff)))
		b := byte(c.Int64())
		if b > byte(0x20) && b < byte(0x80) {
			p += string(b) //if the character is ASCII, add it to the string
		}
	}
	password := tag + "_" + p
	return password
}
