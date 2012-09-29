package main

import (
	"math/rand"
	"time"
)

//
func GenPass(tag string) string {
	rand.Seed(time.Now().UnixNano()) //"randomize" the seed

	p := ""           //initialize an empty string for the password
	for len(p) < 16 { //getting 128 bits
		c := rand.Intn(0x59) //7 bit integers, 0x21-0x79 (inclusive)
		c += 0x21            //we need ascii characters
		p += string(c)
	}
	password := tag + "_" + p
	return password
}
