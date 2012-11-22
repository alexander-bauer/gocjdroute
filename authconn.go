package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/SashaCrofter/cjdngo"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Authorize takes a cjdngo.Conf object, the details to create the displayed connection block with (either TunConn or EthConn), an array index, a signed integer, and a cjdngo.AuthPass, which is meant to be parsed from user input. If the user input has been parsed into an AuthBlock already, then it can be passed and added directly. If they are not supplied (nil) then Authorize initiates a dialogue with the user to recieve them.
//The third argument refers to the index at which to edit or add the new password. If it is -1, then a new password and authorization block is added and appended to the list.
func Authorize(conf *cjdngo.Conf, details string, index int, jsonArg []byte) {
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
	if len(jsonArg) == 0 {
		//Take the name from the user. This is optional.
		ui("Please enter a name", &auth.Name)

		//Likewise, take the location from the user.
		ui("Please enter a location", &auth.Location)

		//The IPv6 address isn't usually known, but ask anyway.
		ui("Please enter an IPv6 address", &auth.IPv6)

		if len(auth.Password) == 0 {
			auth.Password = getPass(auth.Name)
		}
	} else {
		var inAuth cjdngo.AuthPass
		err := json.Unmarshal(jsonArg, &inAuth)
		if err != nil {
			println("Invalid authorized password.")
			return
		}
		auth = &inAuth
	}
	//Now we check whether we should append it to the end,
	//or assume the changes are already in place. (The latter
	//case is always true if we are editing.)
	if willAppend {
		conf.AuthorizedPasswords = append(conf.AuthorizedPasswords, *auth)
	}

	//Finally, we need to generate some connection details
	//that the authorized party can use. We check to make
	//sure that the connection details are in place, but
	//must only ask interactively if we haven't been given
	//JSON.
	if len(conf.Name) == 0 && len(jsonArg) == 0 {
		//If the user did not supply a name, ask for one,
		//which will be written back to the configuration.
		ui("Please enter your name or username", &conf.Name)
	}
	if len(conf.Location) == 0 && len(jsonArg) == 0 {
		//Similarly for location.
		ui("Please enter your displayed location", &conf.Location)
	}
	for len(conf.TunConn) == 0 && len(jsonArg) == 0 {
		var ipv4 string
		//If there aren't any details in place, force
		//the user to add them. These will be written
		//back to the configuration.
		ui("Please enter your IPv4 address", &ipv4)
		ipv4 += conf.Interfaces.UDPInterface.Bind[7:] //cjdns port
		conf.TunConn = ipv4
	}

	//Initialize a map with length 1, which will be used
	//to display our new details.
	display := make(map[string]*cjdngo.Connection, 1)

	display[details] = &cjdngo.Connection{
		Name:      conf.Name,
		Location:  conf.Location,
		IPv6:      conf.IPv6,
		Password:  auth.Password,
		PublicKey: conf.PublicKey,
	}
	b, err := json.MarshalIndent(display, "            ", "    ")
	if err != nil {
		log.Fatal(err)
	}
	println("\n            ", string(b), "\nPlease send these credentials to your peer.")
}

func Connect(conf *cjdngo.Conf, iface *cjdngo.InterfaceBlock, connDetails string, jsonArg []byte) {
	//If credentials are not provided prebuilt, then we must take
	//user input here.
	if len(jsonArg) == 0 {
		var conn *cjdngo.Connection

		//Check if the connDetails are already entered. If they are,
		//we'll be editing, rather than adding a new block.
		existing, isPresent := iface.ConnectTo[connDetails]
		if !isPresent {
			conn = &cjdngo.Connection{}
		} else {
			conn = &existing
		}

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

		iface.ConnectTo[connDetails] = *conn
		println("Connection to", connDetails, "added. You may want to restart cjdns.")
	} else { //If so, just use the JSON.
		var inConn map[string]cjdngo.Connection
		err := json.Unmarshal(jsonArg, &inConn)
		if err != nil || len(inConn) == 0 {
			println("Invalid connection.")
			return
		}
		for k, v := range inConn {
			iface.ConnectTo[k] = v
			println("Connection to", k, "added.")
		}
		println("You may want to restart cjdns.")
	}
}

//ListAuthorization is meant to display authorization blocks based on a search term. All authoriziation blocks are displayed if the term is omitted. Otherwise, only authorization blocks which have a name, location, IPv6, or password which partially matches the term are displayed.
func ListAuthorization(conf *cjdngo.Conf, term string) {
	display := make([]cjdngo.AuthPass, 0)
	indexes := make([]int, 0)

	for i := range conf.AuthorizedPasswords {
		pw := conf.AuthorizedPasswords[i]
		if strings.Contains(pw.Name, term) || strings.Contains(pw.Location, term) || strings.Contains(pw.IPv6, term) || strings.Contains(pw.Password, term) {
			display = append(display, pw)
			indexes = append(indexes, i)
		}
	}
	if len(display) == 0 {
		//If there are no elements to display,
		//don't bother marshalling the result.
		return
	}
	//Marshal the array of relevant results into a
	//displayable form.
	b, err := json.MarshalIndent(display, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(display); i++ {
		//Add a comment to every AuthPass element
		//marking it with its index.
		b = bytes.Replace(b, []byte("    {\n"), []byte("    { // "+strconv.Itoa(indexes[i])+"\n"), 1)
	}
	println(string(b))
}

//ListConnection is equivalent to ListAuthorization, except that it acts on conf.Interfaces.UDPInterface.ConnectTo. Additionally, it searches the PublicKey field of the connection for the term.
func ListConnection(iface *cjdngo.InterfaceBlock, term string) {
	display := make(map[string]cjdngo.Connection)

	for k, v := range conf.Interfaces.UDPInterface.ConnectTo {
		if strings.Contains(k, term) || strings.Contains(v.Name, term) || strings.Contains(v.Location, term) || strings.Contains(v.IPv6, term) || strings.Contains(v.Password, term) || strings.Contains(v.PublicKey, term) {
			display[k] = v
		}
	}
	if len(display) == 0 {
		//If there are no elements to display,
		//don't bother marshalling the result.
		return
	}
	b, err := json.MarshalIndent(display, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	println(string(b))
}

func Remove(conf *cjdngo.Conf, iface *cjdngo.InterfaceBlock, target string) {
	//Try to convert it to a number. If so,
	//then remove a password. Otherwise,
	//remove a connection.
	index, err := strconv.Atoi(target)
	switch err {
	case nil: //Password case
		if index >= len(conf.AuthorizedPasswords) || index < 0 {
			println("There is no password of that index.")
			return
		}
		//Initialize a new array with length - 1
		newAuth := make([]cjdngo.AuthPass, len(conf.AuthorizedPasswords)-1)

		//Copy the first part, stopping before removed index.
		copy(newAuth[:index], conf.AuthorizedPasswords[:index])
		//Copy the second part, starting after the removed index.
		copy(newAuth[index:], conf.AuthorizedPasswords[index+1:])
		conf.AuthorizedPasswords = newAuth

	default: //Connection case
		oldLen := len(iface.ConnectTo)
		delete(iface.ConnectTo, target)
		if oldLen == len(iface.ConnectTo) {
			println("There is no connection identified by that string.")
			return
		}
	}
}

//This is a convenience function which prints a given prompt onscreen in the form 'prompt (valueOfField): ' or just 'prompt:' if valueofField is blank.
func ui(prompt string, field *string) {
	reader := bufio.NewReader(os.Stdin)

	print(prompt + existing(*field) + ": ")
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	line = line[:len(line)-1] //Remove the newline byte.
	replace(field, line)
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
