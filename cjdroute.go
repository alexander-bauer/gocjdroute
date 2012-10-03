package main

import (
	"cjdngo"
	"os"
	"flag"
	"log"
)

var fFile = *flag.String("file", "./example.conf", "The config file to operate on.")
var fListAuth = *flag.Bool("list-auth", false, "List all authorized nodes.")
var fSearch = *flag.String("search", "", "Identify nodes based on matches of this string to their descriptor fields.")


var conf = &cjdngo.Conf{} //conf is the config file specified by fFile

var indexesAuth = []int{} //indexesAuth will be used to identify auth fields based on search terms
var indexesConnectTo = []int{} //connectAuth will be used to identify connection blocks based on search terms

func main() {
	flag.Parse()
	
	conf, err := cjdngo.ReadConf(fFile)
	if err != nil {
		log.Fatal(err)
	}
	
	if fSearch != "" {
		indexesAuth = SearchAuth(conf, fSearch)
		//BUG(DuoNoxSol): Does not implement indexesConnectTo
	}
	
	print(fListAuth + "\n")
	if fListAuth == true {
		print("listing\n")
		print(ListAuth(conf, indexesAuth, false))
	}
	
	os.Exit(0)
}
