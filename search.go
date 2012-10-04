package main

import (
	"cjdngo"
	"strings"
)

//SearchAuth looks for 'term' in the non-essential elements of cjdngo.AuthPass, 'Name,' 'Location,' and 'IPv6.'
//Returns the indexes of every AuthPass with at least one partial match in any of the above fields.
func SearchAuth(conf *cjdngo.Conf, term string) []int {
	matches := make([]int, 0)

	for i := 0; i < len(conf.AuthorizedPasswords); i++ {
		//check if the Name field contains it...
		if strings.Contains(conf.AuthorizedPasswords[i].Name, term) {
			matches = append(matches, i)
		} else if strings.Contains(conf.AuthorizedPasswords[i].Location, term) {
			matches = append(matches, i)
		} else if strings.Contains(conf.AuthorizedPasswords[i].IPv6, term) {
			matches = append(matches, i)
		}
	}
	return matches
}

//SearchConnectTo searches for 'term' in both the non-essential and 'connection detail' fields in the ConnectTo map owned by 'conf.'
//Returns a []string containing a list of the 'connection detail' fields (which identify them within the map) of all connection blocks which match.
func SearchConnectTo(conf *cjdngo.Conf, term string) []string {
	matches := make([]string, 0)

	for i := range conf.Interfaces.UDPInterface.ConnectTo {
		//check if the connection detail field contains it...
		if strings.Contains(i, term) {
			matches = append(matches, i)
		} else if strings.Contains(conf.Interfaces.UDPInterface.ConnectTo[i].Name, term) {
			matches = append(matches, i)
		} else if strings.Contains(conf.Interfaces.UDPInterface.ConnectTo[i].Location, term) {
			matches = append(matches, i)
		} else if strings.Contains(conf.Interfaces.UDPInterface.ConnectTo[i].IPv6, term) {
			matches = append(matches, i)
		}
	}
	return matches
}
