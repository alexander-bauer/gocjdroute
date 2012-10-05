package main

import (
	"cjdngo"
	"flag"
	"log"
	"io/ioutil"
)

var fFile = flag.String("file", "./example.conf", "The config file to operate on.")

var fAuthorize = flag.Bool("auth", false, "Authorize and print connection details for a new node of form: [Node Owner's Name] [Node Location] [Node IPv6 Address]")
var fConnectTo = flag.Bool("connect", false, "Add a ConnectTo block entry for the described node of form: [Connection Information (such as IPv4:port)] [Password] [Public Key] [Node Owner's Name] [Node Location] [Node IPv6 Address]")
var fRemove = flag.Bool("remove", false, "Remove an authorization, identified by index, or connection, identified by connection details, block.")

var fList = flag.Bool("list", false, "List all authorization and connection blocks.")
var fListAuth = flag.Bool("list-auth", false, "List all authorized nodes.")
var fListConnectTo = flag.Bool("list-conn", false, "List all connection blocks.")

var fSearch = flag.String("search", "", "Identify nodes based on matches of this string to their descriptor fields.")

var fShowPass = flag.Bool("show-pass", false, "Prints passwords when listing authorization and config blocks.")

//BUG(DuoNoxSol): There's no flag for displaying passwords.

var config = &cjdngo.Conf{} //config is the config file specified by fFile

var indexesAuth = []int{}         //indexesAuth will be used to identify auth fields based on search terms
var indexesConnectTo = []string{} //connectAuth will be used to identify connection blocks based on search terms

func main() {
	log.SetOutput(ioutil.Discard) //Comment this line out for logging.
	flag.Parse()

	config, err := cjdngo.ReadConf(*fFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read conf file " + *fFile)

	if *fList || *fListAuth || *fListConnectTo {
		//This section will update the indexes for listing.
		indexesAuth = SearchAuth(config, *fSearch)
		indexesConnectTo = SearchConnectTo(config, *fSearch)

		if *fList {
			print(ListAuth(config, indexesAuth, *fShowPass))
			print(ListConnectTo(config, indexesConnectTo, *fShowPass))
			return
		} else if *fListAuth {
			print(ListAuth(config, indexesAuth, *fShowPass))
			return
		} else if *fListConnectTo {
			print(ListConnectTo(config, indexesConnectTo, *fShowPass))
			return
		}
	}
	
	if *fAuthorize {
		UIAuthorize(config, flag.Arg(0), flag.Arg(1), flag.Arg(2))
		return
	}
	if *fConnectTo {
		//UIConnectTo will import a JSON map if it is the first argument, and behave normally if given other arguments.
		UIConnectTo(config, flag.Arg(0), flag.Arg(1), flag.Arg(2), flag.Arg(3), flag.Arg(4), flag.Arg(5))
		return
	}
	if *fRemove {
		UIRemove(config, flag.Arg(0))
		return
	}
}
