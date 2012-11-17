package main

import (
	"encoding/json"
	"fmt"
	"github.com/SashaCrofter/cjdngo"
	"log"
	"math/rand"
	"time"
)

//Authorize takes a cjdngo.Conf object, an array index, a signed integer, and a cjdngo.AuthPass, which is meant to be parsed from user input. If the user input has been parsed into an AuthBlock already, then it can be passed and added directly. If they are not supplied (nil) then Authorize initiates a dialogue with the user to recieve them.
//The first argument refers to the index at which to edit or add the new password. If it is -1, then a new password and authorization block is added and appended to the list.
func Authorize(conf *cjdngo.Conf, index int, userDetails *cjdngo.AuthPass) {
	var auth *cjdngo.AuthPass

	//If index is out of the bounds of the existing array, or
	//if it's -1, then auth will be empty.
	if len(conf.AuthorizedPasswords) <= index || index == -1 {
		auth = &cjdngo.AuthPass{}
	} else {
		auth = &conf.AuthorizedPasswords[index]
	}

	//If we weren't passed any details already, we have to ask
	//for them. This is the hard bit.
	if userDetails == nil {
		var ui string

		//Take the name from the user. This is optional.
		fmt.Print("Please enter a name" + existing(auth.Name) + ": ")
		fmt.Scanln(&ui)
		auth.Name = replace(auth.Name, ui)
		ui = ""

		//Likewise, take the location from the user.
		fmt.Print("Please enter a location" + existing(auth.Location) + ": ")
		fmt.Scanln(&ui)
		auth.Location = replace(auth.Location, ui)
		ui = ""

		//The IPv6 address isn't usually known, but ask anyway.
		fmt.Print("Please enter an IPv6 address" + existing(auth.IPv6) + ": ")
		fmt.Scanln(&ui)
		auth.IPv6 = replace(auth.IPv6, ui)
		ui = ""

		if len(auth.Password) == 0 {
			auth.Password = replace(auth.Password, getPass(auth.Name))
		}
	}
	b, err := json.MarshalIndent(auth, "        ", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("        " + string(b))
}

//This is a convenience function that is meant to show a default value, when passed a string. The string is the existing value in a field, and it is returned as " (value)" if it exists, and "" if it does not exist.
func existing(value string) string {
	if len(value) != 0 {
		return " (" + value + ")"
	}
	return ""
}

//This is a convenience function that is meant to either replace an existing value, if an alternative is provided, or to leave it as-is. The first argument is a pointer to the string, and the second is the replacement candidate. It returns the result, as it determines that it should be..
func replace(original string, replacement string) string {
	if len(replacement) != 0 {
		return replacement
	}
	return original
}

func getPass(tag string) string {
	rand.Seed(time.Now().UnixNano()) //"randomize" the seed

	p := ""           //initialize an empty string for the password
	for len(p) < 16 { //getting 128 bits
		c := rand.Intn(0x59) //7 bit integers, 0x21-0x79 (inclusive)
		c += 0x21            //we need ascii characters
		if c == '\\' || c == '"' {
			continue
		}
		p += string(c)
	}

	if tag != "" {
		p = tag + "_" + p
	}
	return p
}
