package main

import (
	"cjdngo"
	"os"
	"flag"
	"log"
)

var fFile = flag.String("file", "./example.conf", "The config file to operate on.")

var fList = flag.Bool("list", false, "List all authorization and connection blocks.")
var fListAuth = flag.Bool("list-auth", false, "List all authorized nodes.")
var fListConnectTo = flag.Bool("list-conn", false, "List all connection blocks.")

var fSearch = flag.String("search", "", "Identify nodes based on matches of this string to their descriptor fields.")

//BUG(DuoNoxSol): There's no flag for displaying passwords.

var conf = &cjdngo.Conf{} //conf is the config file specified by fFile

var indexesAuth = []int{} //indexesAuth will be used to identify auth fields based on search terms
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
			print(ListAuth(conf, indexesAuth, false))
			print(ListConnectTo(conf, indexesConnectTo, false))
			return
		} else if *fListAuth {
			print(ListAuth(conf, indexesAuth, false))
			return
		} else if *fListConnectTo {
			print(ListConnectTo(conf, indexesConnectTo, false))
			return
		}
	}
	
	os.Exit(0)
}
