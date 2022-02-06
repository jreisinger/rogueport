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
	fmt.Printf("%-25s %s\n", h.name, strings.Join(out, " "))
}

func scanner(host string, ports, results chan int) {
	for port := range ports {
		if canConnect(host, port) {
			results <- port
		} else {
			results <- 0
		}
	}
}

func canConnect(host string, port int) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	for _, timeout := range []int{50, 100, 300} {
		conn, err := net.DialTimeout("tcp", addr, time.Millisecond*time.Duration(timeout))
		if err == nil {
			conn.Close()
			return true
		}
	}
	return false
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
