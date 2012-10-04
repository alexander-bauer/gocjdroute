package main

import (
	//"cjdngo"
)

//UIAuthorize is a wrapper for lower-level authorization functions. It authorizes the described node for connection, then prints the 'connection' block to send to the peer, based on optional fields in the Conf object.
//None of the arguments are required, but they are highly recommended (for finding the block in the future.)
func UIAuthorize(name, location, ipv6 string) {
	p := Authorize(conf, name, location, ipv6)
	print("New authorized node:\n" + ListAuth(conf, []int{len(conf.AuthorizedPasswords)-1}, true) + "\n")
	print(p + "\n")
}
