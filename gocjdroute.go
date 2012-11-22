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
	Version = "2.1"

	authCmd   = "auth"
	connCmd   = "conn"
	lsAuthCmd = "lsa"
	lsConnCmd = "lsc"
	rmCmd     = "rm"
)

var (
	Conf *cjdngo.Conf

	File    string //The file (set by flag) to edit or view.
	UseJSON bool   //Use argument as JSON of the appropriate type.
	UseETH bool //The argument to use ETHInterface

	cmd      string //The action to perform, (auth, conn, lsa, lsc, rm)
	argument string //The argument following cmd, a search term, index, or name
)

func init() {
	const (
		defaultFile = "/etc/cjdroute.conf"
		usageFile   = "the cjdroute configuration file to edit or view"

		defaultUseJSON = false
		usageUseJSON   = "supply a JSON-encoded block instead of interactive arguments"
		
		defaultUseETH = false
		usageUseETH = "act on the ETHInterface block instead of UDPInterface"
	)
	flag.StringVar(&File, "file", defaultFile, usageFile)
	flag.StringVar(&File, "f", defaultFile, usageFile+" (shorthand)")

	flag.BoolVar(&UseJSON, "json", defaultUseJSON, usageUseJSON)
	flag.BoolVar(&UseJSON, "j", defaultUseJSON, usageUseJSON+" (shorthand)")
	
	flag.BoolVar(&UseETH, "eth", defaultUseETH, usageUseETH)
	flag.BoolVar(&UseETH, "e", defaultUseETH, usageUseETH+" (shorthand)")
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

	var jsonArg []byte
	if UseJSON {
		var jsonTmp string
		for i := 2; i < flag.NArg(); i++ {
			//Put all of the remaining arguments into the string json.
			jsonTmp += flag.Arg(i)
		}
		jsonArg = []byte(jsonTmp)
	}

	Conf, err := cjdngo.ReadConf(File)
	if err != nil {
		log.Fatal(err)
	}
	
	var iface *cjdngo.InterfaceBlock
	if UseETH {
		iface = Conf.Interfaces.ETHInterface
	} else {
		iface = Conf.Interfaces.UDPInterface
	}
	//log.SetOutput(ioutil.Discard)

	//This will be used to determine whether the
	//configuration should be rewritten afterward.
	willWrite := false

	//Perform an appropriate action, based on the subcommand.
	switch cmd {
	case authCmd:
		willWrite = true
		
		var details string
		if UseETH {
			details = Conf.EthConn
		} else {
			details = Conf.TunConn
		}
		
		index, err := strconv.Atoi(argument)
		if err != nil {
			//If we can't parse the argument for whatever
			//reason, assume that it's an append.
			index = -1
			if UseJSON {
				//If we couldn't parse the argument, then
				//it might've been JSON, so treat it as
				//such.
				Authorize(Conf, index, append([]byte(argument), jsonArg...))
				break
			}
		}
		Authorize(Conf, details, index, jsonArg)

	case connCmd:
		willWrite = true
		if !UseJSON {
			Connect(Conf, iface, argument, nil)
		} else {
			Connect(Conf, iface, "", append([]byte(argument), jsonArg...))
		}

	case lsAuthCmd:
		ListAuthorization(Conf, argument)

	case lsConnCmd:
		ListConnection(iface, argument)

	case rmCmd:
		willWrite = true
		Remove(Conf, iface, argument)

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
