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

	conf, err := readConfigFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	scan, err := scan(conf)
	if err != nil {
		log.Fatal(err)
	}
	eval(conf, scan)

	// for _, h := range hosts {
	// 	h.scan()
	// 	h.eval()
	// }
}
