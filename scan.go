package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
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
	var shouldOpen []int
	var shouldClose []int
	for _, expected := range h.portsExpected {
		if !contains(h.portsOpen, expected) {
			shouldOpen = append(shouldOpen, expected)
		}
	}
	for _, open := range h.portsOpen {
		if !contains(h.portsExpected, open) {
			shouldClose = append(shouldClose, open)
		}
	}
	fmt.Printf("%s %s\n", h.name, strings.Join(is2ss(h.portsOpen), ", "))
	if len(shouldOpen) > 0 {
		fmt.Printf("👉 you should open %s\n", strings.Join(is2ss(shouldOpen), ", "))
	}
	if len(shouldClose) > 0 {
		fmt.Printf("👉 you should close %s\n", strings.Join(is2ss(shouldClose), ", "))
	}
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

func is2ss(is []int) []string {
	var ss []string
	for _, i := range is {
		ss = append(ss, strconv.Itoa(i))
	}
	return ss
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
