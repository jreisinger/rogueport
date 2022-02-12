package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Ullaakut/nmap/v2"
)

func scan(hosts []string, mostCommonPorts int, syn, udp bool) (map[string][]string, error) {
	options := []nmap.Option{
		nmap.WithTargets(hosts...),
		nmap.WithMostCommonPorts(int(mostCommonPorts)), // --top-ports
		nmap.WithSkipHostDiscovery(),                   // -Pn
	}

	if syn {
		options = append(options, nmap.WithSYNScan()) // -sS
	}

	if udp {
		options = append(options, nmap.WithUDPScan()) // -sU
	}

	scanner, err := nmap.NewScanner(options...)
	if err != nil {
		return nil, err
	}

	progress := make(chan float32, 1)
	go func() { // Function to listen and print the progress
		for p := range progress {
			fmt.Printf("\rProgress: %3.0f%%", p)
		}
	}()
	result, warnings, err := scanner.RunWithProgress(progress)
	if len(warnings) != 0 {
		log.Printf("warnings: %s", strings.Join(warnings, " "))
	}
	if err != nil {
		return nil, err
	}

	ports := make(map[string][]string)
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		hostname := host.Hostnames[0].Name

		for _, port := range host.Ports {
			if port.State.State != "open" {
				continue
			}
			p := fmt.Sprintf("%d/%s", port.ID, port.Protocol)
			ports[hostname] = append(ports[hostname], p)
		}
	}
	return ports, nil
}

// eval evaluates expected and actual open ports on the hosts
func eval(conf map[string][]string, scan map[string][]string) {
	for host, ports := range scan {
		var out []string
		var s string
		for _, port := range ports {
			if contains(conf[host], port) {
				s = fmt.Sprintf("%s ✓", port)
			} else {
				s = fmt.Sprintf("%s ✗", port)
			}
			out = append(out, s)
		}
		fmt.Printf("\r%-25s %s\n", host, strings.Join(out, " "))
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
