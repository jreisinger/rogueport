package main

import (
	"fmt"
	"net"
	"sort"
	"strings"
	"time"
)

type host struct {
	name          string
	portsExpected []int // ports that are expected to be open
	portsOpen     []int // ports that are actually open
}

func (h *host) scan() {
	ports := make(chan int)
	results := make(chan int)

	for i := 0; i < 100; i++ {
		go scanner(h.name, ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	var openports []int

	for i := 1; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)
	h.portsOpen = openports
}

// eval evaluates expected and actual open ports
func (h *host) eval() {
	var out []string
	for _, p := range h.portsOpen {
		var s string
		if contains(h.portsExpected, p) {
			s = fmt.Sprintf("%d ✓", p)
		} else {
			s = fmt.Sprintf("%d ✗", p)
		}
		out = append(out, s)
	}
	fmt.Printf("%s\t%s\n", h.name, strings.Join(out, " "))
}

func scanner(host string, ports, results chan int) {
	for port := range ports {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", addr, time.Millisecond*100)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
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
