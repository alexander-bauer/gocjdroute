package main

import (
	"crypto/rand"
	"math/big"
	"encoding/binary"
)

//
func GenPass(tag string) string {
	upperlim := big.NewInt(int64(0x59))
	lowerlim := uint16(0x21)
	
	p := ""           //initialize an empty string for the password
	for len(p) < 16 { //getting 128 bits
		c, _ := rand.Int(rand.Reader, upperlim) //7 bit integers, 0x21-0x79 (inclusive)
		v := binary.BigEndian.PutUint64(c)
		v += lowerlim //move in to the ascii range
		p += string(v)
	}
	password := tag + "_" + p
	return password
}
