package main

import (
	"cjdngo"
	"strconv"
)

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
		s += "PublicKey: " + conf.Interfaces.UDPInterface.ConnectTo[index].PublicKey
		if showPass == true {
			s += "Password:  " + conf.Interfaces.UDPInterface.ConnectTo[index].Password
		}
		s += "\n"
	}
	return s
}
