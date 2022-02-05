// Monport monitors your TCP ports. Based on config file it shows you which
// ports to open and which to close.
package main

import (
	"flag"
	"log"
	"os"
)

var configFile = flag.String("c", "monport.json", "config file")

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)

	flag.Parse()

	hosts, err := readConfigFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range hosts {
		h.scan()
		h.eval()
	}
}
