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
		s += "name:     " + conf.AuthorizedPasswords[index].Name + "\n"
		if conf.AuthorizedPasswords[index].Location != "" {
			s += "location: " + conf.AuthorizedPasswords[index].Location + "\n"
		}
		if conf.AuthorizedPasswords[index].IPv6 != "" {
			s += "IPv6:     " + conf.AuthorizedPasswords[index].IPv6 + "\n"
		}
		if showPass == true {
			s += "password: " + conf.AuthorizedPasswords[index].Password
		}
		s += "\n"
	}
	return s
}
