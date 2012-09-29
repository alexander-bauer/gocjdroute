package main

import (
	//"fmt"
	"flag"
	//"os"
)

var outpath = flag.String("output", "defaultPath.txt", "The place the out goes.")
var dispHelp = flag.Bool("help", false, "Displays this help message.")

func Parse() {
	flag.Parse()
	//fmt.Println("","");
}
