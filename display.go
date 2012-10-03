package main

import (
	"cjdngo"
)

//ListAuth returns a string containing the information of the authorized nodes indicated by the 'indexes' argument. If 'indexes' is an empty []int, then all authorized nodes are listed. If 'showPass' is false, passwords for nodes will not be printed.
func ListAuth(conf *cjdngo.Conf, indexes []int, showPass bool) string {
	var s string
	
	for i := range indexes {
		if s != "" {
			s += "\n---\n"
		}
		s += "name:     " + conf.AuthorizedPasswords[i].Name + "\n"
		if conf.AuthorizedPasswords[i].Location != "" {
			s += "location: " + conf.AuthorizedPasswords[i].Location + "\n"
		}
		if conf.AuthorizedPasswords[i].IPv6 != "" {
			s += "IPv6:     " + conf.AuthorizedPasswords[i].IPv6 + "\n"
		}
		if showPass == true {
			s += "password: " + conf.AuthorizedPasswords[i].Password
		}
	}
	return s
}
