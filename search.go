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
