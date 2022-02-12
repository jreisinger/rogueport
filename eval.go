package main

import (
	"fmt"
	"strings"
)

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
