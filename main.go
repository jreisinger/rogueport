// Rogueport identifies TCP ports which are not supposed to be open.
package main

import (
	"flag"
	"log"
	"os"
)

var configFile = flag.String("c", "rogueport.json", "config file")

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
