package main

import (
	"cjdngo"
	"flag"
	"log"
	"os"
)

var fFile = flag.String("file", "./example.conf", "The config file to operate on.")

var fAuthorize = flag.Bool("auth", false, "Authorize and print connection details for a new node of form: [Node Owner's Name] [Node Location] [Node IPv6 Address]")
var fConnectTo = flag.Bool("connect", false, "Add a ConnectTo block entry for the described node of form: [Connection Information (such as IPv4:port)] [Password] [Public Key] [Node Owner's Name] [Node Location] [Node IPv6 Address]")

var fList = flag.Bool("list", false, "List all authorization and connection blocks.")
var fListAuth = flag.Bool("list-auth", false, "List all authorized nodes.")
var fListConnectTo = flag.Bool("list-conn", false, "List all connection blocks.")

var fSearch = flag.String("search", "", "Identify nodes based on matches of this string to their descriptor fields.")

var fShowPass = flag.Bool("show-pass", false, "Prints passwords when listing authorization and config blocks.")

//BUG(DuoNoxSol): There's no flag for displaying passwords.

var conf = &cjdngo.Conf{} //conf is the config file specified by fFile

var indexesAuth = []int{}         //indexesAuth will be used to identify auth fields based on search terms
var indexesConnectTo = []string{} //connectAuth will be used to identify connection blocks based on search terms

func main() {
	flag.Parse()

	conf, err := cjdngo.ReadConf(*fFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read conf file " + *fFile)

	if *fList || *fListAuth || *fListConnectTo {
		//This section will update the indexes for listing.
		indexesAuth = SearchAuth(conf, *fSearch)
		indexesConnectTo = SearchConnectTo(conf, *fSearch)

		if *fList {
			print(ListAuth(conf, indexesAuth, *fShowPass))
			print(ListConnectTo(conf, indexesConnectTo, *fShowPass))
			return
		} else if *fListAuth {
			print(ListAuth(conf, indexesAuth, *fShowPass))
			return
		} else if *fListConnectTo {
			print(ListConnectTo(conf, indexesConnectTo, *fShowPass))
			return
		}
	}
	
	if *fAuthorize {
		UIAuthorize(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	}
	if *fConnectTo {
	}

	os.Exit(0)
}
