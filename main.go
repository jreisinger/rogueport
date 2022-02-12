// Rogueport identifies network ports which are not supposed to be open.
package main

import (
	"flag"
	"log"
	"os"
)

var configFile = flag.String("c", "rogueport.json", "config file")
var mostCommonPorts = flag.Int("n", 200, "number of most common ports to scan")
var syn = flag.Bool("s", false, "TCP SYN (half-open) scan; run with sudo")
var udp = flag.Bool("u", false, "UDP scan; run with sudo")

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)

	flag.Parse()

	conf, err := readConfigFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	var hosts []string

	for h := range conf {
		hosts = append(hosts, h)
	}

	scan, err := scan(hosts, *mostCommonPorts, *syn, *udp)
	if err != nil {
		log.Fatal(err)
	}

	eval(conf, scan)
}
