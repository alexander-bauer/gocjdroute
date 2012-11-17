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
	var willAppend bool

	//If index is out of the bounds of the existing array, or
	//if it's -1, then auth will be empty.
	if len(conf.AuthorizedPasswords) <= index || index == -1 {
		auth = &cjdngo.AuthPass{}
		willAppend = true
	} else {
		auth = &conf.AuthorizedPasswords[index]
	}

	//If we weren't passed any details already, we have to ask
	//for them. This is the hard bit.
	if userDetails == nil {
		//Take the name from the user. This is optional.
		ui("Please enter a name", &auth.Name)

		//Likewise, take the location from the user.
		ui("Please enter a location", &auth.Location)

		//The IPv6 address isn't usually known, but ask anyway.
		ui("Please enter an IPv6 address", &auth.IPv6)

		if len(auth.Password) == 0 {
			auth.Password = getPass(auth.Name)
		}
	}
	b, err := json.MarshalIndent(auth, "        ", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("        " + string(b))
	if willAppend {
		conf.AuthorizedPasswords = append(conf.AuthorizedPasswords, *auth)
	}
}

func Connect(conf *cjdngo.Conf, connDetails string, credentials *cjdngo.Connection) {
	var conn *cjdngo.Connection

	//Check if the connDetails are already entered. If they are,
	//we'll be editing, rather than adding a new block.
	existing, isPresent := conf.Interfaces.UDPInterface.ConnectTo[connDetails]
	if !isPresent {
		conn = &cjdngo.Connection{}
	} else {
		conn = &existing
	}

	//If credentials are not provided prebuilt, then we must take
	//user input here.
	if credentials == nil {
		//Ask the user for a name for the connection, optional as usual.
		ui("Please enter a name", &conn.Name)
		ui("Please enter a location", &conn.Location)
		ui("Please enter an IPv6", &conn.IPv6)

		//Now for the technical items. If we weren't supplied
		//connection details, we must be adding a new item. In
		//that case, we must ask for more here.
		for len(connDetails) == 0 {
			ui("Please enter the target connection details", &connDetails)
		}
		ui("Please enter the password", &conn.Password)
		ui("Please enter the target's public key", &conn.PublicKey)
	}
	conf.Interfaces.UDPInterface.ConnectTo[connDetails] = *conn
	fmt.Println("Connection to", connDetails, "added. You may want to restart cjdns.")
}

//This is a convenience function which prints a given prompt onscreen in the form 'prompt (valueOfField): ' or just 'prompt:' if valueofField is blank.
func ui(prompt string, field *string) {
	var input string

	fmt.Print(prompt + existing(*field) + ": ")
	fmt.Scanln(&input)
	replace(field, input)
}

//This is a convenience function that is meant to show a default value, when passed a string. The string is the existing value in a field, and it is returned as " (value)" if it exists, and "" if it does not exist.
func existing(value string) string {
	if len(value) != 0 {
		return " (" + value + ")"
	}
	return ""
}

//This is a convenience function that is meant to either replace an existing value, if an alternative is provided, or to leave it as-is. The first argument is a pointer to the string, and the second is the replacement candidate. It returns the result, as it determines that it should be..
func replace(original *string, replacement string) {
	if len(replacement) != 0 {
		*original = replacement
	}
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
