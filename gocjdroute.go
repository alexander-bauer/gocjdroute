package main

import (
	"github.com/SashaCrofter/cjdngo"
	"log"
)

const (
	Version  = "1.0"
	ConfPath = "./example.conf"
	//LogOutput = ioutil.Discard
)

var (
	Conf *cjdngo.Conf
)

func main() {
	Conf, err := cjdngo.ReadConf(ConfPath)
	if err != nil {
		log.Fatal(err)
	}
	//log.SetOutput(ioutil.Discard)

	Authorize(Conf, -1, nil)
}
