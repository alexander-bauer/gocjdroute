package main

import (
	"cjdngo"
	"log"
)

//UIAuthorize is a wrapper for lower-level authorization functions. It authorizes the described node for connection, then prints the 'connection' block to send to the peer, based on optional fields in the Conf object.
//None of the arguments are required, but they are highly recommended (for finding the block in the future.)
func UIAuthorize(conf *cjdngo.Conf, name, location, ipv6 string) {
	p := Authorize(conf, name, location, ipv6)
	print("New authorized node:\n" + ListAuth(conf, []int{len(conf.AuthorizedPasswords)-1}, true) + "\n")
	print("Details for new node:\n")
	print(MakeConnectTo(conf, p, false) + "\n")
	err := cjdngo.WriteConf(*fFile, *conf)
	log.Println("Wrote conf to file " + *fFile)
	if err != nil {
		log.Fatal(err)
	}
}

//UIConnectTo is a wrapper for lower-level connection functions. It creates a Connection to the described node, given enough information. The arguments 'name,' 'location,' and 'ipv6' are optional.
//UIConnectTo can also create connections from JSON strings. If the first argument begins with '{' and ends with '}' then it is interpreted as JSON, and its contents are imported as connection blocks. 
func UIConnectTo(conf *cjdngo.Conf, connectionDetails, password, publicKey, name, location, ipv6 string) {
	b := []byte(connectionDetails)
	if b[0] == '{' && b[len(b)-1] == '}' {
		//if the first argument looks like JSON, then send it to the JSON parser
		err := ConnectToJSON(conf, b)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		ConnectTo(conf, connectionDetails, password, publicKey, name, location, ipv6)
		print("New connection block:\n" + ListConnectTo(conf, []string{connectionDetails}, false))
	}
	err := cjdngo.WriteConf(*fFile, *conf)
	log.Println("Wrote conf to file " + *fFile)
	if err != nil {
		log.Fatal(err)
	}
	return
}
