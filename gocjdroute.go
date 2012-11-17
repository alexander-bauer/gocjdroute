package main

import (
	"flag"
	"github.com/SashaCrofter/cjdngo"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Version = "1.2"

	authCmd   = "auth"
	connCmd   = "conn"
	lsAuthCmd = "lsa"
	lsConnCmd = "lsc"
	rmCmd     = "rm"
)

var (
	Conf *cjdngo.Conf

	File string //The file (set by flag) to edit or view.

	cmd      string //The action to perform, (auth, conn, lsa, lsc, rm)
	argument string //The argument following cmd, a search term, index, or name
)

func init() {
	const (
		defaultFile = "./example.conf"
		usageFile   = "the cjdroute configuration file to edit or view"
	)
	flag.StringVar(&File, "file", defaultFile, usageFile)
	flag.StringVar(&File, "f", defaultFile, usageFile+" (shorthand)")
}

func usage() {
	println("GoCjdroute version", Version)
	println("usage:", os.Args[0], ""+authCmd+"|"+connCmd+"|"+lsAuthCmd+"|"+lsConnCmd+"|"+rmCmd+"", "[index, identifier, or search term]")
	println("Please use", os.Args[0], "--help", "for a list of flags.")
}

func main() {
	//Define the flags, and parse them.
	flag.Parse()

	cmd = strings.ToLower(flag.Arg(0))
	argument = flag.Arg(1)

	Conf, err := cjdngo.ReadConf(File)
	if err != nil {
		log.Fatal(err)
	}
	//log.SetOutput(ioutil.Discard)

	//This will be used to determine whether the
	//configuration should be rewritten afterward.
	willWrite := false

	//Perform an appropriate action, based on the subcommand.
	switch cmd {
	case authCmd:
		willWrite = true
		index, err := strconv.Atoi(argument)
		if err != nil {
			index = -1
		}
		Authorize(Conf, index, nil)

	case connCmd:
		willWrite = true
		Connect(Conf, argument, nil)

	case lsAuthCmd:
		ListAuthorization(Conf, argument)

	default:
		usage()
	}

	if willWrite {
		err = cjdngo.WriteConf(File, *Conf)
		if err != nil {
			log.Fatal(err)
		}
	}
}
