package main

import (
	"cjdngo"
)

//Authorize adds a new entry to the authorizedPasswords block in 'conf,' generating a cryptographically strong password and returning it. Arguments 'name,' 'location,' and 'ipv6' are optional, but recommended for identifying the authorized node. It is highly recommended that every node be given a unique password.
func Authorize(conf *cjdngo.Conf, name string, location string, ipv6 string) string {
	p := GenPass(name)
	newAuth := &cjdngo.AuthPass{name, location, ipv6, p}
	conf.AuthorizedPasswords = append(conf.AuthorizedPasswords, *newAuth)
	return p
}

//ConnectTo adds an entry to the "connectTo" block (under 'interfaces.UDPInterface') with the given details. The arguments 'connection,' 'password,' and 'publicKey' are required, but 'name,' 'location,' and 'ipv6' are optional. They are recommended for identifying the target of the connection.
//This will exit without making changes to 'conf' if any of the required arguments are not provided.
func ConnectTo(conf *cjdngo.Conf, connection string, password string, publicKey string, name string, location string, ipv6 string) {
	if connection == "" || password == "" || publicKey == "" {
		return //if the required fields are not filled out, exit
	}
	
	conn := &cjdngo.Connection{name, location, ipv6, password, publicKey}
	//creating the connnection
	conf.Interfaces.UDPInterfaceBlock[connection] = conn
	//add the new entry
}

func ConnectToJSON(conf *cjdngo.Conf, connections map[string]conf.Connection) {
	if len(connections) == nil {
		return
	}
	for i, j := range connections { //for every item in connections,
		//i will be the connection information, such as ipv4 and port.
		//j will be the Connection interface itself.
		
		//This will map the new Connection interface to the connection information on the existing Conf object.
		conf.Interfaces.UDPInterface.ConnectTo[i] = connections[j]
	}
}