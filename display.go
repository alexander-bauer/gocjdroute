package main

import (
	"github.com/SashaCrofter/cjdngo"
	"strconv"
	"encoding/json"
	"log"
)

//MakeConnectTo returns a string containing a JSON-formatted ConnectTo block, which can be sent to a potential peer, for them to import either by hand or using ConnectToJSON.
//It uses connection details from 'conf,' either 'TunConn' (if 'usePhys' is false) or 'PhysConn' and the password represented in 'p' to generate this block.
func MakeConnectTo(conf *cjdngo.Conf, p string, usePhys bool) string {
	outMap := map[string]cjdngo.Connection{}
	conn := cjdngo.Connection{conf.Name, conf.Location, conf.IPv6, p, conf.PublicKey}
	connectionDetail := ""
	//Get connection details here.
	if usePhys {
		connectionDetail = conf.PhysConn
	} else {
		connectionDetail = conf.TunConn
	}
	
	//Map the connection to the details (like in the local config,)
	outMap[connectionDetail] = conn
	
	//then encode the map with JSON,
	b, err := json.MarshalIndent(outMap, "", "    ")
	if err != nil {
		log.Fatal(err) //crash if unsuccessful (this probably isn't a good idea)
	}
	
	//then convert the []byte to a string, and return it.
	return string(b)
}

//ListAuth returns a string containing the information of the authorized nodes indicated by the 'indexes' argument. If 'indexes' is an empty []int, then all authorized nodes are listed. If 'showPass' is false, passwords for nodes will not be printed.
func ListAuth(conf *cjdngo.Conf, indexes []int, showPass bool) string {
	var s string

	for i := range indexes {
		index := indexes[i]

		s += "--- ( " + strconv.Itoa(index) + " )\n"
		s += "Name:     " + conf.AuthorizedPasswords[index].Name + "\n"
		if conf.AuthorizedPasswords[index].Location != "" {
			s += "Location: " + conf.AuthorizedPasswords[index].Location + "\n"
		}
		if conf.AuthorizedPasswords[index].IPv6 != "" {
			s += "IPv6:     " + conf.AuthorizedPasswords[index].IPv6 + "\n"
		}
		if showPass == true {
			s += "Password: " + conf.AuthorizedPasswords[index].Password
		}
		s += "\n"
	}
	return s
}

func ListConnectTo(conf *cjdngo.Conf, indexes []string, showPass bool) string {
	var s string

	for i := range indexes {
		index := indexes[i]

		s += "--- ( " + index + " )\n"
		s += "Name:      " + conf.Interfaces.UDPInterface.ConnectTo[index].Name + "\n"
		if conf.Interfaces.UDPInterface.ConnectTo[index].Location != "" {
			s += "Location:  " + conf.Interfaces.UDPInterface.ConnectTo[index].Location + "\n"
		}
		if conf.Interfaces.UDPInterface.ConnectTo[index].IPv6 != "" {
			s += "IPv6:      " + conf.Interfaces.UDPInterface.ConnectTo[index].IPv6 + "\n"
		}
		s += "PublicKey: " + conf.Interfaces.UDPInterface.ConnectTo[index].PublicKey + "\n"
		if showPass == true {
			s += "Password:  " + conf.Interfaces.UDPInterface.ConnectTo[index].Password
		}
		s += "\n"
	}
	return s
}
