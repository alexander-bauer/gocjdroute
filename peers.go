package main

import (
	"github.com/SashaCrofter/cjdngo"
	"encoding/json"
	"log"
)

//Authorize adds a new entry to the authorizedPasswords block in 'conf,' generating a cryptographically strong password and returning it. Arguments 'name,' 'location,' and 'ipv6' are optional, but recommended for identifying the authorized node. It is highly recommended that every node be given a unique password.
func Authorize(conf *cjdngo.Conf, name, location, ipv6 string) string {
	p := GenPass(name)
	newAuth := &cjdngo.AuthPass{name, location, ipv6, p}
	conf.AuthorizedPasswords = append(conf.AuthorizedPasswords, *newAuth)
	return p
}

//Removes the authorization block at 'index' from 'conf.' This permanently removes that node's password and details from the authorized section.
func RemoveAuth(conf *cjdngo.Conf, index int) {
	newAuthBlock := make([]cjdngo.AuthPass, len(conf.AuthorizedPasswords)-1)
	//This will copy all AuthPasses into newAuthBlock except the one specified by 'index'
	j := 0 //j will be the newAuthBlock index iterator
	for i := range conf.AuthorizedPasswords {
		if i != index {
			newAuthBlock[j] = conf.AuthorizedPasswords[i]
			j++
		}
	}
	conf.AuthorizedPasswords = newAuthBlock
}

//ConnectTo adds an entry to the "connectTo" block (under 'interfaces.UDPInterface') with the given details. The arguments 'connection,' 'password,' and 'publicKey' are required, but 'name,' 'location,' and 'ipv6' are optional. They are recommended for identifying the target of the connection.
//This will exit without making changes to 'conf' if any of the required arguments are not provided.
func ConnectTo(conf *cjdngo.Conf, connection, password, publicKey, name, location, ipv6 string) {
	if connection == "" || password == "" || publicKey == "" {
		log.Println("Not all required arguments provided.")
		return //if the required fields are not filled out, exit
	}

	conn := &cjdngo.Connection{name, location, ipv6, password, publicKey}
	//creating the connnection
	conf.Interfaces.UDPInterface.ConnectTo[connection] = *conn
	//add the new entry
}

//Removes the connection identified exactly by 'connectionDetail' from 'conf.'
func RemoveConnectTo(conf *cjdngo.Conf, connectionDetail string) {
	delete(conf.Interfaces.UDPInterface.ConnectTo, connectionDetail)
}

//TODO parse map from JSON []byte
//ConnectToJSON allows for adding one or more Connections from a []byte of JSON data. It unencodes the []byte, then calls ConnectToMap to perform the actual mapping.
func ConnectToJSON(conf *cjdngo.Conf, connections []byte) error {
	
	recv := &cjdngo.UDPInterfaceBlock{}
	
	err := json.Unmarshal(connections, recv)
	
	if err != nil {
		return err
	}
	ConnectToMap(conf, recv.ConnectTo)
	return nil
}
	
//ConnectToMap allows for adding multiple Connections in one operation. It is primarly meant as a back end for ConnectToJSON, but can be used alone, as well.
func ConnectToMap(conf *cjdngo.Conf, connections map[string]cjdngo.Connection) {
	if len(connections) == 0 {
		return
	}
	for i, j := range connections { //for every item in connections,
		//i will be the connection information, such as ipv4 and port.
		//j will be the Connection interface itself.

		//This will map the new Connection interface to the connection information on the existing Conf object.
		conf.Interfaces.UDPInterface.ConnectTo[i] = j
	}
}
