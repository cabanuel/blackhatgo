package main

import (
	"fmt"
	"net"
	"sort"
)
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("192.168.1.184:%d", p)
		conn, err := net.Dial("tcp",address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	// scanner for 1024 ports
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// start workers 
	for i := 0; i < cap(ports); i++ {
		go worker(ports,results)
		fmt.Println("worker started", i)
	}

	// start sending data to worker
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i 
			fmt.Println("sending data to ports", i)
		}
	} ()

	// process the results
	for i := 0; i < 1024; i++ {
		port := <- results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}