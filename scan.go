package main

import (
	"fmt"
	"strings"

	"github.com/Ullaakut/nmap/v2"
)

func scan(conf Ports) (map[string][]int, error) {
	var targets []string
	for host := range conf {
		targets = append(targets, host)
	}

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(targets...),
		nmap.WithMostCommonPorts(100),
	)
	if err != nil {
		return nil, err
	}

	result, _, err := scanner.Run()
	if err != nil {
		return nil, err
	}

	ports := make(map[string][]int)

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		// fmt.Printf("Host %q:\n", host.Addresses[0])
		// address := host.Addresses[0].Addr
		hostname := host.Hostnames[0].Name

		for _, port := range host.Ports {
			if port.Protocol != "tcp" || port.State.State != "open" {
				continue
			}
			// fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			ports[hostname] = append(ports[hostname], int(port.ID))
		}
	}

	return ports, nil
}

// eval evaluates expected and actual open ports
func eval(conf map[string][]int, scan map[string][]int) {
	for host, ports := range scan {
		var out []string
		var s string
		for _, port := range ports {
			if contains(conf[host], port) {
				s = fmt.Sprintf("%d✓", port)
			} else {
				s = fmt.Sprintf("%d✗", port)
			}
			out = append(out, s)
		}
		fmt.Printf("%-25s %s\n", host, strings.Join(out, " "))
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
