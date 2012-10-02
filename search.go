package main

import (
	"cjdngo"
	"strings"
)

//SearchAuth looks for 'term' in the non-essential elements of cjdngo.AuthPass, 'Name,' 'Location,' and 'IPv6.'
//Returns the indexes of every AuthPass with at least one partial match in any of the above fields.
func SearchAuth(conf cjdngo.Conf, term string) []int {
	matches := make([]int, 0, len(conf.AuthorizedPasswords))
	
	for i := 0; i < len(conf.AuthorizedPasswords); i++ {
		//check if the Name field contains it...
		if strings.Contains(conf.AuthorizedPasswords[i].Name, term) {
			matches = append(matches, i)
			continue
		}
		//same with Location
		if strings.Contains(conf.AuthorizedPasswords[i].Location, term) {
			matches = append(matches, i)
			continue
		}
		//and IPv6
		if strings.Contains(conf.AuthorizedPasswords[i].IPv6, term) {
			matches = append(matches, i)
			continue
		}
	}
	return matches
}
