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
